package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		args = append(args, []string{"help"}...)
	}

	all, err := filepath.Glob(args[len(args)-1])
	if err != nil {
		log.Fatal(err)
	}

	args = args[:len(args)-1]

	gofiles := make([]string, 0, len(all))
	for _, filename := range all {
		if args[0] == "run" {
			if !strings.HasSuffix(filename, "_test.go") {
				gofiles = append(gofiles, filename)
			}
		} else {
			gofiles = append(gofiles, filename)
		}
	}

	args = append(args, gofiles...)

	fmt.Println(args)

	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
