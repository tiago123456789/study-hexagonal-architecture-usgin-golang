package application

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetPrice() float64
	GetStatus() string
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string
	Name   string
	Price  float64
	Status string
}

// func (p *Product) IsValid() (bool, error) {

// }

// func (p *Product) Enable() error {}

// func (p *Product) Disable() error {}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetStatus() string {
	return p.Status
}
