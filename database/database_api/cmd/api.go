package cmd_api

import (
	"context"
	"errors"
	"net/http"

	web_api "github.com/creatorkostas/KeyDB/database/database_api/web"
	internal "github.com/creatorkostas/KeyDB/database/database_core/conf"
	database "github.com/creatorkostas/KeyDB/database/database_core/core"
	"github.com/creatorkostas/KeyDB/database/database_core/persistance"
	"github.com/creatorkostas/KeyDB/database/database_core/users"
	db_utils "github.com/creatorkostas/KeyDB/database/database_core/utils"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func GetValue(key string, acc *users.Account) (any, error) {

	var err error = nil
	var result any = database.Get_value(acc.Username, "table.get.all.data", false)
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

func Register(username string, email string, password string, acc_type string) (*users.Account, string, error) {
	var acc, private_key, err = users.Create_account(username, acc_type, email, password)
	return acc, private_key, err
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
	db_utils.SaveDB(internal.DB_filename)
	users.SaveAccounts(internal.Accounts_filename)

	return true, nil
}

func Load() (bool, error) {
	db_utils.LoadDB(internal.DB_filename)
	users.LoadAccounts(internal.Accounts_filename)

	return true, nil
}

func GetAccount(username string) *users.Account {
	var acc = users.Get_account(username)
	return acc
	// fmt.Println(acc)
	// if acc == nil {
	// 	return "The user does not exist"
	// }
	// return acc.Username
}

func StartKeyDB(dev bool, start_web bool, port string) {

	db_utils.LoadDB(internal.DB_filename)
	users.LoadAccounts(internal.Accounts_filename)

	persistance.Start_writers(1)

	if start_web {
		setAndStartRemote(dev, port)
	}

}

func setAndStartRemote(dev bool, port string) {

	if dev {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router = gin.New()

	web_api.Setup_router(router)
	web_api.Add_endpoints(router)

	router.Run(":" + port)
}

func StopWeb() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	srv.Shutdown(context.Background())
}
