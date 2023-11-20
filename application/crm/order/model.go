package order

import "github.com/google/uuid"

type Order struct {
	ID    *uuid.UUID `json:"id,omitempty"`
	Title *string    `json:"title,omitempty"`
	Type  *string    `json:"type,omitempty"`
}

type ResOrder struct {
	Data  []*Order `json:"data,omitempty"`
	Next  *bool    `json:"next,omitempty"`
	Total *int64   `json:"total,omitempty"`
}
