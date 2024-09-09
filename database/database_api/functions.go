package api

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/creatorkostas/KeyDB/database/database_core/conf"
	internal "github.com/creatorkostas/KeyDB/database/database_core/conf"
	database "github.com/creatorkostas/KeyDB/database/database_core/core"
	"github.com/creatorkostas/KeyDB/database/database_core/persistance"
	"github.com/creatorkostas/KeyDB/database/database_core/security"
	"github.com/creatorkostas/KeyDB/database/database_core/users"
	db_utils "github.com/creatorkostas/KeyDB/database/database_core/utils"

	// web_api "github.com/creatorkostas/KeyDB/database/database_interfaces/web"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var router_set bool = false
var RouterSetupFunc func(router *gin.Engine) = nil
var RouterAddEndpointsFunc func(router *gin.Engine) = nil

func GetValue(key string, acc *users.Account) ActionResponce {

	if acc.CanGet() {
		var err error = nil
		var result any = database.Get_value(acc.Username, "table.get.all.data", false)
		if result == nil {
			err = errors.New("key does not exist")
		}

		return ActionResponce{Error: err, Code: http.StatusOK, From: "GetValue", Description: result}
	}

	return ActionResponce{Error: errors.New("account does not have permission to get data"), Code: http.StatusUnauthorized, From: "GetValue", Description: "Get data"}

}

func SetValues(key string, acc *users.Account, value_type string, data string, encrypt bool) ActionResponce {

	if acc.CanAdd() {
		var err error = nil

		if encrypt {
			err = database.Add_value(acc.Username, key, value_type, security.Decrypt_data(acc.Public_key, []byte(data)), false, false, "")
		} else {
		}
		err = database.Add_value(acc.Username, key, value_type, data, false, false, "")

		var ok bool = false

		if err == nil {
			ok = true
		}

		return ActionResponce{Error: err, Code: http.StatusOK, From: "SetValues", Description: ok}
	}

	return ActionResponce{Error: errors.New("account does not have permission to add data"), Code: http.StatusUnauthorized, From: "SetValues", Description: "Add data"}

}

func Register(username string, email string, password string, acc_type string) (*users.Account, string, error) {
	var acc, private_key, err = users.Create_account(username, acc_type, email, password)
	return acc, private_key, err
}

func ChangeApiKey(acc *users.Account) ActionResponce {
	if acc.CanChangeApiKey() {
		var err error = nil
		var new_api string = acc.ChangeApiKey()
		return ActionResponce{Error: err, Code: http.StatusOK, From: "ChangeApiKey", Description: new_api}
	}
	return ActionResponce{Error: errors.New("account does not have permission to change api key"), Code: http.StatusUnauthorized, From: "ChangeApiKey", Description: "Change api key"}
}

func ChangePassword(acc *users.Account, new_pass string) ActionResponce {
	if acc.CanChangePassword() {
		var err error = nil
		acc.ChangePassword(new_pass)
		return ActionResponce{Error: err, Code: http.StatusOK, From: "ChangePassword", Description: "Change password"}
	}
	return ActionResponce{Error: errors.New("account does not have permission to change password"), Code: http.StatusUnauthorized, From: "ChangePassword", Description: "Change password"}
}

func Save(acc *users.Account) ActionResponce {
	if acc.IsAdmin() {
		db_utils.SaveDB(internal.DB_filename)
		users.SaveAccounts(internal.Accounts_filename)

		return ActionResponce{Error: nil, Code: http.StatusOK, From: "Save", Description: "Save DB"}
	}
	return ActionResponce{Error: errors.New("not admin"), Code: http.StatusUnauthorized, From: "Save", Description: "Save DB"}
}

func Load(acc *users.Account) ActionResponce {
	if acc.IsAdmin() {
		db_utils.LoadDB(internal.DB_filename)
		users.LoadAccounts(internal.Accounts_filename)
		return ActionResponce{Error: nil, Code: http.StatusOK, From: "Load", Description: "Load DB"}
	}
	return ActionResponce{Error: errors.New("not admin"), Code: http.StatusUnauthorized, From: "Load", Description: "Load DB"}
}

func GetAccount(acc *users.Account, username string) ActionResponce {
	if acc.IsAdmin() {
		var acc = users.Get_account(username)
		return ActionResponce{Error: nil, Code: http.StatusOK, From: "GetAccount", Description: acc}
	}

	return ActionResponce{Error: errors.New("not admin"), Code: http.StatusUnauthorized, From: "GetAccount", Description: "Get account"}
}

func StartKeyDB(acc *users.Account, dev bool, start_web bool, port string, start_unix bool) ActionResponce {

	if acc.IsAdmin() {
		db_utils.LoadDB(internal.DB_filename)
		users.LoadAccounts(internal.Accounts_filename)

		persistance.Start_writers(conf.Number_of_writers)

		if start_web {
			StartRemote(acc, dev, port)
		}

		if start_unix {
			StartUnix(acc)
		}

		return ActionResponce{Error: nil, Code: http.StatusOK, From: "StartKeyDB", Description: "Start KeyDB"}
	}
	return ActionResponce{Error: errors.New("not admin"), Code: http.StatusUnauthorized, From: "StartKeyDB", Description: "Start KeyDB"}

}

func SetRouter(acc *users.Account) ActionResponce {
	if acc.IsAdmin() {
		router = gin.New()
		RouterSetupFunc(router)
		RouterAddEndpointsFunc(router)
		router_set = true

		return ActionResponce{Error: nil, Code: http.StatusOK, From: "setRouter", Description: "Set router"}
	}
	return ActionResponce{Error: errors.New("not admin"), Code: http.StatusUnauthorized, From: "setRouter", Description: "Set router"}
}

func StartRemote(acc *users.Account, dev bool, port string) ActionResponce {
	if acc.IsAdmin() {

		if !router_set {
			SetRouter(acc)
		}

		if dev {
			gin.SetMode(gin.DebugMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}

		go router.Run(":" + port)

		return ActionResponce{Error: nil, Code: http.StatusOK, From: "startRemote", Description: "Start http server"}
	}
	return ActionResponce{Error: errors.New("not admin"), Code: http.StatusUnauthorized, From: "startRemote", Description: "Start http server"}
}

func StartUnix(acc *users.Account) ActionResponce {
	if acc.IsAdmin() {
		if !router_set {
			SetRouter(acc)
		}

		listener, err := net.Listen("unix", "/tmp/keydb_sock.sock")
		if err != nil {
			panic(err)
		}

		go http.Serve(listener, router)

		return ActionResponce{Error: nil, Code: http.StatusOK, From: "startUnix", Description: "Start unix server"}
	}

	return ActionResponce{Error: errors.New("not admin"), Code: http.StatusUnauthorized, From: "startUnix", Description: "Start unix server"}
}

func StopWeb(acc *users.Account) ActionResponce {
	if acc.IsAdmin() {
		srv := &http.Server{
			Addr:    ":8080",
			Handler: router,
		}
		srv.Shutdown(context.Background())
		return ActionResponce{Error: nil, Code: http.StatusOK, From: "StopWeb", Description: "Stop http web server"}
	}
	return ActionResponce{Error: errors.New("not admin"), Code: http.StatusUnauthorized, From: "StopWeb", Description: "Stop http web server"}
}

func StopUnix() {}
