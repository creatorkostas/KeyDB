package cmd_api

import (
	"errors"

	database "github.com/creatorkostas/KeyDB/database/database_core"
	internal "github.com/creatorkostas/KeyDB/database/database_core/conf"
	"github.com/creatorkostas/KeyDB/database/database_core/persistance"
	"github.com/creatorkostas/KeyDB/database/database_core/users"
)

func GetValue(key string, acc *users.Account) (any, error) {

	var err error = nil
	var result any = database.Get_value(acc.Username, "table.get.all.data")
	if result == nil {
		err = errors.New("key does not exist")
	}

	return result, err

}

func SetValues(key string, acc *users.Account, value_type string, data string) (bool, error) {

	var err error = nil
	var ok bool = false
	err = database.Add_value(acc.Username, key, value_type, data)

	if err == nil {
		ok = true
	}

	return ok, err
}

func Register(username string, email string, password string, acc_type string) (*users.Account, error) {
	var acc *users.Account = nil

	var err error = nil

	acc = users.Create_account(username, acc_type, email, password)
	if acc == nil {
		err = errors.New("something went wrong")
	}

	return acc, err
}

func ChangeApiKey(acc *users.Account) (string, error) {
	var err error = nil
	var new_api string = acc.ChangeApiKey()
	return new_api, err
}

func ChangePassword(acc *users.Account, new_pass string) (bool, error) {
	var err error = nil
	var ok bool = acc.ChangePassword(new_pass)
	return ok, err
}

func Save() (bool, error) {
	persistance.SaveToFile(internal.DB_filename, &database.DB)
	persistance.SaveToFile(internal.Accounts_filename, &users.Accounts)

	return true, nil
}

func Load() (bool, error) {
	persistance.LoadFromFile(internal.DB_filename, &database.DB)
	persistance.LoadFromFile(internal.Accounts_filename, &users.Accounts)

	return true, nil
}
