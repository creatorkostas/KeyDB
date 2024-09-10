package security_test

import (
	"testing"

	"github.com/creatorkostas/KeyDB/database/database_core/security"
	"github.com/creatorkostas/KeyDB/database/database_core/users"
)

var data []byte
var message = "test"
var acc *users.Account
var p_key string
var err error

func TestEncrypt(t *testing.T) {
	t.SkipNow()
	acc, p_key, err = users.Create_account("test", "Admin", "test@test.com", "test")
	if err != nil || acc == nil || p_key == "" {
		t.Fatalf("Security register failed")
	}
	data = security.Encrypt_data(p_key, []byte(message))
	if len(data) == 0 {
		t.Fatalf("Security encrypt failed")
	}
	err = nil
}

func TestDecrypt(t *testing.T) {
	t.SkipNow()
	var acc, p_key, err = users.Create_account("test", "Admin", "test@test.com", "test")
	if err != nil || acc == nil || p_key == "" {
		t.Fatalf("Security register failed")
	}
	var str_data = security.Decrypt_data(acc.Public_key, data)
	if str_data != message {
		t.Fatalf("Security decrypt failed")
	}
}
