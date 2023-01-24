package staff

import (
	"net/http"

	"github.com/albugowy15/basketfood-api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Index(c *gin.Context) {
	type Staff struct {
		ID uint `gorm:"primaryKey" json:"id"`
		Fullname string `json:"fullname"`
		Address string 	`json:"address"`
		PhoneNumber string `json:"phone_number"`
		Salary uint	`json:"salary"`
		Role string	`json:"role"`
	}
	var staffs []Staff

	model.DB.Order("id").Find(&staffs)

	c.JSON(http.StatusOK, gin.H{
		"staffs": staffs,
	})
}

func Show(c *gin.Context) {
	id := c.Param("id")
	var staff model.Staff

	err := model.DB.Model(&staff).Preload(clause.Associations).First(&staff, id).Error

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Staff tidak dapat ditemukan",
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
		"staff": staff,
	})
	
}

func Create(c *gin.Context) {
	var staff model.Staff

	err := c.ShouldBindJSON(&staff)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = model.DB.Create(&staff).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Tidak dapat menambahkan staff",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"id": staff.ID,
	})
	
}

func Update(c *gin.Context) {
	var staff model.Staff
	id := c.Param("id")

	err := c.ShouldBindJSON(&staff)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	rows := model.DB.Model(&staff).Where("id = ?", id).Updates(&staff).RowsAffected == 0
	if rows {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Tidak dapat mengupdate staff",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data staff berhasil diperbarui",
	})
	
}

func Delete(c *gin.Context) {
	var staff model.Staff
	id := c.Param("id")


	if model.DB.Delete(&staff, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Tidak dapat menghapus staff",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Staff berhasil dihapus"})
}