package cmd_api

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	internal "github.com/creatorkostas/KeyDB/database/database_core/conf"
	"github.com/creatorkostas/KeyDB/database/database_core/persistance"
	"github.com/creatorkostas/KeyDB/database/database_core/users"
	db_utils "github.com/creatorkostas/KeyDB/database/database_core/utils"
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
			case "makeacc":
				var username, acc_tier, email, password string = "", "", "", ""

				username = print_and_get("username: ")
				acc_tier = print_and_get("acc_tier: ")
				email = print_and_get("email: ")
				password = print_and_get("password: ")

				var acc = users.Create_account(username, acc_tier, email, password)
				fmt.Print("Account : ")
				fmt.Println(acc)
			case "stopweb":
				StopWeb()
			case "startwebdev":
				var port int64
				port, _ = strconv.ParseInt(print_and_get("Port: "), 10, 64)
				setAndStartRemote(true, int(port))
			case "startweb":
				var port int64
				port, _ = strconv.ParseInt(print_and_get("Port: "), 10, 64)
				setAndStartRemote(false, int(port))
			case "exit":
				StopWeb()
				persistance.Operations <- "||exit||"

				db_utils.SaveDB(internal.DB_filename)
				users.SaveAccounts(internal.Accounts_filename)
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
