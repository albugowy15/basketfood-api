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

	// get income
	var income float64

	res := model.DB.Table("orders").Select("SUM(total_price) as income").Where("created_at BETWEEN ? AND ?", reportRequest.From, reportRequest.To).Scan(&income)

	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": res.Error.Error()})
		return
	}

	var outcome float64

	res = model.DB.Table("orders").Select("SUM(product_orders.quantity * products.production_price) as outcome").Joins("JOIN product_orders ON orders.id = product_orders.order_id").Joins("JOIN products ON product_orders.product_id = products.id").Where("orders.created_at BETWEEN ? AND ?", reportRequest.From, reportRequest.To).Scan(&outcome)
	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": res.Error.Error()})
		return
	}
	
	profit := income - outcome

	report := model.Report{
		From: &reportRequest.From,
		To: &reportRequest.To,
		Income: uint(income),
		Outcome: uint(outcome),
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