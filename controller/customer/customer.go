package customer

import (
	"net/http"

	"github.com/albugowy15/basketfood-api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Customer struct {
	ID uint `json:"id"`
	Fullname string `json:"fullname"`
	Address string	`json:"address"`
	PhoneNumber string `json:"phone_number"`
}
func Index(c *gin.Context) {
	var customers []Customer

	err := model.DB.Order("id").Find(&customers).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"customers": customers,
	})

	
}

func Show(c *gin.Context) {
	var customer model.Customer

	id := c.Param("id")

	err := model.DB.Model(&customer).Preload(clause.Associations).First(&customer, id).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound: 
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Customer tidak ditemukan",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Terjadi kesalahan di server",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"customer": customer,
	})
	
}

func Create(c *gin.Context) {
	var customer model.Customer

	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = model.DB.Create(&customer).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": customer.ID,
	})
	
}

func Update(c *gin.Context) {
	var customer model.Customer
	id := c.Param("id")

	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"messsage": err.Error(),
		})
		return
	}

	err = model.DB.Model(&customer).Where("id = ?", id).Updates(&customer).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Customer berhasil diupdate",
	})
	
}

func Delete(c *gin.Context) {
	var customer model.Customer
	id := c.Param("id")

	countRows := model.DB.Delete(&customer, id).RowsAffected == 0

	if countRows {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Customer tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Customer berhasil dihapus",
	})
	
}