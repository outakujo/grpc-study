package control

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"strconv"
)

func Exe(wr io.Writer, name string, arg ...string) (string, error) {
	command := exec.Command(name, arg...)
	if wr != nil {
		command.Stdout = wr
		err := command.Run()
		return "", err
	}
	output, err := command.CombinedOutput()
	if err != nil {
		return "", err
	}
	code := command.ProcessState.ExitCode()
	if code != 0 {
		return "", errors.New("exit code " + strconv.Itoa(code))
	}
	return string(output), nil
}

const scriptDir = "control/script/"

func ExeScript(wr io.Writer, name string, arg ...string) (string, error) {
	workdir := "workdir"
	_ = os.Mkdir(workdir, os.ModePerm)
	bytes, err := Asset(scriptDir + name)
	if err != nil {
		return "", err
	}
	temp, err := os.CreateTemp(workdir, name)
	if err != nil {
		return "", err
	}
	_, err = temp.Write(bytes)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = os.Remove(temp.Name())
	}()
	err = os.Chmod(temp.Name(), 0500)
	if err != nil {
		return "", err
	}
	stat, err := temp.Stat()
	if err != nil {
		return "", err
	}
	_ = temp.Close()
	command := exec.Command("./"+stat.Name(), arg...)
	command.Dir = workdir
	if wr != nil {
		command.Stdout = wr
		err := command.Run()
		return "", err
	}
	output, err := command.CombinedOutput()
	if err != nil {
		return "", err
	}
	code := command.ProcessState.ExitCode()
	if code != 0 {
		return "", errors.New("exit code " + strconv.Itoa(code))
	}
	return string(output), nil
}
