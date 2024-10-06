package cmd

import "strings"

// ParseCommand will parse the user's command and populate the main Command struct
func ParseCommand(input string) *Command {
	splittedInput := strings.Split(input, " ")
	var userCmd Command
	userCmd.Name = splittedInput[0]
	userCmd.Args = splittedInput[1:]
	return &userCmd
}
