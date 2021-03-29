package main

import (
	"fmt"
	"os"
	"os/exec"
	"bufio"
  )
  

func main() {
	fmt.Print("Create local react project - init process")

	prg := "npx"
    arg1 := "create-react-app"
    arg2 := "my-app"
	arg3 := "--yes"
	cmd := exec.Command(prg, arg1, arg2, arg3)
	cmd.Dir = "/Users/jefersonagudelo/project/own/progenerator/src/core/templates/frontend"


	cmdReader, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
        return
    }

    scanner := bufio.NewScanner(cmdReader)
    go func() {
        for scanner.Scan() {
            fmt.Printf("\t > %s\n", scanner.Text())
        }
    }()

    err = cmd.Start()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
        return
    }

    err = cmd.Wait()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
        return
    }

	fmt.Print("Create local react project - finish process")
  }