package account

import (
	"net/http"

	"github.com/albugowy15/basketfood-api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Index(c *gin.Context) {
	type EcommerceAccount struct {
		ID uint `gorm:"primaryKey" json:"id"`
		Name string	`json:"name"`
		Username string	`json:"username"`
		ProfileUrl string `json:"profile_url"`
		EcommerceID uint `json:"ecommerce_id"`
	}

	var accounts []EcommerceAccount

	err := model.DB.Model(&accounts).Order("id").Find(&accounts).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"accounts": accounts,
	})
}

func Show(c *gin.Context) {
	var account model.EcommerceAccount
	id := c.Param("id")

	err := model.DB.Model(&account).Preload(clause.Associations).First(&account, id).Error

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Akun ecommerce tidak ditemukan",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"akun": account})
	
}

func Create(c *gin.Context) {
	var account model.EcommerceAccount

	err := c.ShouldBindJSON(&account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = model.DB.Model(&account).Create(&account).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": account.ID})
}

func Update(c *gin.Context) {
	var account model.EcommerceAccount
	id := c.Param("id")

	err := c.ShouldBindJSON(&account)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = model.DB.Model(&account).Where("id = ?", id).Updates(account).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Akun ecommerce berhasil diperbarui"})
}

func Delete(c *gin.Context) {
	var account model.EcommerceAccount
	id := c.Param("id")

	countRows := model.DB.Delete(&account, id).RowsAffected == 0

	if countRows {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Akun ecommerce tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Akunn ecommerce berhasil dihapus",
	})
}