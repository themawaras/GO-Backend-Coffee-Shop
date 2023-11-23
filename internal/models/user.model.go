package models

import "time"

type UserModel struct {
	User_id       int         `db:"user_id"`
	User_fullname string      `db:"user_fullname" form:"user_fullname" json:"user_fullname"`
	User_email    string      `db:"user_email" form:"user_email" json:"user_email"`
	User_phone    string      `db:"user_phone" form:"user_phone" json:"user_phone"`
	User_address  string      `db:"user_address" form:"user_address" json:"user_address"`
	IsAdmin       bool        `db:"isAdmin" form:"isAdmin" json:"isAdmin"`
	User_password string      `db:"user_password" form:"user_password" json:"user_password"`
	User_picture  interface{} `db:"user_picture" form:"user_picture" json:"user_picture"`
	User_otp      int         `db:"user_otp" form:"user_otp" json:"user_otp"`
	IsActivated   bool        `db:"isActivated" form:"isActivated" json:"isActivated"`
	Created_at    *time.Time  `db:"created_at"`
	Updated_at    *time.Time  `db:"updated_at"`
}
