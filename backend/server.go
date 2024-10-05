package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users = []User{
	{Id: "1", Username: "samiransari", Email: "samiransari@gmail.com"},
	{Id: "2", Username: "krish", Email: "krish@gmail.com"},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserById)
	router.POST("/users", postUsers)

	router.Run("localhost:8080")
}

// getUsers returns a list of users
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func postUsers(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getUserById(c *gin.Context) {
	id := c.Param("id")

	for _, user := range users {
		if user.Id == id {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found!"})
}
