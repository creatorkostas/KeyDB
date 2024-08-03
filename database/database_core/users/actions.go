package users

import (
	"time"

	database "github.com/creatorkostas/KeyDB/database/database_core"
)

func Get_account(username string) *Account {
	for _, account := range Accounts {
		if account.Username == username {
			return &account
		}
	}
	return nil
}

func Create_account(username string, acc_tier string, email string, password string) *Account {
	if Get_account(username) != nil {
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
	} else if acc_tier == "FreeUser" {
		acc.MakeFreeUser()
	}

	Accounts = append(Accounts, acc)
	database.MakeTable(acc.Username)
	return &acc
}
