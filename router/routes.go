package router

import (
	"github.com/Ruberson1/api-gojobs/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpeningHandler)

		v1.POST("/opening", handler.CreatOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.GET("/openings", handler.ListOpeningHandler)
	}
}
