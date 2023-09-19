package errors

import (
	"space-traders-playground/pkg/dom"
)

var (
	ErrInvalidFaction    = dom.NewError(100, "invalid faction")
	ErrFailedInsertMongo = dom.NewError(101, "failed to insert to MongoDB")
	ErrFailedAPIRequest  = dom.NewError(102, "failed to request API")
)
