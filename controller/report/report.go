package report

import (
	"net/http"
	"time"

	"github.com/albugowy15/basketfood-api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	type Report struct {
		ID uint `gorm:"primaryKey" json:"id"`
		From *time.Time `json:"from" time_format:"2006-01-02"`
		To *time.Time	`json:"to" time_format:"2006-01-02"`
		Income uint	`json:"income"`
		Outcome uint	`json:"outcome"`
		Profit int `json:"profit"`
		CreatedAt *time.Time	`gorm:"auto_preload" json:"created_at" time_format:"2006-01-02"`
	}
	var reports []Report

	res := model.DB.Find(&reports)
	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reports": reports})
	
}

func Show(c *gin.Context) {
	var report model.Report
	id := c.Param("id")

	res := model.DB.Preload("Staff").First(&report, id)
	if res.Error != nil {
		switch res.Error {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Report not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": res.Error.Error()})
			return			
		}
	}

	c.JSON(http.StatusOK, gin.H{"report": report})
	
}

func Create(c *gin.Context) {
	type ReportRequest struct {
		From time.Time `json:"from"`
		To time.Time	`json:"to"`
		StaffID uint `json:"staff_id"`
	}

	var reportRequest ReportRequest

	err := c.ShouldBindJSON(&reportRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var orders []model.Order

	res := model.DB.Where("created_at BETWEEN ? AND ?", reportRequest.From, reportRequest.To).Preload("Orders").Preload("Orders.Product").Find(&orders)

	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": res.Error.Error()})
		return
	} else if res.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "No orders found"})
		return
	}

	// get income
	var income uint
	var outcome uint

	for _, order := range orders {
		for _, productOrder := range order.Orders {
			outcome += productOrder.Product.ProductionPrice * productOrder.Quantity
		}
		income += order.TotalPrice
	}
	
	
	profit := income - outcome

	report := model.Report{
		From: &reportRequest.From,
		To: &reportRequest.To,
		Income: income,
		Outcome: outcome,
		Profit: int(profit),
		StaffID: reportRequest.StaffID,
	}

	res = model.DB.Create(&report)

	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": res.Error.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Report created successfully", "income": report.Income, "outcome": report.Outcome, "profit": report.Profit})
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	var report model.Report


	err := model.DB.Where("id = ?", id).Delete(&report).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Report deleted successfully"})
	
}