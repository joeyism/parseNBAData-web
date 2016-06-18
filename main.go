package main

import (
	"os/exec"
	"fmt"
)

func getData()(string, error){
	cmd := exec.Command("./python/parseNBAData.py", "playbyplay20092010reg20100418", "20091027BOSCLE", "Garnett")
	out, err := cmd.Output()

	if err != nil {
		println("error!")
		println(err.Error())
		return "", err
	}

	return string(out), nil;
}

func main() {
	result, _:= getData()
	fmt.Printf(result)
}
