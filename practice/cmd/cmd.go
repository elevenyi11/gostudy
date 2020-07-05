package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	name, _ := os.Hostname()
	fmt.Println(name)
	//command := "ls"
	//params := []string{"-l"}
	cmd := exec.Command("ping", "127.0.0.1", "-n", "2")
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("error")
	}
	//execCommand(command, params)
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return false
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return true
}
