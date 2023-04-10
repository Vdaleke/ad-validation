package validator

import "github.com/pkg/errors"

var ValidationError = errors.New("validation error")

func ValidateAd(title string, text string) error {
	if title == "" || text == "" || len(title) > 100 || len(text) > 100 {
		return ValidationError
	}
	return nil
}
