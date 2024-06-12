package handlers

import (
	"crypto/sha256"
	"encoding/hex"
)

const (
	ADMIN     = 0
	USER      = 1
	FREE_USER = 2
)

type Account struct {
	Username string
	Tier     Tier
	Api_key  string
	Email    string
	Password string
}

type Rules struct {
	Admin                 bool
	Add                   bool
	Get                   bool
	Change_password       bool
	Change_change_api_key bool
	Analytics             bool
}

type Tier struct {
	Type  int8
	Rules Rules
}

var Admin_rules Rules = Rules{Admin: true, Add: true, Get: true, Change_password: true, Change_change_api_key: true, Analytics: true}
var User_rules Rules = Rules{Admin: false, Add: true, Get: true, Change_password: true, Change_change_api_key: true, Analytics: false}
var Free_user_rules Rules = Rules{Admin: false, Add: true, Get: true, Change_password: false, Change_change_api_key: false, Analytics: false}

// var Accounts = make([]Account, 10)
var Accounts []Account = []Account{}

func Get_account(username string) *Account {
	for _, account := range Accounts {
		if account.Username == username {
			return &account
		}
	}
	return nil
}

func Create_account(username string, acc_tier string, email string, password string) *Account {
	if Get_account(username) == nil {
		return nil
	}
	var api_string = username + email + password

	var tier Tier = Tier{ADMIN, Admin_rules}
	var acc = Account{Username: username, Tier: tier, Api_key: hash(api_string)[0:16], Email: email, Password: hash(password)}

	Accounts = append(Accounts, acc)
	return &acc
}

func hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
