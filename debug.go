package main

import (
	"os"
	"strconv"

	"github.com/keegancsmith/shell"
)

const debugExecEnv = "DEBUG_EXEC"

func isDebugMode() bool {
	gbeDebug, ok := os.LookupEnv(debugExecEnv)
	if !ok {
		return false
	}
	if isGbeDebug, err := strconv.ParseBool(gbeDebug); err != nil {
		return false
	} else {
		return isGbeDebug
	}
}

func fmtArgs(args []string) string {
	return shell.Sprintf("%S", args)
}
