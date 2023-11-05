package database

import (
	"context"
	database "database/sql"
)

type DBTransaction struct {
	DB *database.DB
	TX *database.Tx
}

func NewTransaction() (tx *DBTransaction, err error) {
	if tx.DB, err = database.Open("postgre", "miniature-potato"); err != nil {
		return
	}

	opts := new(database.TxOptions)
	tx1, err := tx.DB.BeginTx(context.Background(), opts)
	if err != nil {
		return
	}

	tx.TX = tx1
	return
}

func (tx *DBTransaction) Rollback() {
	tx.TX.Rollback()
}

func (tx *DBTransaction) Commit() {
	tx.TX.Commit()
}
