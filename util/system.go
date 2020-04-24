package util

import (
	"io/ioutil"
	"os/exec"
)

// ReadFile //
func ReadFile(file string) (string, error) {
	content, e := ioutil.ReadFile(file)
	return string(content), e
}

// Execute //
func Execute(cmd string, arg ...string) (string, error) {
	output, e := exec.Command(cmd, arg...).Output()
	return string(output), e
}
