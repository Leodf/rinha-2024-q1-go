package utils

import (
	"errors"

	"github.com/Leodf/leodf-go/internal/dto"
	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

func IdValidator(id int) error {
	if id < 1 || id > 5 {
		return errors.New("not found")
	}
	return nil
}

func TransactionReqValidator(body *dto.TransactionRequest) error {
	err := Validator.Struct(body)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
