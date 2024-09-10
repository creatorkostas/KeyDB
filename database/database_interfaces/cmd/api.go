package cmd_interface

import (
	"fmt"
	"os"
	"strings"
	"sync"

	api "github.com/creatorkostas/KeyDB/database/database_api"
	"github.com/creatorkostas/KeyDB/database/database_core/persistance"
	"github.com/creatorkostas/KeyDB/database/database_core/users"
	"github.com/gin-gonic/gin"
	"github.com/jedib0t/go-pretty/table"
	stats "github.com/semihalev/gin-stats"
)

var wg sync.WaitGroup

func Cmd_start(user *users.Account) {
	wg.Add(1)
	go func(user *users.Account) {
		var command string = ""
		for {
			command = print_and_get("keyDB >> ")
			switch command {
			case "save":
				api.Save(user)
			case "load":
				api.Load(user)
			case "getacc":
				var acc_name = print_and_get("Username: ")
				fmt.Println((api.GetAccount(user, acc_name)).Description.(*users.Account))
			case "getallacc":
				c_table := table.NewWriter()
				c_table.SetOutputMirror(os.Stdout)
				c_table.AppendHeader(table.Row{"Username", "Email", "Api_key", "Active", "Tokens", "Tier"})
				for _, acc := range users.GetAllAccounts() {
					c_table.AppendRows([]table.Row{
						{acc.UserInfo.Username, acc.UserInfo.Email, acc.UserInfo.Api_key, acc.AccountState.Active, acc.AccountState.Tokens, acc.Tier},
					})
				}
				c_table.Render()
			case "editacc":
				var acc_name = print_and_get("Username: ")
				var acc = (api.GetAccount(user, acc_name)).Description.(*users.Account)

				fmt.Println(acc)
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
				if err == nil {
					fmt.Print("Account : ")
					fmt.Println(acc)
					fmt.Print("Private RSA key : ")
					fmt.Println(private_key)
				} else {
					fmt.Println(err.Error())
				}
			case "deleteacc":
				var acc_name = print_and_get("Username: ")
				var sure = print_and_get("Are you sure? [y/n]: ")
				if sure == "y" {
					users.DeleteAccount(acc_name)
					fmt.Println("Account deleted!")
				} else {
					fmt.Println("Aborted!")
				}
			case "chengeacctier":
				var acc *users.Account
				var tier = print_and_get("1) Admin\n2) User\n3) Guest\n--> ")
				switch tier {
				case "1":
					var acc_name = print_and_get("Username: ")
					acc = (api.GetAccount(user, acc_name)).Description.(*users.Account)
					acc.MakeAdmin()
					fmt.Println(acc)
				case "2":
					var acc_name = print_and_get("Username: ")
					acc = (api.GetAccount(user, acc_name)).Description.(*users.Account)
					acc.MakeUser()
					fmt.Println(acc)
				case "3":
					var acc_name = print_and_get("Username: ")
					acc = (api.GetAccount(user, acc_name)).Description.(*users.Account)
					acc.MakeGuestUser()
					fmt.Println(acc)
				default:
					fmt.Println("Invalid command!")
				}
				fmt.Println(acc)
			case "stopweb":
				api.StopWeb(user)
			case "startwebdev":
				var port string
				port = print_and_get("Port: ")
				api.StartWeb(user, true, port)
			case "startweb":
				var port string
				port = print_and_get("Port: ")
				api.StartWeb(user, false, port)
			case "startunix":
				api.StartUnix(user)
			case "exit":
				api.StopWeb(user)
				persistance.Operations <- "||exit||"

				api.Save(user)
				wg.Done()
				os.Exit(0)
			case "getcurrentacc":
				c_table := table.NewWriter()
				c_table.SetOutputMirror(os.Stdout)
				c_table.AppendHeader(table.Row{"Username", "Email", "Api_key", "Active", "Tokens", "Tier"})
				c_table.AppendRows([]table.Row{
					{user.UserInfo.Username, user.UserInfo.Email, user.UserInfo.Api_key, user.AccountState.Active, user.AccountState.Tokens, user.Tier},
				})
				c_table.Render()
			case "help":
				c_table := table.NewWriter()
				c_table.SetOutputMirror(os.Stdout)
				c_table.AppendHeader(table.Row{"Command", "Description"})
				c_table.AppendRows([]table.Row{
					{"save", "save the database"},
					{"load", "load the database"},
					{"getacc", "get account"},
					{"getcurrentacc", "get current account"},
					{"getallacc", "get all accounts"},
					{"editacc", "edit account"},
					{"report", "get stats"},
					{"makeacc", "make account"},
					{"deleteacc", "delete account"},
					{"chengeacctier", "change account tier"},
					{"stopweb", "stop web server"},
					{"startwebdev", "start web server in dev mode"},
					{"startweb", "start web server"},
					{"startunix", "start unix server"},
					{"exit", "exit"},
				})
				c_table.Render()
			case "":
				continue
			default:
				fmt.Println("Invalid command!")
			}
		}
	}(user)
	wg.Wait()
}

func print_and_get(print string) string {
	var str string = ""
	fmt.Print(print)
	fmt.Scanln(&str)
	return strings.Replace(str, print, "", 1)
}

func StartKeyDB(acc *users.Account, dev bool, start_web bool, port string, start_unix bool, RouterSetupFunc func(router *gin.Engine), RouterAddEndpointsFunc func(router *gin.Engine)) {
	api.RouterSetupFunc = RouterSetupFunc
	api.RouterAddEndpointsFunc = RouterAddEndpointsFunc
	api.StartKeyDB(acc, dev, start_web, port, start_unix)
}
