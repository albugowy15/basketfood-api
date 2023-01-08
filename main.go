package main

import (
	"github.com/albugowy15/basketfood-api/model"
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	model.ConnectDB()
	r.Run()
}