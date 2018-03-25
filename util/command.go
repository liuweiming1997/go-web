package util

import (
	"io/ioutil"
	"os/exec"

	"github.com/sundayfun/go-web/logs"
)

//use to exec some terminal command
func Command(sCmd string) (res string) {
	cmd := exec.Command(sCmd)

	stdout, err := cmd.StdoutPipe()
	defer stdout.Close()
	logs.CheckErr(err)

	err = cmd.Start()
	logs.CheckErr(err)

	opBytes, err := ioutil.ReadAll(stdout)
	logs.CheckErr(err)
	return string(opBytes)
}
