package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	devenv "github.com/Minizbot2012/DevEnv"
)

var exetab map[string]string

func main() {
	jsf, err := os.Open(path.Join(devenv.ExeDir, "loki.json"))
	if err != nil {
		panic(err)
	}
	var cmd *exec.Cmd
	json.NewDecoder(jsf).Decode(&exetab)
	if exe, ok := exetab[os.Args[0]]; ok {
		exe = devenv.Replace(exe)
		if filepath.IsAbs(exe) {
			cmd = exec.Command(exe, os.Args[1:]...)
		} else {
			cmd = exec.Command(path.Join(devenv.ExeDir, exe), os.Args[1:]...)
		}
	} else if exe, ok := exetab[os.Args[1]]; ok {
		exe = devenv.Replace(exe)
		if filepath.IsAbs(exe) {
			cmd = exec.Command(exe, os.Args[2:]...)
		} else {
			cmd = exec.Command(path.Join(devenv.ExeDir, exe), os.Args[2:]...)
		}
	}
	if cmd != nil {
		cmd.Stdout, _ = os.OpenFile(os.DevNull, os.O_WRONLY, 0755)
		cmd.Stderr, _ = os.OpenFile(os.DevNull, os.O_WRONLY, 0755)
		cmd.Start()
	}
}
