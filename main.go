package main

import (
	"github.com/albugowy15/basketfood-api/controller/account"
	"github.com/albugowy15/basketfood-api/controller/customer"
	"github.com/albugowy15/basketfood-api/controller/discount"
	"github.com/albugowy15/basketfood-api/controller/order"
	"github.com/albugowy15/basketfood-api/controller/product"
	"github.com/albugowy15/basketfood-api/controller/report"
	"github.com/albugowy15/basketfood-api/controller/staff"
	"github.com/albugowy15/basketfood-api/model"
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	model.ConnectDB()

	// products api
	r.GET("/api/products", product.Index)
	r.GET("/api/product/:id", product.Show)
	r.POST("/api/product", product.Create)
	r.PUT("/api/product/:id", product.Update)
	r.DELETE("/api/product/:id", product.Delete)


	// orders api
	r.GET("/api/orders", order.Index)
	r.GET("/api/order/:id", order.Show)
	r.POST("/api/order", order.Create)
	r.DELETE("/api/order/:id", order.Delete)

	// staffs api
	r.GET("/api/staffs", staff.Index)
	r.GET("/api/staff/:id", staff.Show)
	r.POST("/api/staff", staff.Create)
	r.PUT("/api/staff/:id", staff.Update)
	r.DELETE("/api/staff/:id", staff.Delete)

	// cutomers api
	r.GET("/api/customers", customer.Index)
	r.GET("/api/customer/:id", customer.Show)
	r.POST("/api/customer", customer.Create)
	r.PUT("/api/customer/:id", customer.Update)
	r.DELETE("/api/customer/:id", customer.Delete)

	// reports api
	r.GET("/api/reports", report.Index)
	r.GET("/api/report/:id", report.Show)
	r.POST("/api/report", report.Create)
	r.DELETE("/api/report/:id", report.Delete)

	// discounts api
	r.GET("/api/discounts", discount.Index)
	r.GET("/api/discount/:id", discount.Show)
	r.POST("/api/discount", discount.Create)
	r.PUT("/api/discount/:id", discount.Update)
	r.DELETE("/api/discount/:id", discount.Delete)

	// accounts api
	r.GET("/api/accounts", account.Index)
	r.GET("/api/account/:id", account.Show)
	r.POST("/api/account", account.Create)
	r.PUT("/api/account/:id", account.Update)
	r.DELETE("/api/account/:id", account.Delete)


	r.Run()
}