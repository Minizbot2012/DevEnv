package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path"

	devenv "github.com/Minizbot2012/DevEnv"
)

type vars struct {
	Name  string
	Value string
}

type Env struct {
	//add to environment
	Environ []vars `json:"vars"`
	//append to PATH
	PathAppend []string `json:"path"`
	//Vars to keep and pass
	Passthrough []string `json:"keep"`
	//Terminal to use
	Terminal string
}

func main() {
	//Init config file
	jsf, err := os.Open(path.Join(devenv.CWD, "config_"+devenv.OS+".json"))
	if err != nil {
		jsf, err = os.Open(path.Join(devenv.CWD, "config.json"))
		if err != nil {
			panic(err)
		}
	}
	defer jsf.Close()
	var environ Env
	//Decode config file
	err = json.NewDecoder(jsf).Decode(&environ)
	if err != nil {
		panic(err)
	}
	// Default passthrough
	environ.Passthrough = append(environ.Passthrough, "SSH_AUTH_SOCK")
	if devenv.OS == "windows" {
		environ.Passthrough = append(environ.Passthrough,
			"PROCESSOR_ARCHITECTURE", "PROCESSOR_IDENTIFIER",
			"NUMBER_OF_PROCESSORS", "USERNAME",
			"PSModulePath", "PATHEXT", "OS", "DriverData",
			"PROGRAMFILES", "ProgramFiles(x86)", "ProgramData",
			"APPDATA", "LOCALAPPDATA", "TMP", "TEMP",
			"COMSPEC")
	} else if devenv.OS == "linux" {
		environ.Passthrough = append(environ.Passthrough,
			"DISPLAY", "WAYLAND_DISPLAY", "LANG", "HOME", "XDG_RUNTIME_DIR",
			"XDG_SESSION_TYPE", "XDG_SESSION_DESKTOP", "USER", "SHELL",
			"SESSION_MANAGER", "TERM", "XAUTHORITY", "DESKTOP_SESSION",
			"HISTCONTROL", "COLORTERM", "DBUS_SESSION_BUS_ADDRESS")
	}

	//do our replacements
	for i, v := range environ.Environ {
		environ.Environ[i].Value = devenv.Replace(v.Value)
	}
	for i, v := range environ.PathAppend {
		environ.PathAppend[i] = devenv.Replace(v)
	}

	//Initiate Terminal
	term, err := exec.LookPath(environ.Terminal)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(term)
	cmd.Env = []string{}
	path := os.Getenv("PATH")
	if path[len(path)-1] != ';' {
		path = path + string(os.PathListSeparator)
	}
	//Path Appends
	for _, v := range environ.PathAppend {
		path = path + v + string(os.PathListSeparator)
	}
	cmd.Env = append(cmd.Env, "PATH="+path)

	//Environment Variables
	for _, v := range environ.Environ {
		cmd.Env = append(cmd.Env, v.Name+"="+v.Value)
	}

	//Passthrough Variables
	for _, v := range environ.Passthrough {
		cmd.Env = append(cmd.Env, v+"="+os.Getenv(v))
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
