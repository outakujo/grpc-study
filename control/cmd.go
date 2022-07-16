package control

import (
	"io"
	"os/exec"
)

func Exe(wr io.Writer, name string, arg ...string) (string, error) {
	command := exec.Command(name, arg...)
	if wr != nil {
		command.Stdout = wr
		err := command.Run()
		return "", err
	}
	output, err := command.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
