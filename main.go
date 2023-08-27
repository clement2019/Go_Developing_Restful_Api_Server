package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// creating the strcut for the data layer for the album database

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// creating a slice of albums struct that will be used to create the album records

var albums = []album{

	{ID: "1", Title: "Gold Truck music for the lord", Artist: "Truck wills", Price: 81.29},
	{ID: "2", Title: "clement Songs", Artist: "Clement Doe", Price: 36.79},
	{ID: "3", Title: "williams Songs of prraise", Artist: "John Wein", Price: 56.49},
}

//Now lets write the code to return our first endpoints

// getAlbums responds with the list of all albums as JSON.

func getalbums(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, albums)
}

// am, now writti g the new alsbum code

func postnewalbums(c *gin.Context) {

	var newalbum album

	if err := c.Bind(&newalbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	albums = append(albums, newalbum)
	c.IndentedJSON(http.StatusCreated, newalbum)

}

//now write a function that will get an album by ID

func getalbumByID(c *gin.Context) {

	id := c.Param("id")

	for _, a := range albums {

		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {

	router := gin.Default()
	router.GET("/albums", getalbums)

	fmt.Printf("server starting......")
	router.GET("/albums:id", getalbumByID)
	router.POST("/albums", postnewalbums)
	router.Run("localhost:8084")

}
