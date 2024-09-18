package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	Router := gin.Default()

	userGroup := Router.Group("/users")
	{
		userGroup.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"data": "ok",
			})
		})
	}

	return Router
}
