package repositories

import (
	"BE-Golang-Coffee-Shop/internal/models"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	*sqlx.DB
}

func InitializeProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) RepositoryGetAllProduct() ([]models.ProductModel, error) {
	result := []models.ProductModel{}
	query := `SELECT * FROM products`

	err := r.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *ProductRepository) RepositoryGetProductID(id string) ([]models.ProductModel, error) {
	result := []models.ProductModel{}
	query := `SELECT * from products WHERE product_id = $1`
	err := r.Select(&result, query, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *ProductRepository) RepositoryCreateProduct(body *models.ProductModel) (sql.Result, error) {
	query := `INSERT INTO products (product_name, product_desc, product_category, product_price) VALUES (:product_name, :product_desc, :product_category, :product_price)`

	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *ProductRepository) RepositoryUpdateProduct(body *models.ProductModel, id string) (sql.Result, error) {
	query := `UPDATE products SET product_name=:product_name updated_at = now() WHERE product_id =` + id
	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *ProductRepository) RepositoryDeleteProduct(body *models.ProductModel, id string) (sql.Result, error) {
	query := `DELETE FROM products WHERE product_id = ` + id
	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *ProductRepository) RepositorySearchProduct(name string) ([]models.ProductModel, error) {
	result := []models.ProductModel{}
	query := `SELECT product_id, product_name, product_desc, product_price FROM products`

	if name != "" {
		query += ` WHERE product_name ilike $1`
		err := r.Select(&result, query, "%"+name+"%")
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (r *ProductRepository) RepositoryFilterPriceProduct(minPrice, maxPrice string) ([]models.ProductModel, error) {
	result := []models.ProductModel{}
	query := `SELECT product_id, product_name, product_desc, product_price FROM products`

	if minPrice != "" && maxPrice != "" {
		query += ` WHERE product_price >= $1 AND product_price <= $2`
		err := r.Select(&result, query, minPrice, maxPrice)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
