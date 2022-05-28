package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type menu struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	InStock bool   `json:"instock"`
}

var menuList = []menu{
	{Id: 1, Name: "皮", Price: 120, InStock: true},
	{Id: 2, Name: "レバー", Price: 150, InStock: true},
	{Id: 3, Name: "ぼんじり", Price: 180, InStock: true},
}

func getMenuList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, menuList)
}

func postItem(c *gin.Context) {
	var newMenu menu
	err := c.BindJSON(&newMenu)
	if err != nil {
		return
	}
	menuList = append(menuList, newMenu)
	c.IndentedJSON(http.StatusCreated, newMenu)
}

func getItemById(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	for _, a := range menuList {
		if a.Id == i {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

func updateStockById(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	for _, a := range menuList {
		if a.Id == i {
			stock := a.InStock
			if stock == false {
				a.InStock = true
			} else {
				a.InStock = false
			}
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

func DeleteById(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	for j, a := range menuList {
		if a.Id == i {
			menuList = append(menuList[:j], menuList[j+1:]...)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

func main() {
	router := gin.Default()
	router.GET("/menu", getMenuList)
	router.GET("/menu/:id", getItemById)
	router.POST("/menu", postItem)
	router.GET("/menu/outofstock/:id", updateStockById)
	router.GET("/menu/delete/:id", DeleteById)
	router.Run("localhost:8080")
}
