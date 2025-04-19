package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/uagarwal54/saberSlashes/terminal_chat_app/config"
)

func main() {
	fmt.Println("Welcome to Terminal Chat!")
	fmt.Println("Type 'help' to see available commands.")

	reader := bufio.NewReader(os.Stdin)
	absPath, _ := filepath.Abs("config/config.yaml")
	controlConfigStruct, err := config.InitializeConfigStruct(absPath)
	if err != nil {
		log.Fatal(err)
	}
	var dbInUseConfig config.Db
	for _, db := range controlConfigStruct.DBs {
		if db.Name == controlConfigStruct.DbInUse {
			dbInUseConfig = db
			break
		}
	}
	_, err = config.InitializeSupabase(dbInUseConfig.ConnString)
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Print(">>> ")
		input, _ := reader.ReadString('\n')
		command := strings.TrimSpace(input)

		switch command {
		case "help":
			showHelp()
		case "login":
			login()
		case "register":
			register()
		case "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command. Type 'help' for options.")
		}
	}
}

func showHelp() {
	fmt.Println("Available commands:")
	fmt.Println("- register: Create a new account.")
	fmt.Println("- login: Log into your account.")
	fmt.Println("- exit: Exit the chat application.")
}

func login() {
	fmt.Println("Login feature not implemented yet!")
	// Placeholder for future implementation
}

func register() {
	fmt.Println("Register feature not implemented yet!")
	// Placeholder for future implementation
}
