package users

import "time"

const (
	ADMIN     = 0
	USER      = 1
	FREE_USER = 2
	DEFAULT   = 3
)

const (
	CAN_GET = 0
	CAN_ADD = 1
)

type UserInfo struct {
	Username string
	Api_key  string
	Email    string
	Password string
}

type Account struct {
	UserInfo
	AccountState
	Tier Tier
}

func (acc *Account) IsAdmin() bool {
	return acc.Tier.Rules.Admin
}

func (acc *Account) CanAdd() bool {
	return acc.Tier.Rules.Add
}

func (acc *Account) CanGet() bool {
	return acc.Tier.Rules.Get
}

func (acc *Account) CanChangePassword() bool {
	return acc.Tier.Rules.Change_password
}

func (acc *Account) ChangePassword(new_password string) bool {
	var new_pass string = hash(new_password)
	acc.UserInfo.Api_key = new_pass
	return true
}

func (acc *Account) CanChangeApiKey() bool {
	return acc.Tier.Rules.Change_api_key
}

func (acc *Account) ChangeApiKey() string {
	var new_api_key string = hash(acc.Username + acc.Email + acc.Password + time.Now().String())[0:16]
	acc.UserInfo.Api_key = new_api_key
	return new_api_key
}

func (acc *Account) CanGetAnalytics() bool {
	return acc.Tier.Rules.Analytics
}

func MakeDefaultUser() Account {
	var acc Account
	acc.Tier.Type = DEFAULT
	acc.Tier.Rules = Default_rules
	acc.AccountState = Default_state
	return acc
}

func (acc *Account) MakeAdmin() {
	acc.Tier.Type = ADMIN
	acc.Tier.Rules = Admin_rules
	acc.AccountState = Admin_state
}

func (acc *Account) MakeUser() {
	acc.Tier.Type = USER
	acc.Tier.Rules = User_rules
	acc.AccountState = User_state
}

func (acc *Account) MakeFreeUser() {
	acc.Tier.Type = FREE_USER
	acc.Tier.Rules = Free_user_rules
	acc.AccountState = Free_user_state
}

var Accounts = make([]Account, 50)
