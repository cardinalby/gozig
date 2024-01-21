package zig

import "errors"

func GetCmd(args []string) (cmd string, zigCmdArgs []string, err error) {
	if len(args) == 0 {
		return "", nil, errors.New("no zig command provided")
	}
	if args[0] != "cc" && args[0] != "c++" {
		return "", nil, errors.New("only 'cc' or 'c++' zig commands are supported")
	}
	return args[0], args[1:], nil
}
