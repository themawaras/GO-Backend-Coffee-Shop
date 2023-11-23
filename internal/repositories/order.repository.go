package repositories

import (
	"BE-Golang-Coffee-Shop/internal/models"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	*sqlx.DB
}

func InitializeOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) RepositoryGetAllOrder() ([]models.OrderModel, error) {
	result := []models.OrderModel{}
	// query := `SELECT order_id, user_id, subtotal, promo_id, serve_id, total_disc, order_total, payment_id, order_status FROM orders`
	query := `SELECT * FROM orders`

	err := r.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *OrderRepository) RepositoryGetOrderID(id string) ([]models.OrderModel, error) {
	result := []models.OrderModel{}
	query := `SELECT * FROM orders WHERE order_id = $1`
	// query := `SELECT o.order_id, u.user_fullname, o.subtotal, o.promo_id, o.serve_id, o.total_disc, o.order_total, o.payment_id, o.order_status FROM orders o join users u on o.user_id = u.user_id WHERE o.order_id = $1`

	err := r.Select(&result, query, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *OrderRepository) RepositoryCreateOrder(body *models.OrderModel) (sql.Result, error) {
	query := `INSERT INTO orders (user_id, subtotal, promo_id, serve_id, total_disc, order_total, payment_id, order_status) VALUES (:user_id, :subtotal, :promo_id, :serve_id, :total_disc, :order_total, :payment_id, :order_status)`

	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *OrderRepository) RepositoryUpdateOrder(body *models.OrderModel, id string) (sql.Result, error) {
	query := `UPDATE orders SET order_status = :order_status, updated_at = now() WHERE order_id =` + id
	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *OrderRepository) RepositoryDeleteOrder(body *models.OrderModel, id string) (sql.Result, error) {
	query := `DELETE FROM orders WHERE order_id = ` + id
	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err
	}
	return result, nil
}
