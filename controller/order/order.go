package order

import (
	"log"
	"net/http"

	"github.com/albugowy15/basketfood-api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Index(c *gin.Context) {
	type Order struct {
		ID uint `gorm:"primaryKey" json:"id"`
		TotalPrice uint `json:"total_price"`
		CustomerID uint `json:"customer_id"`
		DiscoutID uint	`json:"discount_id"`
		EcommerceAccountID uint	`json:"ecommerce_account_id"`
		StaffID uint	`json:"staff_id"`
	}

	var orders []Order

	model.DB.Order("id").Find(&orders)

	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
	
}

func Show(c *gin.Context) {
	type Customer struct {
		ID uint `gorm:"primaryKey" json:"id"`
		Fullname string `json:"fullname"`
		Address string	`json:"address"`
	}
	type Discout struct {
		ID uint `gorm:"primaryKey" json:"id"`
		Title string `json:"title"`
		Percentage float32	`json:"percentage"`
	}

	type EcommerceAccount struct {
		ID uint `gorm:"primaryKey" json:"id"`
		Name string	`json:"name"`
		Username string	`json:"username"`
	}

	type Staff struct {
		ID uint `gorm:"primaryKey" json:"id"`
		Fullname string `json:"fullname"`
		Address string 	`json:"address"`
	}
	type ProductOrder struct {
		ID uint `gorm:"primaryKey" json:"id"`
		ProductID uint	`json:"product_id"`
		OrderID uint	`json:"order_id"`
		Quantity uint	`json:"quantity"`
	Price uint	`json:"price"`
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
	}
	var order Order

	id := c.Param("id")

	err := model.DB.Model(Order{}).Preload(clause.Associations).First(&order, id).Error

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Order tidak dapat ditemukan",
			})
			return

		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"order": order,
	})
	
}

func Create(c *gin.Context) {

	type ProductOrderRequest struct {
		ProductID uint	`json:"product_id"`
		Quantity uint	`json:"quantity"`
		Price uint	`json:"price"`
	}

	type OrderRequest struct {
		CustomerID uint `json:"customer_id"`
		DiscoutID uint	`json:"discount_id"`
		EcommerceAccountID uint	`json:"ecommerce_account_id"`
		StaffID uint	`json:"staff_id"`
		Orders []ProductOrderRequest	`json:"orders"`
	}
	var orderRequest OrderRequest

	err := c.BindJSON(&orderRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	order := model.Order{
		TotalPrice: 0,
		CustomerID: orderRequest.CustomerID,
		DiscoutID: orderRequest.DiscoutID,
		EcommerceAccountID: orderRequest.EcommerceAccountID,
		StaffID: orderRequest.StaffID,
	}

	if err := model.DB.Save(&order).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var orderId = order.ID

	var totalPrice uint = 0
	tx := model.DB.Begin()
	for _, order := range orderRequest.Orders {
		type Product struct {
			ID uint `gorm:"primaryKey" json:"id"`
			SellingPrice uint `json:"selling_price"`
		}

		var product Product
		if err := tx.First(&product, order.ProductID).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Product tidak dapat ditemukan",
			})
			return
		}

		if err := tx.Model(&model.Product{}).Where("id = ?", order.ProductID).Update("stock", gorm.Expr("stock - ?", order.Quantity)).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Stock tidak cukup",
			})
			return
		}
		
		itemPrice := product.SellingPrice * order.Quantity

		if err := tx.Create(&model.ProductOrder{
			ProductID: order.ProductID,
			OrderID: orderId,
			Quantity: order.Quantity,
			Price: itemPrice,
		}).Error; err != nil {
			tx.Rollback()
			model.DB.Delete(&model.Order{}, orderId)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		totalPrice += itemPrice
	}
	tx.Commit()
	

	err = model.DB.Model(&model.Order{}).Where("id = ?", orderId).Updates(map[string]interface{}{"total_price": totalPrice}).Error
	if err != nil {
		log.Fatal(err)
	}
	
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Order berhasil dibuat",
	})
	
}


func Delete(c *gin.Context) {
	id := c.Param("id")
	var order model.Order

	err := model.DB.Model(&model.ProductOrder{}).Where("order_id = ?", id).Delete(&model.ProductOrder{}).Error

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Order tidak dapat ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
	}
	}

	model.DB.Delete(&order, id)

	c.JSON(http.StatusOK, gin.H{"message": "Order berhasil dihapus"})
	
}