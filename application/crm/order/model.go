package order

import "github.com/google/uuid"

type Order struct {
	ID    *uuid.UUID `json:"id"`
	Title *string    `json:"title"`
	Type  *string    `json:"type"`
}

type ResOrder struct {
	Data  []*Order `json:"data"`
	Next  *bool    `json:"next"`
	Total *int64   `json:"total"`
}
