package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BlogIndex(c *gin.Context) {
	c.String(http.StatusOK, "blog index")
}

func Blog(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"category": c.Param("category"), "post": c.Param("post")})
}

func User(c *gin.Context) {
	c.String(http.StatusOK, "hello, %s!\n", c.Param("username"))
}

// --- 定义一些处理器 ---
func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"userid": "1", "username": "kingofzihua"})
}
func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"userid": "1", "username": "kingofzihua"})
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Welcome!\n")
}

func HelloGin(c *gin.Context) {
	c.String(http.StatusOK, "hello gin")
}

type Person struct {
	Name    string `uri:"name"`
	Address string `form:"address"`
}

func Testing(c *gin.Context) {
	var person Person

	err := c.ShouldBind(&person)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"name": person.Name, "address": person.Address})
}
