//go:build darwin
// +build darwin

package vm

import (
	"os/user"
)

const libName = "libhera.dylib"

func WASMLibLocation() string {
	usr, err := user.Current()
	if err != nil {
		return ""
	}
	return usr.HomeDir + "/kalyan3104-vm-binaries/" + libName
}
