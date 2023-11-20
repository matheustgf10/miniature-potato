package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

var DB *sql.DB

func Open() {
	var (
		HOST     = "localhost"
		PORT     = "5432"
		USER     = "root"
		PASSWORD = "1234@mudar"
		DBNAME   = "postgres"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		HOST,
		PORT,
		USER,
		PASSWORD,
		DBNAME)

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
