package validator

import (
	"fmt"
	"regexp"
	"strings"

	validatorError "github.com/Brotiger/per-painted_poker-backend/internal/validator/error"
	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate
var reArr *regexp.Regexp

func init() {
	Validator = validator.New()
	reArr = regexp.MustCompile(`(.+)\[(\d+)\]*`)
}

func ValidateErr(err error) map[string][]string {
	fieldErrors := map[string][]string{}

	for _, err := range err.(validator.ValidationErrors) {
		textError := validatorError.DefaultError

		if objError, ok := validatorError.Map[err.Tag()]; ok {
			textError = objError.GetError(err.Param())
		}

		field := fmt.Sprintf("%s%s", strings.ToLower(string(err.Field()[0])), err.Field()[1:])
		match := reArr.FindStringSubmatch(field)
		if len(match) == 3 {
			field = fmt.Sprintf("%s.%s", match[1], match[2])
		}

		fieldErrors[field] = append(fieldErrors[field], textError)
	}

	return fieldErrors
}
