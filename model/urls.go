package model

import (
	"fmt"
	"github.com/GenesisBlock3301/url_shortner_golang/config/helpers"
	"github.com/GenesisBlock3301/url_shortner_golang/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"hash/fnv"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	CollectionName = "url"
)

type URL struct {
	Key     string `json:"key" bson:"_id,omitempty"`
	LongUrl string `bson:"longUrl omitempty"`
}

// CreateUrl for create short url
// Firstly hashed the long url which get from client
// Then calculate value based on Base 62

func (u *URL) CreateUrl(c *gin.Context) {
	// Create of get collection on database.
	_db := helpers.GetDB(CollectionName)

	ctx, cancel := helpers.GetContext()
	defer cancel()
	var urlData URL
	// Get value from frontend as a json value.
	mappingErr := c.Bind(&urlData)

	// Check long url correct or incorrect.
	if !IsUrl(urlData.LongUrl) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid url error"})
		return
	}
	if mappingErr != nil {
		c.JSON(http.StatusOK, gin.H{"payload": mappingErr.Error()})
		logger.Log{Message: "Mapping error"}.Error()
		return
	}
	// In url only use these characters
	BASE62 := os.Getenv("BASE64")
	//Hashed the LongUrl
	hashValue := hash(urlData.LongUrl)
	var str []string
	// Calculate value based on Base 62 and find
	//value from `BASE62` according to this value
	for hashValue > 0 {
		rem := hashValue % 62
		hashValue = hashValue / 62
		str = append(str, string(BASE62[rem]))
	}
	// Concat host with hashed value
	u.Key = strings.Join(str, "")
	u.LongUrl = urlData.LongUrl
	// Insert long and short link into mongodb database
	res, insertErr := _db.InsertOne(ctx, u)
	if insertErr != nil {
		logger.Log{Message: insertErr.Error()}.Error()
		c.JSON(http.StatusBadRequest, gin.H{"error": insertErr.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"payload": res})
	}
}
func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
func hash(s string) uint64 {
	h := fnv.New64()
	h.Write([]byte(s))
	return h.Sum64()
}

func (u *URL) ForwardToTargetUrl(c *gin.Context) {
	id := c.Params.ByName("key")
	_db := helpers.GetDB(CollectionName)
	ctx, cancel := helpers.GetContext()
	defer cancel()
	var _url URL
	filter := bson.M{"_id": bson.M{"$eq": id}}
	err := _db.FindOne(ctx, filter).Decode(&_url)
	if err != nil {
		logger.Log{Context: c, Message: err.Error()}.Error()
		c.JSON(http.StatusBadRequest, gin.H{"payload": err.Error()})
		return
	}
	//redirectUrl := c.Request.Host + "/" + _url.Key
	fmt.Println(_url.LongUrl)
	c.Redirect(http.StatusMovedPermanently, _url.LongUrl)
	c.JSON(http.StatusOK, gin.H{"payload": _url})
}
