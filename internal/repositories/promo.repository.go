package repositories

import (
	"BE-Golang-Coffee-Shop/internal/models"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type PromoRepository struct {
	*sqlx.DB
}

func InitializePromoRepository(db *sqlx.DB) *PromoRepository {
	return &PromoRepository{db}
}

func (r *PromoRepository) RepositoryGetAllPromo() ([]models.PromoModel, error) {
	result := []models.PromoModel{}
	query := `SELECT * FROM promos`

	err := r.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *PromoRepository) RepositoryGetPromoID(id string) ([]models.PromoModel, error) {
	result := []models.PromoModel{}
	query := `SELECT * FROM promos WHERE promo_id = $1`
	err := r.Select(&result, query, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *PromoRepository) RepositoryCreatePromo(body *models.PromoModel) (sql.Result, error) {
	query := `INSERT INTO promos (promo_name, promo_desc, flat_amount, percent_amount, expired_at) VALUES (:promo_name, :promo_desc, :flat_amount, :percent_amount, :expired_at)`

	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *PromoRepository) RepositoryUpdatePromo(body *models.PromoModel, id string) (sql.Result, error) {
	query := `UPDATE promos SET promo_name = :promo_name, promo_desc = :promo_desc WHERE promo_id =` + id
	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *PromoRepository) RepositoryDeletePromo(body *models.PromoModel, id string) (sql.Result, error) {
	query := `DELETE FROM promos WHERE promo_id = ` + id
	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err
	}
	return result, nil
}
