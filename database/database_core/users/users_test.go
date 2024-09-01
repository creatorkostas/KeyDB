package users_test

import (
	"os"
	"testing"

	"github.com/creatorkostas/KeyDB/database/database_core/users"
	"github.com/jedib0t/go-pretty/table"
)

func TestCreateUser(t *testing.T) {

	// Run_read_test(100)
	c_table := table.NewWriter()
	c_table.SetOutputMirror(os.Stdout)
	c_table.AppendHeader(table.Row{"acc"})

	var acc, _, err = users.Create_account("test", "Admin", "test@test.com", "123456789")
	acc.Api_key = "123456789"
	acc.Public_key = "key"

	var local_acc = users.MakeDefaultUser()
	var userInfo users.UserInfo = users.UserInfo{Username: "test", Api_key: "123456789", Email: "test@test.com", Password: "15e2b0d3c33891ebb0f1ef609ec419420c20e320ce94c65fbc8c3312448eb225"}
	local_acc.MakeAdmin()
	local_acc.UserInfo = userInfo
	local_acc.Public_key = "key"

	c_table.AppendRows([]table.Row{{acc}})
	c_table.AppendRows([]table.Row{{local_acc}})

	if *acc != local_acc {
		c_table.Render()
		t.Fatalf("[ERROR]: Creating a new account")
	}

	if err != nil {
		c_table.Render()
		t.Fatalf("[ERROR]: " + err.Error())
	}

	c_table.Render()
}
