package order

import "github.com/google/uuid"

type Order struct {
	ID    *uuid.UUID
	Title *string
	Type  *string
}

type ResOrder struct {
	Data  []*Order
	Next  *bool
	Total *int64
}
