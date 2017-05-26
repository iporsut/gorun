package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	all, err := filepath.Glob("*.go")
	if err != nil {
		log.Fatal(err)
	}

	gofiles := make([]string, 0, len(all))
	for i := range all {
		if !strings.HasSuffix(all[i], "_test.go") {
			gofiles = append(gofiles, all[i])
		}
	}

	args := append([]string{"run"}, gofiles...)

	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
