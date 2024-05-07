package main

import (
	"fmt"
	"net/http"
	"request-counter/database"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

type Name struct {
	Name string `json:"name"`
}

func getRequests(c *gin.Context) {
	access_key := c.Param("access_key")
	c.IndentedJSON(http.StatusOK, database.GetRequests(access_key))
}

func putRequest(c *gin.Context) {
	access_key := c.Param("access_key")
	err := database.PutRequest(access_key)
	if err != nil {
		c.HTML(http.StatusOK, "instantclose.html", gin.H{})
	}
	c.String(http.StatusOK, fmt.Sprint(database.GetRequestsCount(access_key)))
}

func main() {
	database.Connect()

	router := gin.Default()

	router.LoadHTMLGlob("html/*")

	router.GET("/requests/get/:access_key", getRequests)
	router.GET("/requests/log/:access_key", putRequest)

	router.Run("0.0.0.0:8080")
}
