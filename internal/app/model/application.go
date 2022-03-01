package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Application struct {
	ID         int    `json:"id"`
	LoanAmount string `json:"loanAmount"`
	LoanLength string `json:"loanLength"`
}

func (application *Application) Validate() error {
	return validation.ValidateStruct(
		application,
		validation.Field(&application.ID, validation.Required, is.Int),
		validation.Field(&application.LoanAmount, validation.Required, validation.Length(100, 36000)),
		validation.Field(&application.LoanLength, validation.Required, validation.In(4, 6, 12, 24, 36)),
	)
}
