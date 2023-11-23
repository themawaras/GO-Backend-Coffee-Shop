package models

import "time"

type PromoModel struct {
	Promo_id       int        `db:"promo_id"`
	Promo_name     string     `db:"promo_name" form:"promo_name" json:"promo_name"`
	Promo_desc     string     `db:"promo_desc" form:"promo_desc" json:"promo_desc"`
	Flat_amount    int        `db:"flat_amount" form:"flat_amount" json:"flat_amount"`
	Percent_amount float32    `db:"percent_amount" form:"percent_amount" json:"percent_amount"`
	Created_at     *time.Time `db:"created_at"`
	Expired_at     *time.Time `db:"expired_at"`
}
