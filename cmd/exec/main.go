package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var exetab map[string]string

func main() {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	jsf, err := os.Open(path.Join(exPath, "exetab.json"))
	if err != nil {
		panic(err)
	}
	var cmd *exec.Cmd
	json.NewDecoder(jsf).Decode(&exetab)
	if exe, ok := exetab[os.Args[0]]; ok {
		cmd = exec.Command(path.Join(exPath, exe), os.Args[1:]...)
	} else if exe, ok := exetab[os.Args[1]]; ok {
		cmd = exec.Command(path.Join(exPath, exe), os.Args[2:]...)
	}
	if cmd != nil {
		cmd.Stdout = os.NewFile(0, os.DevNull)
		cmd.Stderr = os.NewFile(0, os.DevNull)
		cmd.Start()
	}
}
