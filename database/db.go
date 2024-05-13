package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

var (
	dbName   = viper.GetString("DB_NAME")
	dbHost   = viper.GetString("DB_HOST")
	dbPass   = viper.GetString("DB_PASSWORD")
	dbUser   = viper.GetString("DB_USERNAME")
	dbPort   = viper.GetString("DB_PORT")
	dbParams = viper.GetString("DB_PARAMS")
)

func GetConnPool() *pgxpool.Pool {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", dbUser, dbPass, dbHost, dbPort, dbName, dbParams)
	cfg, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatal("Error when parsing config DB URL: ", err)
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Fatal("Error when creating Database Pool Context: ", err)
	}

	return dbPool
}
