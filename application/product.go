package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

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
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float"`
	Status string  `valid:"required"`
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("The status must be enabled or disabled")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("Product need have price greather tha 0 to enabled")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("Product must be disabled when the price equal 0")
}

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
