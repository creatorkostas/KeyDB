package users

import (
	"time"

	database "github.com/creatorkostas/KeyDB/database/database_core/core"
	"github.com/creatorkostas/KeyDB/internal/tools"
)

func Get_account(username string) *Account {
	var acc Account = accounts[username]
	return &acc
}

func Create_account(username string, acc_tier string, email string, password string) *Account {
	if Get_account(username).Username != "" {
		return nil
	}
	var api_string = username + email + password + time.Now().String()

	var acc Account = MakeDefaultUser()
	var userInfo UserInfo = UserInfo{Username: username, Api_key: hash(api_string)[0:16], Email: email, Password: hash(password)}
	acc.UserInfo = userInfo
	if acc_tier == "Admin" {
		acc.MakeAdmin()
	} else if acc_tier == "User" {
		acc.MakeUser()
	} else if acc_tier == "GuestUser" {
		acc.MakeGuestUser()
	}

	accounts[username] = acc
	// Accounts = append(Accounts, acc)
	database.MakeTable(acc.Username)
	return &acc
}

func SaveAccounts(filename string) error {
	var err error = nil

	tools.SaveToFile(filename, &accounts)

	return err
}

func LoadAccounts(filename string) error {
	var err error = nil

	tools.LoadFromFile(filename, &accounts)

	return err
}
