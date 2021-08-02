package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Get all users
func GetUsers(c *gin.Context) {
	var user []models.User

	err := models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.Json(http.StatusOK, user)
	}
}

//Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	err := models.GetUSerByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//Create user
func CreateUser(c *gin.Context) {
	var user models.User

	c.BindJson(&user)
	err := models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//Update the user information
func Updateuser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	err := models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.BindJSON(&user)

	err = models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//Delete the use
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	err := models.DeleteUSer(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
