package models

import "time"

type OrderModel struct {
	Order_id      int        `db:"order_id"`
	User_id       int        `db:"user_id" form:"user_id" json:"user_id"`
	Subtotal      int        `db:"subtotal" form:"subtotal" json:"subtotal"`
	Promo_id      int        `db:"promo_id" form:"promo_id" json:"promo_id"`
	Serve_id      int        `db:"serve_id" form:"serve_id" json:"serve_id"`
	Total_disc    int        `db:"total_disc" form:"total_disc" json:"total_disc"`
	Order_total   int        `db:"order_total" form:"order_total" json:"order_total"`
	Payment_id    int        `db:"payment_id" form:"payment_id" json:"payment_id"`
	Order_status  string     `db:"order_status" form:"order_status" json:"order_status"`
	User_fullname string     `db:"user_fullname"`
	Created_at    *time.Time `db:"created_at"`
	Updated_at    *time.Time `db:"updated_at"`
}
