package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/cardinalby/gozig/pkg/zig"
)

func main() {
	target, err := zig.GetTargetByGoEnv()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	args := os.Args[1:]
	if len(target) > 0 {
		args = append([]string{"-target", target}, args...)
	}

	cmd := exec.Command("zig", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	err = cmd.Run()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			os.Exit(exitError.ExitCode())
		} else {
			_, _ = fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
	}
}
