package zig

const BinaryName = "zig"

func PrepareCompilerArgs(
	appArgs []string,
	goos string,
	goarch string,
) (res []string, err error) {
	zigCmd, zigCmdArgs, err := DecomposeCompilerCmd(appArgs)
	if err != nil {
		return nil, err
	}
	res = []string{zigCmd}

	if target, err := GetTarget(goos, goarch); err != nil {
		return nil, err
	} else if len(target) > 0 {
		res = append(res, "-target", target)
	}

	res = append(res, GetPlatformArgs(goos)...)
	res = append(res, zigCmdArgs...)

	return res, nil
}
