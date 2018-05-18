package tool

import (
	"io/ioutil"
	"os/exec"
)

//use to exec some terminal command
func Command(sCmd string) (res string) {
	cmd := exec.Command(sCmd)

	stdout, err := cmd.StdoutPipe()
	defer stdout.Close()
	CheckErr(err)

	err = cmd.Start()
	CheckErr(err)

	opBytes, err := ioutil.ReadAll(stdout)
	CheckErr(err)
	return string(opBytes)
}
