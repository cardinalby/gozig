package zig

import (
	"fmt"
	"os"
	"path"
)

func GetPlatformArgs(goos string) []string {
	if goos == "darwin" {
		return GetOsxSdkArgs()
	}
	return nil
}

func GetOsxSdkArgs() []string {
	osxSdk := os.Getenv("GOZIG_OSX_SDK")
	if osxSdk == "" {
		return nil
	}

	return []string{
		fmt.Sprintf("--sysroot=%s", osxSdk),
		fmt.Sprintf("-I%s", path.Join(osxSdk, "usr/include")),
		fmt.Sprintf("-L%s", path.Join(osxSdk, "usr/lib")),
		fmt.Sprintf("-F%s", path.Join(osxSdk, "System/Library/Frameworks")),
	}
}
