package postgres_test

import (
	"os"
	"testing"

	"github.com/luckyss6/go-kit/config"
	"github.com/luckyss6/go-kit/postgres"
)

var cfg *config.Config

func TestMain(m *testing.M) {
	os.Setenv("ENV_PATH", "../")
	cfg = config.New()

	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	db := postgres.New(cfg)
	if db == nil {
		t.Errorf("db is nil")
		return
	}
}

func TestReadDatabase(t *testing.T) {
	db := postgres.New(cfg)
	err := db.Exec("show database").Error
	if err != nil {
		t.Errorf("exec show databases err: %s", err)
		return
	}
}
