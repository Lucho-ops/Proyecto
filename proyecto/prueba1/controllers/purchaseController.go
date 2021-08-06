package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"proyecto/prueba1/services"
	"proyecto/prueba1/shared"
	"strconv"
)

func GetPurchaseController(c *gin.Context)  {
	response := shared.Context{}
	days,_ := strconv.ParseInt(c.Param("days"), 10,64)
	dateProcess := c.Param("dateProcess")

	fmt.Println(dateProcess,days)

	data, err := services.GetPurchaseServices(dateProcess ,days)
	if err != nil {
		c.JSON(400,response.ResponseData(shared.StatusError,err.Error(),""))
	}

	c.JSON(200,response.ResponseData(shared.StatusSuccess,"",data))

}