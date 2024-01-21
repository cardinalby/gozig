package gobuild

import (
	"os"
	"runtime"
)

var Os string
var Arch string

func init() {
	Os = os.Getenv("GOOS")
	Arch = os.Getenv("GOARCH")
	if Os == "" {
		Os = runtime.GOOS
	}
	if Arch == "" {
		Arch = runtime.GOARCH
	}
}
