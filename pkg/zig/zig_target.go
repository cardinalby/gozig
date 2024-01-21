package zig

import (
	"fmt"
	"runtime"
)

func GetTarget(goos, goarch string) (t string, err error) {
	// not a cross-build
	if goarch == runtime.GOARCH && goos == runtime.GOOS {
		return "", nil
	}

	switch goos {
	case "windows":
		switch goarch {
		case "386":
			t = "i386-windows-gnu"
		case "amd64":
			t = "x86_64-windows-gnu"
		case "arm":
			t = "arm-windows-gnu"
		case "arm64":
			t = "aarch64-windows-gnu"
		}
	case "linux":
		switch goarch {
		case "386":
			t = "i386-linux-gnu"
		case "amd64":
			t = "x86_64-linux-gnu"
		case "arm":
			t = "arm-linux-gnueabi"
		case "arm64":
			t = "aarch64-linux-gnu"
		case "mips":
			t = "mips-linux-gnu"
		case "mips64":
			t = "mips64-linux-musl"
		case "mips64le":
			t = "mips64el-linux-musl"
		case "mipsle":
			t = "mipsel-linux-gnu"
		}
	case "darwin":
		switch goarch {
		case "amd64":
			t = "x86_64-macos"
		case "arm64":
			t = "aarch64-macos"
		}
	case "wasip1":
		switch goarch {
		case "wasm":
			t = "wasm32-wasi"
		}
	}
	if t == "" {
		err = fmt.Errorf("unsupported GOOS/GOARCH: %s/%s", goos, goarch)
	}
	return
}
