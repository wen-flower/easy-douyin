package db

import (
	"github.com/wen-flower/easy-douyin/cmd/chat/cfg"
	"github.com/wen-flower/easy-douyin/pkg/mlog/gormlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	DB   *gorm.DB
	once sync.Once
)

// Init 初始化数据库连接
func Init() {
	once.Do(func() {
		db, err := gorm.Open(mysql.Open(cfg.MySqlDNS),
			&gorm.Config{
				PrepareStmt: true,
				Logger:      gormlog.GormLogger(cfg.LogJson, cfg.LogPretty),
			},
		)

		if err != nil {
			panic(err)
		}

		DB = db
	})
}
