package main

import (
	"bufio"
	"fmt"
	"os"
	c "saberShell/cmd"
	"strings"
)

func main() {
	var input string
	for {
		fmt.Print("saberShell> ")
		reader := bufio.NewReader(os.Stdin)
		var err error
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while taking input: ", err)
			os.Exit(1)
		}
		input = strings.TrimSpace(input)
		userCmd := c.ParseCommand(input)
		if err := userCmd.Execute(); err != nil {
			fmt.Println(err)
		}
	}
}
