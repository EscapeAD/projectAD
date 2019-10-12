package controllers

import (
	"encoding/json"
	"escapead/projectAD/backend/api/database"
	"escapead/projectAD/backend/api/models"
	"escapead/projectAD/backend/api/repository"
	"escapead/projectAD/backend/api/repository/crud"
	_response "escapead/projectAD/backend/api/responses"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetUsers - GET all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	_response.JSON(w, http.StatusOK, map[string]string{})
}

// CreateUser - GET user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		_response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		_response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	repo := crud.NewRepositoryUsersCRUD(db)
	func(userRepository repository.UserRepository) {
		user, err = userRepository.Save(user)
		if err != nil {
			_response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, user.ID))
		_response.JSON(w, http.StatusCreated, user)
	}(repo)
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
