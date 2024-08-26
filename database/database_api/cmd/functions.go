package cmd_api

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/creatorkostas/KeyDB/database/database_core/persistance"
	"github.com/creatorkostas/KeyDB/database/database_core/users"
	stats "github.com/semihalev/gin-stats"
)

var wg sync.WaitGroup

func Cmd_start() {
	wg.Add(1)
	go func() {
		var command string = ""
		for {
			command = print_and_get("keyDB >> ")
			switch command {
			case "save":
				Save()
			case "load":
				Load()
			case "getacc":
				var acc_name = print_and_get("Username: ")
				fmt.Println(GetAccount(acc_name))
			case "report":
				fmt.Println(stats.Report().GetAll())
				fmt.Println(stats.Report())
			case "makeacc":
				var username, acc_tier, email, password string = "", "", "", ""

				username = print_and_get("username: ")
				acc_tier = print_and_get("acc_tier: ")
				email = print_and_get("email: ")
				password = print_and_get("password: ")

				var acc, private_key, err = users.Create_account(username, acc_tier, email, password)
				if err != nil {
					fmt.Print("Account : ")
					fmt.Println(acc)
					fmt.Print("Private RSA key : ")
					fmt.Println(private_key)
				} else {
					fmt.Println(err.Error())
				}
			case "stopweb":
				StopWeb()
			case "startwebdev":
				var port string
				port = print_and_get("Port: ")
				setAndStartRemote(true, port)
			case "startweb":
				var port string
				port = print_and_get("Port: ")
				setAndStartRemote(false, port)
			case "exit":
				StopWeb()
				persistance.Operations <- "||exit||"

				Save()
				wg.Done()
				os.Exit(0)
			default:
				fmt.Println("Invalid command!")
			}
		}
	}()
	wg.Wait()
}

func print_and_get(print string) string {
	var str string = ""
	fmt.Print(print)
	fmt.Scanln(&str)
	return strings.ToLower(strings.Replace(str, print, "", 1))
}
