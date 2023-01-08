package model

import (
	"time"
)

type Ecommerce struct {
	ID uint `gorm:"primaryKey"`
	Name string
	Url string
}

type Product struct {
	ID uint `gorm:"primaryKey"`
	Name string
	Description string
	Stock uint
	SellingPrice uint
	ProductionPrice uint
	Profit uint
}

type Staff struct {
	ID uint `gorm:"primaryKey"`
	Fullname string
	Address string
	PhoneNumber string
	Salary uint
	Role string
}

type Customer struct {
	ID uint `gorm:"primaryKey"`
	Fullname string
	Address string
	PhoneNumber string
	EcommerceID uint
	Ecommerce Ecommerce
}

type Report struct {
	ID uint `gorm:"primaryKey"`
	Income uint
	Outcome uint
	Profit int
	Excel string
	StaffID uint
	Staff Staff
	CreatedAt time.Time
}

type Discout struct {
	ID uint `gorm:"primaryKey"`
	Title string
	Percentage float32
	EcommerceID uint
	Ecommerce Ecommerce
}

type EcommerceAccount struct {
	ID uint `gorm:"primaryKey"`
	Name string
	Username string
	ProfileUrl string
	EcommerceID uint
	Ecommerce Ecommerce
}

type Order struct {
	ID uint `gorm:"primaryKey"`
	TotalPrice uint
	CustomerID uint
	DiscoutID uint
	EcommerceAccountID uint
	StaffID uint
	Customer Customer
	Discout Discout
	EcommerceAccount EcommerceAccount
	Staff Staff
}

type ProductOrder struct {
	ID uint `gorm:"primaryKey"`
	ProductID uint
	OrderID uint
	Product Product
	Order Order
	Quantity uint
	Price uint
}