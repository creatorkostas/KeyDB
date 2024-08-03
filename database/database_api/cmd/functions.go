package cmd_api

import (
	"fmt"
	"strings"
)

func Cmd_start() {
	go func() {
		var command string = ""
		for {
			command = print_and_get()
			switch command {
			case "save":
				Save()
			case "load":
				Load()
			default:
				fmt.Println("Invalid command!")
			}
		}
	}()
}

func print_and_get() string {
	var str string = ""
	fmt.Print("KeyDB >> ")
	fmt.Scanln(&str)
	return strings.ToLower(strings.Replace(str, "keyDB >> ", "", 1))
}
