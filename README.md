# URL shortener #

URL shortener is a third-party website that converts that long URL to a short, case-sensitive alphanumeric code..


## Installation ##

URLshortener is compatible with modern Go releases in module mode, with Go installed:
Install all package by using `go get <git path>`
```bash "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"hash/fnv"
	"net/url"
	"os"
```

### Project Layout

```tree
├── config
│   ├── db
│   │   └── database-config.go
|   |   └── loader.go
│   ├── helper
│   │   └── GetDB.go
├── controllers
│   └── urls.go
├── logger
│   └── log.go
├── model
│   └── url.go
├── router
│   └── route-router.go
├── main.go
├── .gitignore
├── README.md
```
## Run project:
```bash
go run main.go
```

## End points ##

```bash
localhost:8000/create/
```
Request body
```json
{
	"longUrl": "https://github.com/GenesisBlock3301/full_featured_golang_app"
}
```
Redirect to target url using short url
```bash
localhost:8000/<url_token>
```

## Learning Scope ##
1. Url shortener.
2. how to work 301 & 302 url redirect work.
3. Base64 
4. How to work hash


## License ##

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE)
file.
