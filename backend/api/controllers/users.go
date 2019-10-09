package controllers

import (
	"net/http"
)

// GetUsers - GET all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List all Users"))
}

// CreateUser - GET user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create an User"))
}

// GetUser - GET an user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("An User"))
}

// UpdateUser - Update an User
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update an User"))
}

// DeleteUser - Delete an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
