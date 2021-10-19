package application_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tiago123456789/study-hexagonal-architecture-usgin-golang/application"
	mock_application "github.com/tiago123456789/study-hexagonal-architecture-usgin-golang/application/mocks"
)

func TestProduct_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).Times(1)

	productService := application.ProductService{
		Persistence: persistence,
	}

	_, err := productService.Get("1")
	require.Nil(t, err)
	persistence.EXPECT().Get(gomock.Any()).Return(nil, errors.New("Error")).Times(1)
	_, err = productService.Get("1")
	require.Error(t, err)
}

func TestProduct_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().IsValid().Return(false, errors.New("")).AnyTimes()

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(nil, errors.New("")).AnyTimes()
	productService := application.ProductService{
		Persistence: persistence,
	}

	_, err := productService.Create("ps5", 5000)
	require.Error(t, err)

	product.EXPECT().IsValid().Return(true, nil).AnyTimes()
	_, err = productService.Create("ps5", 5000)
	require.Error(t, err)

}
