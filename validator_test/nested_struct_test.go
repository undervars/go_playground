package validator_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type User struct {
	bankAccount *BankAccount
}

type BankAccount struct {
	Money uint `validate:"gte=1000"`
}

func TestValidateNestedStruct(t *testing.T) {
	u1 := User{
		bankAccount: &BankAccount{
			Money: 10,
		},
	}

	u2 := User{
		bankAccount: &BankAccount{
			Money: 1000,
		},
	}

	v := validator.New()

	assert.NoError(t, v.Struct(u1))
	assert.NoError(t, v.Struct(u2))
}
