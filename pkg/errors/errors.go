package errors

import (
	"space-traders-playground/pkg/dom"
)

var (
	ErrInvalidFaction = dom.NewError(100, "invalid faction")
)