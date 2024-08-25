package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/JordanMarcelino/drivers-backend/internal/config"
	"github.com/JordanMarcelino/drivers-backend/internal/pkg/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitStdLib(cfg *config.Config) *sql.DB {
	dbCfg := cfg.Database

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Jakarta",
		dbCfg.Host,
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.DbName,
		dbCfg.Port,
		dbCfg.Sslmode,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Log.Fatal("error initializing database: ", err.Error())
	}

	err = db.Ping()
	if err != nil {
		logger.Log.Fatal("error connecting to database; ", err.Error())
	}

	db.SetMaxIdleConns(dbCfg.MaxIdleConn)
	db.SetMaxOpenConns(dbCfg.MaxOpenConn)
	db.SetConnMaxLifetime(time.Duration(dbCfg.MaxConnLifetimeMinute) * time.Minute)

	return db
}
