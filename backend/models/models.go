package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/telepathy/telepathy-admin/backend/pkg/setting"
)

var db *sql.DB

func Setup() {
	var err error
	db, err = sql.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf(err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
}
