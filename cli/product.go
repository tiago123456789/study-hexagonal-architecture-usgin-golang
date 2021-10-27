package cli

import (
	"fmt"

	"github.com/tiago123456789/study-hexagonal-architecture-usgin-golang/application"
)

func Run(
	service application.ProductServiceInterface, action string,
	productId string, productName string, price float64) (string, error) {

	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(),
			product.GetStatus())
	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	}

	return result, nil
}
