package utils

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ErrResponse(err error) []*ErrorResponse {
	var errs []*ErrorResponse
	for _, e := range err.(validator.ValidationErrors) {
		errs = append(errs, &ErrorResponse{
			FailedField: e.StructNamespace(),
			Tag:         e.Tag(),
			Value:       e.Param(),
		})
	}
	return errs
}
