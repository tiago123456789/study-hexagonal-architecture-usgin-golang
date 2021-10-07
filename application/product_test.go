package application_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/tiago123456789/study-hexagonal-architecture-usgin-golang/application"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{
		Name:   "Test product",
		Status: application.DISABLED,
		Price:  10,
	}

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "Product need have price greather tha 0 to enabled", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{
		Name:   "Test product",
		Status: application.ENABLED,
		Price:  0,
	}

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "Product must be disabled when the price equal 0", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{
		ID:     uuid.NewV4().String(),
		Name:   "Test product",
		Status: "test",
		Price:  0,
	}

	_, err := product.IsValid()
	require.Error(t, err)
	product.Status = application.DISABLED
	product.Price = 10
	_, err = product.IsValid()
	require.Nil(t, err)
}
