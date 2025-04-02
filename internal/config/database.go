package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

func DatabaseProvider(lc fx.Lifecycle) *sqlx.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}

	lc.Append(fx.Hook{
		OnStop: func(c context.Context) error {
			if err := db.Close(); err != nil {
				return err
			}
			return nil
		},
	})

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)

	return db
}
