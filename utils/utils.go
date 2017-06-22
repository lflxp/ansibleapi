package utils

import (
	"fmt"
	"os/exec"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ExecCommand(commands string) string {
	println(commands)
	out,err := exec.Command("bash", "-c", commands).Output()
	if err != nil {
	    	fmt.Println(err.Error())
	}
	return string(out)
}