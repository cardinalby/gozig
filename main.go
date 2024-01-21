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
	var targetArgs []string
	if target, err := zig.GetTarget(gobuild.Os, gobuild.Arch); err == nil {
		targetArgs = []string{"-target", target}
	} else {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	sdkArgs := zig.GetPlatformArgs(gobuild.Os)

	args := append(targetArgs, sdkArgs...)
	args = append(args, os.Args[1:]...)

	cmd := exec.Command("zig", args...)
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
