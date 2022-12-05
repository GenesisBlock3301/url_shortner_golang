package route

import (
	"github.com/GenesisBlock3301/url_shortner_golang/controller"
	"github.com/gin-gonic/gin"
)

func RootRouter(router *gin.Engine)  {
	router.POST("/create", func(context *gin.Context) {
		controller.CreateUrlController(context)
	})

	router.GET("/:id", func(context *gin.Context) {
		controller.ForwardToTargetUrl(context)
	})
}