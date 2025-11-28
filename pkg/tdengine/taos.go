package tdengine

import (
	"fmt"

	"gitee.com/chunanyong/zorm"
)

type Config struct {
	Protocal              string `json:"protocal"`
	Host                  string `json:"host" yaml:"host"`
	Port                  string `json:"port" yaml:"port"`
	Username              string `json:"username" yaml:"username"`
	Password              string `json:"password" yaml:"password"`
	DBName                string `json:"db_name" yaml:"db_name"`
	MaxOpenConns          int    `json:"max_open_conns" yaml:"max_open_conns"`
	ConnMaxLifeTimeSecond int    `json:"conn_max_life_time_second" yaml:"conn_max_life_time_second"`
	MaxIdleConns          int    `json:"max_idle_conns" yaml:"max_idle_conns"`
	SlowSQLMillis         int    `json:"slow_sql_millis" yaml:"slow_sql_millis"` // default: 1000ms
}

type Taos struct {
	*zorm.DBDao
}

func NewTaos(cfg Config) *Taos {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Protocal, cfg.Host, cfg.Password, cfg.DBName)
	driverName := "taosSql"
	if cfg.Protocal == "http" {
		driverName = "taosRestful"
	}
	dbDaoConfig := zorm.DataSourceConfig{
		DSN:                   dsn,
		DriverName:            driverName,
		Dialect:               "tdengine",
		MaxOpenConns:          cfg.MaxOpenConns,
		MaxIdleConns:          cfg.MaxIdleConns,
		ConnMaxLifetimeSecond: cfg.ConnMaxLifeTimeSecond,
		DisableTransaction:    true,
		SlowSQLMillis:         cfg.SlowSQLMillis,
	}
	db, err := zorm.NewDBDao(&dbDaoConfig)
	if err != nil {
		panic(err)
	}
	return &Taos{db}
}

func (t *Taos) Close() error {
	return t.DBDao.CloseDB()
}
