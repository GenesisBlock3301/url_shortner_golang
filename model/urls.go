package model

import (
	"github.com/GenesisBlock3301/url_shortner_golang/config/helpers"
	"github.com/GenesisBlock3301/url_shortner_golang/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CollectionName = "url"
)

type URL struct {
	Key       string `bson:"key omitempty"`
	SecretKey string `bson:"secretKey omitempty"`
	TargetUrl string `bson:"targetUrl omitempty"`
	IsActive  bool `bson:"isActive omitempty"`
	Clicks    string `bson:"clicks omitempty"`
}

func (u *URL) CreateUrl(c *gin.Context)  {
	_db := helpers.GetDB(CollectionName)
	ctx, cancel := helpers.GetContext()
	defer cancel()
	var urlData URL
	mappingErr := c.BindJSON(&urlData)
	if mappingErr != nil{
		c.JSON(http.StatusOK, gin.H{"payload": mappingErr.Error()})
		logger.Log{Message: "Mapping error"}.Error()
		return
	}
	res, insertErr := _db.InsertOne(ctx,urlData)
	if insertErr != nil{
		logger.Log{Message: insertErr.Error()}.Error()
		c.JSON(http.StatusBadRequest, gin.H{"error": insertErr.Error()})
	}else {
		c.JSON(http.StatusCreated, gin.H{"payload": res})
	}
}

func (u *URL) GetURL(){

}