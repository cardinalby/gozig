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
	zigCmd, zigCmdArgs, err := zig.GetCmd(os.Args[1:])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	var targetArgs []string
	if target, err := zig.GetTarget(gobuild.Os, gobuild.Arch); err == nil && len(target) > 0 {
		targetArgs = []string{"-target", target}
	} else {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	sdkArgs := zig.GetPlatformArgs(gobuild.Os)

	args := []string{zigCmd}
	args = append(args, targetArgs...)
	args = append(args, sdkArgs...)
	args = append(args, zigCmdArgs...)

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
