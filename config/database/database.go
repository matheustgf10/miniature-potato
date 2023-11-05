package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

var DB *sql.DB

func Open() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("NAME"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Could not connect the database: %s", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not connect the database: %s", err.Error())
	}

	DB = db
	log.Print("Database Connected")
}

func Close() {
	if DB != nil {
		_ = DB.Close()
	}
}

func NewTx() (*sql.Tx, error) {
	var (
		tx  *sql.Tx
		err error
	)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-time.After(time.Duration(5) * time.Second)
		if tx == nil {
			cancel()
		}
	}()

	tx, err = DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
