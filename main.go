package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// spirit represents data about a record gin.
type spirit struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    Description string  `json:"description"`
    Price  float64 `json:"price"`
}

// gins slice to seed record spirit data.
var spirits = []spirit{
    {ID: "1", Name: "Monkey 47 Gin 0,5 Liter", Description: "Monkey 47 - Der populäre Renner aus dem Schwarzwald", Price: 34.99},
    {ID: "2", Name: "Elephant London Dry Gin 0,5 Liter", Description: "Elephant Gin - Der exotische Riese", Price: 37.99},
    {ID: "3", Name: "Boar Gin Royal 0,5 Liter", Description: "Boar Gin 0,5 Liter - Blackforest Gin mit Trüffel", Price: 54.99},
}

func main() {
    router := gin.Default()
    router.GET("/gins", getGins)
    router.GET("/gins/:id", getGinByID)
    router.POST("/gins", postGins)
    router.StaticFile("/openapi.json", "openapi.json")

    router.RunTLS(":8000","server.crt", "server.key")
}

// getGins responds with the list of all gins as JSON.
func getGins(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, spirits)
}

// postGins adds an gin from JSON received in the request body.
func postGins(c *gin.Context) {
    var newGin spirit

    // Call BindJSON to bind the received JSON to
    // newGin.
    if err := c.BindJSON(&newGin); err != nil {
        return
    }

    // Add the new gin to the slice.
    spirits = append(spirits, newGin)
    c.IndentedJSON(http.StatusCreated, newGin)
}

// getGinByID locates the gin whose ID value matches the id
// parameter sent by the client, then returns that gin as a response.
func getGinByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of gins, looking for
    // an gin whose ID value matches the parameter.
    for _, a := range spirits {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Gin not found"})
}
