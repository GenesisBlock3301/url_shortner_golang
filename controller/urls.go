package controller

import (
	"github.com/GenesisBlock3301/url_shortner_golang/model"
	"github.com/gin-gonic/gin"
)

func CreateUrlController(c *gin.Context) {
	url := model.URL{}
	url.CreateUrl(c)
}

func ForwardToTargetUrl(c *gin.Context)  {

}