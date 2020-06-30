package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r := gin.Default()
	r.GET("/", hello)
	r.Run()
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
}
