package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	conection *sql.DB
}

func NewProductRepository(conection *sql.DB) ProductRepository {
	return ProductRepository{
		conection: conection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "Select * from product"
	rows, err := pr.conection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productLista []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productLista = append(productLista, productObj)
	}

	rows.Close()

	return productLista, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {

	var id int

	query, err := pr.conection.Prepare("INSERT INTO product " +
		"(product_name, price)" +
		" VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *ProductRepository) GetProductById(product_id int) (*model.Product, error) {

	query, err := pr.conection.Prepare("Select * from product where id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product model.Product

	err = query.QueryRow(product_id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Println(err)
		return nil, err
	}

	query.Close()

	return &product, nil
}
