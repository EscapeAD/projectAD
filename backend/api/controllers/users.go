package controllers

import (
	_response "escapead/projectAD/backend/api/responses"
	"net/http"
)

// GetUsers - GET all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	_response.JSON(w, http.StatusOK, map[string]string{})
}

// CreateUser - GET user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	_response.JSON(w, http.StatusOK, map[string]string{})
}

// GetUser - GET an user
func GetUser(w http.ResponseWriter, r *http.Request) {
	_response.JSON(w, http.StatusOK, map[string]string{})
}

// UpdateUser - Update an User
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update an User"))
}

// DeleteUser - Delete an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
