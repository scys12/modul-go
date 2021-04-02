package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/scys12/modul-go/internal/config"
)

func InitDatabase(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open(cfg.DBDriver, dsn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * time.Duration(cfg.MaxLifetime))
	db.SetMaxOpenConns(cfg.OpenConn)
	db.SetMaxIdleConns(cfg.IdleConn)
	return db, nil
}
