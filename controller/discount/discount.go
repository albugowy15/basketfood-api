package discount

import (
	"net/http"

	"github.com/albugowy15/basketfood-api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Index(c *gin.Context) {
	type Discout struct {
		ID uint `gorm:"primaryKey" json:"id"`
		Title string `json:"title"`
		Percentage float32	`json:"percentage"`
	}

	var discounts []Discout

	err := model.DB.Model(&discounts).Order("id").Find(&discounts).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"discounts": discounts,
	})

}

func Show(c *gin.Context) {
	var discount model.Discout
	id := c.Param("id")

	err := model.DB.Model(&discount).Preload(clause.Associations).First(&discount, id).Error

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Discout tidak ditemukan",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"discount": discount})
	
}

func Update(c *gin.Context) {
	var discount model.Discout
	id := c.Param("id")

	err := c.ShouldBindJSON(&discount)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = model.DB.Model(&discount).Where("id = ?", id).Updates(discount).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Diskon berhasil diperbarui"})
	
}

func Create(c *gin.Context) {
	var discount model.Discout

	err := c.ShouldBindJSON(&discount)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = model.DB.Model(&discount).Create(&discount).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": discount.ID})
	
}

func Delete(c *gin.Context) {
	var discount model.Discout
	id := c.Param("id")

	countRows := model.DB.Delete(&discount, id).RowsAffected == 0

	if countRows {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Discout tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Discout berhasil dihapus",
	})
	
}