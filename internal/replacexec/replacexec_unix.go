//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package replacexec

import (
	"os"

	"golang.org/x/sys/unix"
)

func Replacexec(cmd string, args ...string) error {
	argv := []string{cmd}
	argv = append(argv, args...)
	return unix.Exec(cmd, argv, os.Environ())
}
