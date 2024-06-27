package users

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/creatorkostas/KeyDB/internal/database"
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

func hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
