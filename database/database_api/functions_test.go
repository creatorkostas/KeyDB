package api_test

import (
	"net/http"
	"testing"

	api "github.com/creatorkostas/KeyDB/database/database_api"
	"github.com/creatorkostas/KeyDB/database/database_core/users"
	"github.com/creatorkostas/KeyDB/internal/tools"
)

// GetValue(key , acc )
// SetValues(key , acc , value_type , data , encrypt )
// Register(username , email , password , acc_type ) (, , error)
// ChangeApiKey(acc )
// ChangePassword(acc , new_pass )
// Save(acc )
// Load(acc )
// GetAccount(acc , username )
// StartKeyDB(acc , dev , start_web , port , start_unix )
// SetRouter(acc )
// StartRemote(acc , dev , port )
// StartUnix(acc )
// StopWeb(acc )
// StopUnix()

var local_acc users.Account = users.MakeDefaultUser()

func setTestUser() {
	local_acc.MakeAdmin()
	local_acc.Username = "Local cmd admin account"
}

func TestRegister(t *testing.T) {

	var acc, private_key, err = api.Register("test", "test@test.com", "test", "Admin")
	if err != nil || acc == nil || private_key == "" {
		t.Fatalf("Admin register failed")
	}
	if acc.UserInfo.Username != "test" || acc.UserInfo.Email != "test@test.com" || acc.UserInfo.Password != "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08" || acc.Tier != users.Admin_tier {
		t.Fatalf("Admin wrong account data")
	}

	var acc1, private_key1, err1 = api.Register("test1", "test@test.com", "test", "User")
	if err1 != nil || acc1 == nil || private_key1 == "" {
		t.Fatalf("User register failed")
	}
	if acc1.UserInfo.Username != "test1" || acc1.UserInfo.Email != "test@test.com" || acc1.UserInfo.Password != "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08" || acc1.Tier != users.User_tier {
		t.Fatalf("User wrong account data")
	}
	var acc2, private_key2, err2 = api.Register("test2", "test@test.com", "test", "GuestUser")
	if err2 != nil || acc2 == nil || private_key2 == "" {
		t.Fatalf("Guest register failed")
	}
	if acc2.UserInfo.Username != "test2" || acc2.UserInfo.Email != "test@test.com" || acc2.UserInfo.Password != "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08" || acc2.Tier != users.Guest_user_tier {
		t.Fatalf("Guest wrong account data")
	}

}
func TestGetAccount(t *testing.T) {
	setTestUser()
	var acc = (api.GetAccount(&local_acc, "test")).Description.(*users.Account)
	if acc.UserInfo.Username != "test" || acc.UserInfo.Email != "test@test.com" || acc.UserInfo.Password != "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08" || acc.Tier != users.Admin_tier {
		t.Fatalf("Wrong account data")
	}
}
func TestSetValues(t *testing.T) {

	setTestUser()
	var acc = (api.GetAccount(&local_acc, "test")).Description.(*users.Account)
	var data = api.SetValues(acc, "test1_int", "int", "2", false)
	if data.Error != nil || data.Code != http.StatusOK || data.From != "SetValues" || data.Description.(bool) != true {
		t.Fatalf("SetValues int failed")
	}

	data = api.SetValues(acc, "test1_string", "string", "some string", false)
	if data.Error != nil || data.Code != http.StatusOK || data.From != "SetValues" || data.Description.(bool) != true {
		t.Fatalf("SetValues string failed")
	}

	data = api.SetValues(acc, "test1_bool_true", "bool", "1", false)
	if data.Error != nil || data.Code != http.StatusOK || data.From != "SetValues" || data.Description.(bool) != true {
		t.Fatalf("SetValues bool truefailed")
	}

	data = api.SetValues(acc, "test1_bool_false", "bool", "false", false)
	if data.Error != nil || data.Code != http.StatusOK || data.From != "SetValues" || data.Description.(bool) != true {
		t.Fatalf("SetValues bool false failed")
	}

	data = api.SetValues(acc, "test1_float32", "float32", "21.23", false)
	if data.Error != nil || data.Code != http.StatusOK || data.From != "SetValues" || data.Description.(bool) != true {
		t.Fatalf("SetValues float32failed")
	}

	data = api.SetValues(acc, "test1_float64", "float64", "21.234", false)
	if data.Error != nil || data.Code != http.StatusOK || data.From != "SetValues" || data.Description.(bool) != true {
		t.Fatalf("SetValues float64 failed")
	}
}

func TestGetValue(t *testing.T) {
	setTestUser()
	var acc = (api.GetAccount(&local_acc, "test")).Description.(*users.Account)

	var data = api.GetValue(acc, "test1_int")
	if data.Error != nil || data.Code != http.StatusOK || data.From != "GetValue" || data.Description.(int64) != 2 {
		t.Fatalf("GetValue int failed")
	}

	data = api.GetValue(acc, "test1_string")
	if data.Error != nil || data.Code != http.StatusOK || data.From != "GetValue" || data.Description.(string) != "some string" {
		t.Fatalf("GetValue string failed")
	}

	data = api.GetValue(acc, "test1_bool_true")
	if data.Error != nil || data.Code != http.StatusOK || data.From != "GetValue" || data.Description.(bool) != true {
		t.Fatalf("GetValue bool true failed")
	}

	data = api.GetValue(acc, "test1_bool_false")
	if data.Error != nil || data.Code != http.StatusOK || data.From != "GetValue" || data.Description.(bool) != false {
		t.Log(data.Description)
		t.Fatalf("GetValue bool false failed")
	}

	data = api.GetValue(acc, "test1_float32")
	if data.Error != nil || data.Code != http.StatusOK || data.From != "GetValue" || data.Description.(float32) != 21.23 {
		t.Fatalf("GetValue float32 failed")
	}

	data = api.GetValue(acc, "test1_float64")
	if data.Error != nil || data.Code != http.StatusOK || data.From != "GetValue" || data.Description.(float64) != 21.234 {
		t.Fatalf("GetValue float64 failed")
	}

}

func TestChangeApiKey(t *testing.T) {
	setTestUser()
	var acc = (api.GetAccount(&local_acc, "test")).Description.(*users.Account)
	var current_api_key = acc.UserInfo.Api_key
	acc.ChangeApiKey()
	if acc.UserInfo.Api_key == current_api_key {
		t.Fatalf("ChangeApiKey failed")
	}
}

func TestChangePassword(t *testing.T) {
	setTestUser()
	var acc = (api.GetAccount(&local_acc, "test")).Description.(*users.Account)
	var current_password = acc.UserInfo.Password
	acc.ChangePassword("new password")
	if acc.UserInfo.Password == current_password {
		t.Fatalf("ChangePassword failed")
	}
}

func TestSave(t *testing.T) {
	var resp = api.Save(&local_acc)
	if resp.Error != nil || resp.Code != http.StatusOK || resp.From != "Save" {
		t.Fatalf("Save failed")
	}
}

func TestLoad(t *testing.T) {
	var resp = api.Load(&local_acc)
	if resp.Error != nil || resp.Code != http.StatusOK || resp.From != "Load" {
		t.Fatalf("Load failed")
	}
}

func TestStartKeyDB(t *testing.T) {
	var resp = api.StartKeyDB(&local_acc, true, false, "8080", false)
	if resp.Error != nil || resp.Code != http.StatusOK || resp.From != "StartKeyDB" {
		t.Fatalf("StartKeyDB failed")
	}
}

func TestStartWeb(t *testing.T) {
	t.SkipNow()
	var resp = api.StartWeb(&local_acc, true, "8080")
	api.StopWeb(&local_acc)
	if resp.Error != nil || resp.Code != http.StatusOK || resp.From != "StartRemote" {
		t.Fatalf("StartRemote failed")
	}
}

func TestStartUnix(t *testing.T) {
	t.SkipNow()
	var resp = api.StartUnix(&local_acc)
	if resp.Error != nil || resp.Code != http.StatusOK || resp.From != "StartUnix" {
		t.Fatalf("StartUnix failed")
	}
}

func TestStopWeb(t *testing.T) {
	t.SkipNow()
	var resp = api.StopWeb(&local_acc)
	if resp.Error != nil || resp.Code != http.StatusOK || resp.From != "StopWeb" {
		t.Fatalf("StopWeb failed")
	}
}

func TestStopUnix(t *testing.T) {
	t.SkipNow()
	api.StopUnix()
}

func TestCleanUp(t *testing.T) {
	tools.DeleteFile("db.gob")
	tools.DeleteFile("accounts.gob")
}
