package main

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{"1", "Blue Train", "John Coltrane", 56.99},
	{"2", "Jeru", "Gerry Mulligan", 17.99},
	{"3", "Sarah Vaughan and Clifford Brown", "Sarah Vaughan", 39.99},
}

func main() {
	// initialize a GIN router
	router := setupRouter()

	router.Run("localhost:8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/album/list", getList)
	router.GET("/album/:id", getById)
	router.POST("/album", create)
	router.PUT("/album/:id", updateAlbum)
	router.DELETE("/album/:id", deleteAlbum)

	return router
}

func getList(c *gin.Context) {
	// sort list
	var ascSortedAlbums = albums
	sort.SliceStable(ascSortedAlbums, func(i, j int) bool {
		return ascSortedAlbums[i].Price < ascSortedAlbums[j].Price
	})

	c.IndentedJSON(http.StatusOK, ascSortedAlbums)
	return
}

func getById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "album not found"})
	return
}

func checkExistedById(c *gin.Context) bool {
	var id = c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			return true
		}
	}
	return false
}

func create(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if checkExistedById(c) {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
	return
}

func updateAlbum(c *gin.Context) {
	var id = c.Param("id")
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for _, a := range albums {
		if a.ID == id {
			albums = append(albums, newAlbum)
			c.IndentedJSON(http.StatusCreated, newAlbum)
			return
		}
	}

}

func deleteAlbum(c *gin.Context) {
	var id = c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusFound, gin.H{"message": "album found"})
			deleteAlbum(c)
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "album not found"})
	return
}
