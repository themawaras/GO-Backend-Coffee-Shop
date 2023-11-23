package models

import "time"

type ProductModel struct {
	Product_id       int        `db:"product_id"`
	Product_name     string     `db:"product_name" form:"product_name" json:"product_name"`
	Product_desc     string     `db:"product_desc" form:"product_desc" json:"product_desc"`
	Product_category int        `db:"product_category" form:"product_category" json:"product_category"`
	Product_price    int        `db:"product_price" form:"product_price" json:"product_price"`
	Created_at       *time.Time `db:"created_at"`
	Updated_at       *time.Time `db:"updated_at"`
	Deleted_at       *time.Time `db:"deleted_at"`
}
