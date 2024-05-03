package main

import (
	"net/http"
	"request-counter/database"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

type Name struct {
	Name string `json:"name"`
}

func getRequests(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.GetRequests())
}

func putRequest(c *gin.Context) {
	err := database.PutRequest()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.Status(http.StatusOK)
}

func main() {
	database.Connect()

	router := gin.Default()

	router.GET("/requests", getRequests)
	router.GET("/", putRequest)

	router.Run("0.0.0.0:8080")
}
