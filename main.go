package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/cardinalby/gozig/pkg/gobuild"
	"github.com/cardinalby/gozig/pkg/zig"
)

func main() {
	zigArgs, err := zig.PrepareCompilerArgs(os.Args[1:], gobuild.Os, gobuild.Arch)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	executeZigCmd(zigArgs)
}

func executeZigCmd(args []string) {
	cmd := exec.Command("zig", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			os.Exit(exitError.ExitCode())
		} else {
			_, _ = fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
	}
}
