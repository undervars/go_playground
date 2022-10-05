package validators

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type RequiredFields struct {
	IntType           int     `validate:"required"`
	StringType        string  `validate:"required"`
	StringPointerType *string `validate:"required"`
}

func TestRequiredFields(t *testing.T) {
	v := validator.New()

	r1 := RequiredFields{}
	r2 := RequiredFields{
		IntType:           0,
		StringType:        "",
		StringPointerType: nil,
	}

	str := ""
	r3 := RequiredFields{
		IntType:           1,
		StringType:        "1",
		StringPointerType: &str,
	}

	assert.Error(t, v.Struct(r1))
	assert.Error(t, v.Struct(r2))

	assert.NoError(t, v.Struct(r3))

}
