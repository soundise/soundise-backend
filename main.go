package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "golang.org/x/crypto/argon2"
)

func main() {

	router := gin.Default()

	router.GET("/hello", hello)
	router.Run("0.0.0.0:8000")

}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
