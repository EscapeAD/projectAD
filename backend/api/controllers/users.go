package controllers

import (
	"escapead/projectAD/backend/api/models"
	_response "escapead/projectAD/backend/api/responses"
	"net/http"
)

var users = []models.User{
	models.User{Nickname: "user 1", Email: "oops@youseemypassword.com", Password: "123456"},
	models.User{Nickname: "user 2", Email: "2@batman.com", Password: "123456"},
}

// GetUsers - GET all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	_response.JSON(w, http.StatusOK, users)
}

// CreateUser - GET user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	_response.JSON(w, http.StatusOK, map[string]string{})
}

// GetUser - GET an user
func GetUser(w http.ResponseWriter, r *http.Request) {
	_response.JSON(w, http.StatusOK, users[len(users)-1])
}

// UpdateUser - Update an User
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update an User"))
}

// DeleteUser - Delete an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
