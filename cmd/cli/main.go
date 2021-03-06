package main

import (
	"errors"
	"os"

	"github.com/cristianortiz/blacksmith"
	"github.com/fatih/color"
)

const version = "1.0.0"

var bls blacksmith.Blacksmith

func main() {
	var message string
	//create command line arguments
	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGracefully(err)
	}
	//setup blacksmith types for running migrations
	setup()

	switch arg1 {
	case "help":
		showHelp()
	case "version":
		color.Yellow("App version: " + version)

		//Ex command:blacksmith migrate up
	case "migrate":
		//set 'up' as default migrate sub-cmd
		if arg2 == "" {
			arg2 = "up"
		}
		err = doMigrate(arg2, arg3)
		if err != nil {
			exitGracefully(err)
		}
		message = "Migrations complete!"
	case "make":
		if arg2 == "" {
			exitGracefully(errors.New("'make' requieres subcommand: (migration|model|handler)"))
		}
		err = doMake(arg2, arg3)
		if err != nil {
			exitGracefully(err)
		}
	default:
		showHelp()

	}

	exitGracefully(nil, message)
}

//validateInput() checks if commands are entered in prompt line after
//the CLI command argument (blacksmith cmd sub-cmd sub-cmd-Option)
func validateInput() (string, string, string, error) {
	var arg1, arg2, arg3 string
	//check if command line arguments are being entered in prompt
	if len(os.Args) > 1 {
		//pos[0] is the command line argument (blacksmith arg1 arg2 arg3)
		arg1 = os.Args[1]
		//check for more than 2 cli arguments entered
		if len(os.Args) >= 3 {
			arg2 = os.Args[2]
		}
		if len(os.Args) >= 4 {
			arg3 = os.Args[3]
		}
	} else {
		color.Red("Error: command required")
		showHelp()
		return "", "", "", errors.New("command required")
	}

	return arg1, arg2, arg3, nil
}

//exitGracefully() ends the CLI app showing the appropiate messages to user
func exitGracefully(err error, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]

	}
	if err != nil {
		color.Red("Error: %v\n", err)
	}
	if len(message) > 0 {
		color.Yellow(message)
	} else {
		color.Green("Finished!")
	}
	os.Exit(0)
}
