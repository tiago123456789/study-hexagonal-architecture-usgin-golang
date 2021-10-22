package application

import (
	uuid "github.com/satori/go.uuid"
)

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{
		Persistence: persistence,
	}
}

func (p *ProductService) Get(id string) (ProductInterface, error) {
	product, err := p.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductService) Create(name string, price float64) (ProductInterface, error) {
	productToService := &Product{
		Name:   name,
		Price:  price,
		Status: ENABLED,
		ID:     uuid.NewV4().String(),
	}

	_, err := productToService.IsValid()
	if err != nil {
		return nil, err
	}

	productReturned, err := p.Persistence.Save(productToService)
	if err != nil {
		return nil, err
	}

	return productReturned, nil

}
