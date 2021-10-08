package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (p *ProductService) Get(id string) (ProductInterface, error) {
	product, err := p.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
