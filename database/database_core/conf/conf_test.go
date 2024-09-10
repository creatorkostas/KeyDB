package conf_test

import (
	"testing"

	"github.com/creatorkostas/KeyDB/database/database_core/conf"
)

func TestStartWeb(t *testing.T) {
	t.SkipNow()
	conf.Load_configs("conf.yaml")
	if conf.DB_filename != "db.gob" ||
		conf.Accounts_filename != "accounts.gob" ||
		conf.Send_all_errors_in_requests != true ||
		conf.Append_only_in_file != false ||
		conf.Append_file != "aof.op" ||
		conf.StartWeb != false ||
		conf.StartUnix != false ||
		conf.WebPort != "8080" ||
		conf.Number_of_writers != 1 {
		t.Fatalf("Load_configs failed")
	}
}
