package routes

import (
	"github.com/gin-gonic/gin"
	"proyecto/prueba2/controllers"
)

func Init() *gin.Engine  {
	r :=  gin.New()

	router(r)

	return r
}

func router(r *gin.Engine)  {
	v2 := r.Group("/v2")
	{
		v2.GET("load",controllers.UploadController)
	}
}