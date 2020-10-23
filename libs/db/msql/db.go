package msql

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

type Conf struct {
	Addr string        `yaml:"addr"` // 连接地址
	Open int           `yaml:"open"` // 最大打开连接数
	Idle int           `yaml:"idle"` // 最大闲置连接数
	Life time.Duration `yaml:"life"` // 连接最大生存时长
}

func NewDB(cfg Conf) *DB {
	db, err := sqlx.Open("mysql", cfg.Addr)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(cfg.Open)
	db.SetMaxIdleConns(cfg.Idle)
	db.SetConnMaxLifetime(cfg.Life)
	return &DB{DB: db}
}
