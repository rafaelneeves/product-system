package service

import (
	"database/sql"
	"errors"
)

type Product struct {
	id     int
	name   string
	price  float32
	idUser int
}

type ProductService struct {
	db *sql.DB
}

func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{db: db}
}

// Retorna lista de Produtos
func (s *ProductService) GetProducts() ([]Product, error) {
	query := "SELECT * FROM products"
	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product

		if err := rows.Scan(&product.id, &product.name, &product.price, &product.idUser); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

// Retorna um produto pelo seu ID.
func (s *ProductService) GetProductById(id int) (*Product, error) {
	query := "SELECT * FROM product WHERE id = ?"
	row := s.db.QueryRow(query, id)

	var product Product
	if err := row.Scan(&product.id, &product.name, &product.price, &product.idUser); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

// Cria um Produto
func (s *ProductService) CreateProduct(product *Product) error {
	query := "INSERT INTO products (name, price, id_user) VALUES (?,?,?)"
	result, err := s.db.Exec(query, product.name, product.price, product.idUser)

	if err != nil {
		return err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	product.id = int(lastInsertID)

	return nil
}

// Atualiza um produto
func (s *ProductService) UpdateBook(product *Product) error {
	query := "UPDATE products SET name = ?, price = ? WHERE id = ?"
	_, err := s.db.Exec(query, product.name, product.price)

	return err
}

// Deleta um Produto
func (s *ProductService) DeleteBook(id int) error {
	query := "DELETE FROM product WHERE id = ?"
	_, err := s.db.Exec(query, id)
	return err
}
