package model

import (
	"time"
)

type Ecommerce struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Url string	`json:"url"`
}

type Product struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Stock uint `json:"stock"`
	SellingPrice uint `json:"selling_price"`
	ProductionPrice uint `json:"production_price"`
	Profit uint `json:"profit"`
	ProductOrders []ProductOrder `json:"-"`
}

type Staff struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Fullname string `json:"fullname"`
	Address string 	`json:"address"`
	PhoneNumber string `json:"phone_number"`
	Salary uint	`json:"salary"`
	Role string	`json:"role"`
	Orders []Order	`json:"orders"`
	Reports []Report	`json:"reports"`
}

type Customer struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Fullname string `json:"fullname"`
	Address string	`json:"address"`
	PhoneNumber string `json:"phone_number"`
	EcommerceID uint `json:"ecommerce_id"`
	Ecommerce Ecommerce `json:"ecommerce"`
	Orders []Order `json:"orders"`
}

type Report struct {
	ID uint `gorm:"primaryKey" json:"id"`
	From *time.Time `json:"from" time_format:"2006-01-02"`
	To *time.Time	`json:"to" time_format:"2006-01-02"`
	Income uint	`json:"income"`
	Outcome uint	`json:"outcome"`
	Profit int `json:"profit"`
	StaffID uint `json:"staff_id"`
	Staff Staff	`json:"staff"`
	CreatedAt *time.Time	`gorm:"auto_preload" json:"created_at" time_format:"2006-01-02"`
}

type Discout struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	Percentage float32	`json:"percentage"`
	EcommerceID uint `json:"ecommerce_id"`
	Ecommerce Ecommerce `json:"ecommerce"`
	Orders []Order `json:"orders"`
}

type EcommerceAccount struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Name string	`json:"name"`
	Username string	`json:"username"`
	ProfileUrl string `json:"profile_url"`
	EcommerceID uint `json:"ecommerce_id"`
	Ecommerce Ecommerce	`json:"ecommerce"`
	Orders []Order	`json:"orders"`
}

type Order struct {
	ID uint `gorm:"primaryKey" json:"id"`
	TotalPrice uint `json:"total_price"`
	CustomerID uint `json:"customer_id"`
	Customer Customer `json:"customer"`
	DiscoutID uint	`json:"discount_id"`
	Discout Discout	`json:"discount"`
	EcommerceAccountID uint	`json:"ecommerce_account_id"`
	EcommerceAccount EcommerceAccount	`json:"ecommerce_account"`
	StaffID uint	`json:"staff_id"`
	Staff Staff	`json:"staff"`
	Orders []ProductOrder	`json:"orders"`
	CreatedAt *time.Time `gorm:"auto_preload" json:"created_at" time_format:"2006-01-02"`
}

type ProductOrder struct {
	ID uint `gorm:"primaryKey" json:"id"`
	ProductID uint	`json:"product_id"`
	Product Product	`json:"product"`
	OrderID uint	`json:"order_id"`
	Order Order	`json:"order"`
	Quantity uint	`json:"quantity"`
	Price uint	`json:"price"`
}