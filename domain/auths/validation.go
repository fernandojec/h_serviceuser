package auths

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStructAuth(user interface{}) []string {
	// var errors []ErrorResponse
	err := validate.Struct(user)
	errMsgs := make([]string, 0)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.Field(),
				err.Param(),
				err.Tag(),
			))
		}
	}
	return errMsgs
}
