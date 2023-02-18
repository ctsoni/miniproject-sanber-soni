package helper

import "github.com/go-playground/validator/v10"

func FormatThickness(page int) string {
	var thickness string

	switch {
	case page <= 100:
		thickness = "tipis"
	case page <= 200 && page >= 101:
		thickness = "sedang"
	case page >= 201:
		thickness = "tebal"
	}

	return thickness
}

func FormatError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
