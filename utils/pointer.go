package utils

import "github.com/google/uuid"

func GetStringPointer(s string) *string {
	return &s
}

func GetUUIDPointer(id uuid.UUID) *uuid.UUID {
	return &id
}
