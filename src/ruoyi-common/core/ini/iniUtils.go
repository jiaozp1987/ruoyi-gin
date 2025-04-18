package ini

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var (
	JWT_DURATION int
	JWT_ISSUER   string
	JWT_KEY      string

	MYSQL_DSN               string
	MYSQL_MAX_IDLE_CONNS    int
	MYSQL_MAX_OPEN_CONNS    int
	MYSQL_CONN_MAX_LIFETIME int

	CRYPTOR_KEY string

	REDIS_ADDR     string
	REDIS_PASSWORD string
	REDIS_DB       int
	REDIS_EXPIRE   int
)

func init() {
	cfg := func() *ini.File {
		cfg, err := ini.Load("env.ini")
		if err != nil {
			fmt.Printf("Fail to read file: %v", err)
			os.Exit(1)
		}
		//r := cfg.Section(section).Key(key).String()
		return cfg
	}()

	JWT_DURATION, _ = cfg.Section("jwt").Key("duration").Int()
	JWT_ISSUER = cfg.Section("jwt").Key("issuer").String()
	JWT_KEY = cfg.Section("jwt").Key("signingKey").String()

	MYSQL_DSN = cfg.Section("mysql").Key("dsn").String()
	MYSQL_MAX_IDLE_CONNS, _ = cfg.Section("mysql").Key("maxIdleConns").Int()
	MYSQL_MAX_OPEN_CONNS, _ = cfg.Section("mysql").Key("maxOpenConns").Int()
	MYSQL_CONN_MAX_LIFETIME, _ = cfg.Section("mysql").Key("connMaxLifetime").Int()

	CRYPTOR_KEY = cfg.Section("cryptor").Key("cryptorKey").String()

	REDIS_ADDR = cfg.Section("redis").Key("addr").String()
	REDIS_PASSWORD = cfg.Section("redis").Key("password").String()
	REDIS_DB, _ = cfg.Section("redis").Key("db").Int()
	REDIS_EXPIRE, _ = cfg.Section("redis").Key("duration").Int()
}
