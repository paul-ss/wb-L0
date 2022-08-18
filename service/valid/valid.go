package valid

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

var (
	v *validator.Validate
)

func V() *validator.Validate {
	if v != nil {
		return v
	}

	v = validator.New()
	return v
}

func ErrorMsg(err error) string {
	errors, ok := err.(validator.ValidationErrors)
	if !ok {
		return ""
	}

	sb := strings.Builder{}
	for i, e := range errors {
		sb.WriteString(fmt.Sprintf("#%d (field = %s; structField = %s; tag = %s)  ",
			i+1, e.Field(), e.StructField(), e.Tag()))
	}

	return sb.String()
}
