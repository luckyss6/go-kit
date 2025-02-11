package config_test

import (
	"os"
	"testing"

	"github.com/luckyss6/go-kit/pkg/config"
)

func TestMain(m *testing.M) {
	os.Setenv("ENV_PATH", "../../")
	os.Exit(m.Run())
}
func TestReadConfig(t *testing.T) {
	cfg := config.New()

	if cfg == nil {
		t.Errorf("cfg is nil")
		return
	}
}

func TestReadPostgres(t *testing.T) {
	cfg := config.New()

	if cfg.Postgres.Host == "" {
		t.Errorf("host is empty")
		return
	}

	if cfg.Postgres.Port == 0 {
		t.Errorf("port is 0")
		return
	}

	if cfg.Postgres.User == "" {
		t.Errorf("user is empty")
		return
	}

	if cfg.Postgres.Password == "" {
		t.Errorf("password is empty")
		return
	}

	if cfg.Postgres.DBname == "" {
		t.Errorf("dbname is empty")
		return
	}
}

func TestReadRedis(t *testing.T) {
	os.Setenv("ENV_PATH", "../")
	cfg := config.New()

	if cfg.Redis.Addr == "" {
		t.Errorf("addr is empty")
		return
	}

	if cfg.Redis.Password == "" {
		t.Errorf("password is empty")
		return
	}

	if cfg.Redis.DB == 0 {
		t.Errorf("db is 0")
		return
	}
}
