package db

import (
	"database/sql"
	"fmt"

	"github.com/tiago123456789/study-hexagonal-architecture-usgin-golang/application"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where id = $1")
	if err != nil {
		return nil, err
	}
	fmt.Printf("ID => %s", id)
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var productRegister application.Product

	stmt, err := p.db.Prepare("SELECT id FROM products WHERE id = $1")
	if err != nil {
		fmt.Println(err.Error())

		return nil, err
	}

	err = stmt.QueryRow(product.GetID()).Scan(&productRegister.ID)
	if err != nil {
		productRegister = application.Product{}
	}

	if (productRegister == application.Product{}) {
		return p.create(product)
	} else {
		return p.update(product)
	}

}

func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare("INSERT INTO products(id, name, price, status) VALUES($1, $2, $3, $4)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus(),
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare("UPDATE products SET name = $1, price = $2, status = $3 WHERE id = $4")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetName(), product.GetPrice(), product.GetStatus(),
		product.GetID(),
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}
