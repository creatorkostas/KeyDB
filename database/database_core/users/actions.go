package users

import (
	"errors"
	"time"

	database "github.com/creatorkostas/KeyDB/database/database_core/core"
	"github.com/creatorkostas/KeyDB/internal/tools"
)

func Get_account(username string) *Account {
	var acc Account = accounts[username]
	return &acc
}

func Create_account(username string, acc_tier string, email string, password string) (*Account, string, error) {
	var err error = nil
	var acc *Account = Get_account(username)
	var username_exist bool = acc.Username != ""
	var email_exist bool = acc.Email != ""
	var private_key string

	if !username_exist && !email_exist {

		*acc = MakeDefaultUser()
		var api_string = username + email + password + time.Now().String()

		var userInfo UserInfo = UserInfo{Username: username, Api_key: hash(api_string)[0:16], Email: email, Password: hash(password)}
		acc.UserInfo = userInfo
		if acc_tier == "Admin" {
			acc.MakeAdmin()
		} else if acc_tier == "User" {
			acc.MakeUser()
		} else if acc_tier == "GuestUser" {
			acc.MakeGuestUser()
		}

		private_key = acc.create_RSA_keys()
	} else if username_exist {
		err = errors.New("username already exist")
	} else if email_exist {
		err = errors.New("email already exist")
	} else {
		err = errors.New("something went wrong")
	}

	accounts[username] = *acc
	// Accounts = append(Accounts, acc)
	database.MakeTable(acc.Username)
	return acc, private_key, err
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

// TODO add and handle errors
func DeleteAccount(username string) {
	delete(accounts, username)
	database.DeleteTable(username)
}
