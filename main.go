package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	version   = "" // 编译时注入
	buildTime = "" // 编译时注入
	gitCommit = "" // 编译时注入
)

func info(c *gin.Context) {
	c.JSON(200, gin.H{
		"version":   version,
		"buildTime": buildTime,
		"gitCommit": gitCommit,
	})
}

func v1(c *gin.Context) {
	c.JSON(200, gin.H{"code": 0, "version": "v1"})
}
func v2(c *gin.Context) {
	c.JSON(200, gin.H{"code": 0, "version": "v2"})
}
func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func main() {
	router := gin.Default()
	router.GET("hello", hello)
	router.GET("v1", v1)
	router.GET("v2", v2)
	router.GET("", info)
	router.Run(":8000")
}
