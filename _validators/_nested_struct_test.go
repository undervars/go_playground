package validators

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type A struct {
	B B `validate:"required"`
}

type B struct {
	C int `validate:"required"`
}

func TestValidateNestedStruct(t *testing.T) {
	a1 := A{
		B: B{},
	}

	a2 := A{
		B: B{
			C: 10,
		},
	}
	// assert.Equal(t, 1, 2)
	v := validator.New()

	// Key: 'A.B.C' Error:Field validation for 'C' failed on the 'required' tag
	assert.Error(t, v.Struct(a1))
	assert.NoError(t, v.Struct(a2))
}

type A2 struct {
	B2 *B2 `validate:"required"`
}

type B2 struct {
	C2 int `validate:"required"`
}

func TestValidateNestedStructPointer(t *testing.T) {
	a1 := A2{
		B2: &B2{},
	}

	a2 := A2{
		B2: &B2{
			C2: 10,
		},
	}
	// assert.Equal(t, 1, 2)
	v := validator.New()

	assert.NoError(t, v.Struct(a1))
	assert.NoError(t, v.Struct(a2))
}

type A3 struct {
	B3 B3 `validate:"required"`
}

type B3 struct {
	C3 int `validate:"gt=1"`
}

func TestGt(t *testing.T) {
	a1 := A3{
		B3: B3{
			C3: 1,
		},
	}

	a2 := A3{
		B3: B3{
			C3: 2,
		},
	}
	// assert.Equal(t, 1, 2)
	v := validator.New()

	// 	Key: 'A3.B3.C3' Error:Field validation for 'C3' failed on the 'gt' tag
	assert.Error(t, v.Struct(a1))
	assert.NoError(t, v.Struct(a2))
}
