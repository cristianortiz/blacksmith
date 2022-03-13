package main

import (
	"errors"
	"log"
	"os"

	"github.com/cristianortiz/blacksmith"
	"github.com/fatih/color"
)

const version = "1.0.0"

var bls blacksmith.Blacksmith

func main() {
	//create command line arguments
	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGracefully(err)
	}
	switch arg1 {
	case "help":
		showHelp()
	case "version":
		color.Yellow("App version: " + version)
	default:
		log.Print(arg2, arg3)

	}
}

//validateInput() checks if commands are entered in prompt line after the command line argumnt
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

//showHelp() print a text with help info about the cli commands an dtheir functions
func showHelp() {
	color.Yellow(`Available commands:
	help         - show the help commands
	version      - print app version
	
	`)
}

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
