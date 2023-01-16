package product

import (
	"net/http"

	"github.com/albugowy15/basketfood-api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Description string `json:"description"`
	Stock uint `json:"stock"`
	SellingPrice uint `json:"selling_price"`
}

func Index(c *gin.Context) {
	var products []Product

	model.DB.Order("id").Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func Show(c *gin.Context) {
	var product model.Product
	id := c.Param("id")
	err := model.DB.First(&product, id).Error

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Data product tidak ditemukan",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Terjadi kesalahan di server",
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func Create(c *gin.Context) {
	// retrive data from body
	// convert to object variable
	// send it to database
	// return the response

	var product model.Product

	err := c.ShouldBindJSON(&product)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = model.DB.Create(&product).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": product.ID,
	})
}

func Update(c *gin.Context) {
	var product model.Product
	id := c.Param("id")

	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	rows := model.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0
	if rows {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Tidak dapat mengupdate product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data product berhasil diperbarui",
	})
	
}

func Delete(c *gin.Context) {
	var product model.Product
	id := c.Param("id")


	if model.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Tidak dapat menghapus product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product berhasil dihapus"})
}