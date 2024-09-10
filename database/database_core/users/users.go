package users

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"strings"
	"time"
)

const (
	ADMIN      = 0
	USER       = 1
	GUEST_USER = 2
	DEFAULT    = 3
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
	Tier       Tier
	Public_key string
}

func (acc *Account) create_RSA_keys() string {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	publicKey := &privateKey.PublicKey

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	// privateKey

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	acc.Public_key = string(publicKeyPEM)

	return string(privateKeyPEM)
}

func (acc *Account) ToSting() string {
	var acc_string strings.Builder
	// acc_string.WriteString(acc.UserInfo.Username)
	// acc_string.WriteString(acc.UserInfo.Email)
	// acc_string.WriteString(acc.UserInfo.Password)
	// acc_string.WriteString(acc.UserInfo.Api_key)
	// acc_string.WriteString(acc.AccountState.Active)
	// acc_string.WriteString(acc.AccountState.Active)
	// acc_string.WriteString(acc.Tier)
	return acc_string.String()
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
	acc.UserInfo.Password = new_pass
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
	acc.Tier = Default_tier
	acc.AccountState = Default_state
	return acc
}

func (acc *Account) MakeAdmin() {
	acc.AccountState = Admin_state
	acc.Tier = Admin_tier
}

func (acc *Account) MakeUser() {
	acc.AccountState = User_state
	acc.Tier = User_tier
}

func (acc *Account) MakeGuestUser() {
	acc.AccountState = Guest_user_state
	acc.Tier = Guest_user_tier
}

func GetAllAccounts() map[string]Account {
	return accounts
}

// var Accounts = make([]Account, 50)
var accounts map[string]Account = make(map[string]Account, 100)
