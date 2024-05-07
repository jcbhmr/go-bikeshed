package replacexec

import (
	"errors"
	"os"
	"os/exec"
)

func Replacexec(cmd string, argv ...string) error {
	c := exec.Command(cmd, argv[1:]...)
	c.Args[0] = argv[0]
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	var exitError *exec.ExitError
	if errors.As(err, &exitError) {
		os.Exit(exitError.ExitCode())
	}
	return err
}
