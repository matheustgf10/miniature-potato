package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	domain "github.com/matheustgf10/miniature-potato/domain/crm/order"
)

type PGOrder struct {
	DB      *sql.DB
	Builder squirrel.StatementBuilderType
}

func (pg *PGOrder) ListOrder() (res *domain.ResOrder, err error) {
	consulta := pg.Builder.
		Select("*").
		From("tabela1").
		Where("id = 2")

	query, values, err := consulta.ToSql()
	if err != nil {
		return
	}

	rows, err := pg.DB.Query(query, values)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, nil
		}
	}

	for rows.Next() {
		var order *domain.Order
		if err = rows.Scan(&order.ID, &order.Title, &order.Type); err != nil {
			return
		}

		res.Data = append(res.Data, order)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	if err = rows.Close(); err != nil {
		return
	}

	return
}
