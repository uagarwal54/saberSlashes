package cmd

import (
	"fmt"
	"os/exec"
)

// Execute is the function that executes entered command
func (cmd *Command) Execute() error {
	var cmdExecutor *exec.Cmd
	if len(cmd.Args) > 0 {
		cmdExecutor = exec.Command(cmd.Name, cmd.Args...)
	} else {
		cmdExecutor = exec.Command(cmd.Name)
	}

	stdout, err := cmdExecutor.Output()
	if err != nil {
		fmt.Println("Error while executing the command")
		return err
	}

	fmt.Print(string(stdout))

	return nil
}
