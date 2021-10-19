package db

import (
	"database/sql"

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
	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	result, err := p.Get(product.GetID())
	if err != nil {
		return nil, err
	}

	if result == nil {
		return p.create(product)
	} else {
		return p.update(product)
	}
}

func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare("INSERT INTO products(name, price, status) VALUES(?, ?, ?)")
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(
		product.GetName(), product.GetPrice(), product.GetStatus(),
	)

	if err != nil {
		return nil, err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare("UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(
		product.GetName(), product.GetPrice(), product.GetStatus(),
		product.GetID(),
	)

	if err != nil {
		return nil, err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return product, nil
}
