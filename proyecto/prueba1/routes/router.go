package routes

import (
	"github.com/gin-gonic/gin"
	"proyecto/prueba1/controllers"
)

func Init() *gin.Engine  {
	r :=gin.New()
	Routes(r)
	return  r
}

func Routes(r *gin.Engine)  {
	v1 :=r.Group("/resume")
	{
		v1.GET(":dateProcess/:days",controllers.GetPurchaseController)
	}
}