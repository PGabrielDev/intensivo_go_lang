package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_if_Id_is_null(t *testing.T) {
	order := Order{}
	err := order.Validate()
	if err == nil {
		t.Error("Error Expected")
	}
}

func Test_if_Id_is_nulll(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "invalid id")
}

func Test_If_price_equals_zero(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "Invalid Price")
}

func Test_If_tax_equals_zero(t *testing.T) {
	order := Order{ID: "123", Price: 20.0}
	assert.Error(t, order.Validate(), "Invalid Tax")
}

func Test_all_valid_params(t *testing.T) {
	order := Order{ID: "123", Price: 20.0, Tax: 1.0}
	assert.NoError(t, order.Validate())
	assert.NoError(t, order.CalculateFinalPrice())
	assert.Equal(t, 21.0, order.FinalPrice)
	assert.Equal(t, 20.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
}
