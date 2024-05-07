//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package replacexec

import (
	"os"

	"golang.org/x/sys/unix"
)

func Replacexec(cmd string, argv ...string) error {
	return unix.Exec(cmd, argv, os.Environ())
}
