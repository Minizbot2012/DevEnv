package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	devenv "github.com/Minizbot2012/DevEnv"
)

var exetab map[string]struct {
	Exec   string
	Silent bool
}

func main() {
	jsf, err := os.Open(path.Join(devenv.ExeDir, "loki.json"))
	if err != nil {
		panic(err)
	}
	var cmd *exec.Cmd
	var silent bool
	json.NewDecoder(jsf).Decode(&exetab)
	if exe, ok := exetab[os.Args[0]]; ok {
		silent = exe.Silent
		exe := devenv.Replace(exe.Exec)
		if filepath.IsAbs(exe) {
			cmd = exec.Command(exe, os.Args[1:]...)
		} else {
			cmd = exec.Command(path.Join(devenv.ExeDir, exe), os.Args[1:]...)
		}
	} else if exe, ok := exetab[os.Args[1]]; ok {
		silent = exe.Silent
		exe := devenv.Replace(exe.Exec)
		if filepath.IsAbs(exe) {
			cmd = exec.Command(exe, os.Args[2:]...)
		} else {
			cmd = exec.Command(path.Join(devenv.ExeDir, exe), os.Args[2:]...)
		}
	}
	if cmd != nil {
		cmd.Dir = devenv.CWD
		if silent {
			cmd.Stdout, _ = os.OpenFile(os.DevNull, os.O_WRONLY, 0755)
			cmd.Stderr, _ = os.OpenFile(os.DevNull, os.O_WRONLY, 0755)
			cmd.Start()
		} else {
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			cmd.Stderr = os.Stderr
			cmd.Run()
		}
	}
}
