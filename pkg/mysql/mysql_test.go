package mysql_test

import (
	"os"
	"testing"

	"github.com/luckyss6/go-kit/pkg/config"
	"github.com/luckyss6/go-kit/pkg/mysql"
)

func TestMain(m *testing.M) {
	os.Setenv("ENV_PATH", "../../")
	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	cfg := config.New()
	db := mysql.New(cfg)
	if db == nil {
		t.Errorf("db is nil")
		return
	}
	t.Log("connect successfully")
}
