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
}

type Staff struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Fullname string `json:"fullname"`
	Address string 	`json:"address"`
	PhoneNumber string `json:"phone_number"`
	Salary uint	`json:"salary"`
	Role string	`json:"role"`
}

type Customer struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Fullname string `json:"fullname"`
	Address string	`json:"address"`
	PhoneNumber string `json:"phone_number"`
	EcommerceID uint `json:"ecommerce_id"`
	Ecommerce Ecommerce
}

type Report struct {
	ID uint `gorm:"primaryKey" json:"id"`
	From time.Time `json:"from"`
	To time.Time	`json:"toe"`
	Income uint	`json:"income"`
	Outcome uint	`json:"outcome"`
	Profit int `json:"profit"`
	StaffID uint `json:"staff_id"`
	Staff Staff
	CreatedAt time.Time
}

type Discout struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	Percentage float32	`json:"percentage"`
	EcommerceID uint `json:"ecommerce_id"`
	Ecommerce Ecommerce `json:"ecommerce"`
}

type EcommerceAccount struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Name string	`json:"name"`
	Username string	`json:"username"`
	ProfileUrl string `json:"profile_url"`
	EcommerceID uint `json:"ecommerce_id"`
	Ecommerce Ecommerce
}

type Order struct {
	ID uint `gorm:"primaryKey" json:"id"`
	TotalPrice uint `json:"total_price"`
	CustomerID uint `json:"customer_id"`
	DiscoutID uint	`json:"discount_id"`
	EcommerceAccountID uint	`json:"ecommerce_account_id"`
	StaffID uint	`json:"staff_id"`
	Customer Customer
	Discout Discout
	EcommerceAccount EcommerceAccount
	Staff Staff
}

type ProductOrder struct {
	ID uint `gorm:"primaryKey" json:"id"`
	ProductID uint	`json:"product_id"`
	OrderID uint	`json:"order_id"`
	Product Product
	Order Order
	Quantity uint	`json:"quantity"`
	Price uint	`json:"price"`
}