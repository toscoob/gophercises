package main

import (
	task "github.com/gophercises/taskmanager"
	"github.com/mitchellh/cli"
	"log"
	"os"
)

func main() {
	c := cli.NewCLI("task", "1.0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"add":  func() (cli.Command, error) {
			return &task.AddCommand{}, nil
		},
		"do":  func() (cli.Command, error) {
			return &task.DoCommand{}, nil
		},
		"list":  func() (cli.Command, error) {
			return &task.ListCommand{}, nil
		},
		"reset":  func() (cli.Command, error) {
			return &task.ResetCommand{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
