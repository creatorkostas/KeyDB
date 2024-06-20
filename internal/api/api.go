package api

import (
	"net/http"

	"github.com/creatorkostas/KeyDB/internal"
	"github.com/creatorkostas/KeyDB/internal/database"
	"github.com/creatorkostas/KeyDB/internal/tools"
	"github.com/creatorkostas/KeyDB/internal/users"
	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"
)

func GetValue(c *gin.Context) {

	var key, found = c.GetQuery("key")

	var acc = c.MustGet("Account").(*users.Account)

	var result any
	if found {
		result = database.Get_value(acc.Username, key)
	} else {
		result = database.Get_value(acc.Username, "user.get.all.data")
	}

	var res = &Responce{C: c, ErrorMessage: "Key does not exist", Result: result, OkCode: http.StatusOK, ErrorCode: http.StatusBadRequest}
	sendResponce(res)

}

// func SetValues(key string, value_type string, value interface{}) {
func SetValues(c *gin.Context) {

	var acc = c.MustGet("Account").(*users.Account)

	var key, _ = c.GetQuery("key")

	var value_type, _ = c.GetQuery("type")

	// var value, _ = c.GetQuery("value")

	var ok = database.Add_value(acc.Username, key, value_type, c)

	var res = &Responce{C: c, ErrorMessage: "Error", Result: ok, OkCode: http.StatusOK, ErrorCode: http.StatusInternalServerError}
	sendResponce(res)

}

func Register(c *gin.Context) {

	var username, _ = c.GetQuery("username")
	// var acc_tier, _ = c.GetQuery("acc_tier")
	var email, _ = c.GetQuery("email")
	var password, _ = c.GetQuery("password")

	var acc = users.Create_account(username, "Admin", email, password)

	var res = &Responce{C: c, ErrorMessage: "Key does not exist", Result: acc, OkCode: http.StatusOK, ErrorCode: http.StatusBadRequest}
	sendResponce(res)
}

func GetStats(c *gin.Context) {
	c.JSON(http.StatusOK, stats.Report())
	c.Request.Context().Done()
}

func Save(c *gin.Context) {
	tools.SaveToFile(internal.DB_filename, &database.DB)
	tools.SaveToFile(internal.Accounts_filename, &users.Accounts)

	c.JSON(http.StatusOK, JsonResponce{"ok"})
}

func Load(c *gin.Context) {
	tools.LoadFromFile(internal.DB_filename, &database.DB)
	tools.LoadFromFile(internal.Accounts_filename, &users.Accounts)

	c.JSON(http.StatusOK, JsonResponce{"ok"})
}
