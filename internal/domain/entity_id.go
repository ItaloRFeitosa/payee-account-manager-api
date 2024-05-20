package domain

import (
	"github.com/google/uuid"
)

type EntityID struct {
	value string
}

func NewEntityID() EntityID {
	return EntityID{uuid.NewString()}
}

func (e EntityID) Value() string {
	return e.value
}
