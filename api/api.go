package api

import (
	"net/http"

	"github.com/creatorkostas/KeyDB/internal"
	"github.com/creatorkostas/KeyDB/internal/handlers"
	"github.com/creatorkostas/KeyDB/internal/tools"
	"github.com/gin-gonic/gin"
)

type Responce struct {
	Message string `json:"response"`
}

func GetValue(c *gin.Context) {

	var key, _ = c.GetQuery("key")
	var user = c.Param("user")

	var result = handlers.Get_value(user, key)
	if result == nil {
		c.IndentedJSON(http.StatusBadRequest, "Key does not exist")
	} else {
		c.IndentedJSON(http.StatusOK, result)
	}

}

func GetAll(c *gin.Context) {

	var user = c.Param("user")

	var result = handlers.Get_Users_Data(user)
	if result == nil {
		c.IndentedJSON(http.StatusBadRequest, "Key does not exist")
	} else {
		c.IndentedJSON(http.StatusOK, result)
	}

}

// func SetValues(key string, value_type string, value interface{}) {
func SetValues(c *gin.Context) {

	var key, _ = c.GetQuery("key")

	var value_type, _ = c.GetQuery("type")

	var value, _ = c.GetQuery("value")
	var user = c.Param("user")

	var ok = handlers.Add_value(user, key, value_type, value)
	// var val = handlers.Get_value(user, key)

	// log.Println("val: ", val)
	if ok {
		c.IndentedJSON(http.StatusOK, Responce{"ok"})
	} else {
		c.IndentedJSON(http.StatusOK, Responce{"Error"})
	}
}

func Register(c *gin.Context) {
	// Create_account(username string, acc_tier string, email string, password string)

	var username, _ = c.GetQuery("username")
	// var acc_tier, _ = c.GetQuery("acc_tier")
	var email, _ = c.GetQuery("email")
	var password, _ = c.GetQuery("password")

	var acc = handlers.Create_account(username, "Admin", email, password)

	if acc == nil {
		c.IndentedJSON(http.StatusBadRequest, Responce{Message: "User already exist"})
		return
	}
	c.IndentedJSON(http.StatusOK, acc)
}

func Save(c *gin.Context) {
	tools.SaveDB(internal.DB_filename, &handlers.DB)
	tools.SaveDB(internal.Accounts_filename, &handlers.Accounts)

	c.IndentedJSON(http.StatusOK, Responce{"ok"})
}

func Load(c *gin.Context) {
	tools.LoadDB(internal.DB_filename, &handlers.DB)
	tools.LoadDB(internal.Accounts_filename, &handlers.Accounts)

	c.IndentedJSON(http.StatusOK, Responce{"ok"})
}
