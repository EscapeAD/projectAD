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
	"strconv"

	"github.com/go-chi/chi"
)

// GetUsers - GET all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		_response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	repo := crud.NewRepositoryUsersCRUD(db)

	func(userRepository repository.UserRepository) {
		users, err := userRepository.FindAll()
		if err != nil {
			_response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, len(users)))
		_response.JSON(w, http.StatusCreated, users)
	}(repo)
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
	userIDParam := chi.URLParam(r, "id") // from a route like /users/{id}
	userID, err := strconv.ParseUint(userIDParam, 10, 32)

	if err != nil {
		_response.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		_response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	repo := crud.NewRepositoryUsersCRUD(db)

	func(userRepository repository.UserRepository) {
		user, err := userRepository.FindByID(uint32(userID))
		if err != nil {
			_response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, user))
		_response.JSON(w, http.StatusCreated, user)
	}(repo)
}

// UpdateUser - Update an User
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userIDParam := chi.URLParam(r, "id") // from a route like /users/{id}
	userID, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		_response.Error(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		_response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	repo := crud.NewRepositoryUsersCRUD(db)

	func(userRepository repository.UserRepository) {
		rows, err := userRepository.Update(uint32(userID), user)
		if err != nil {
			_response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, rows))
		_response.JSON(w, http.StatusCreated, rows)
	}(repo)
}

// DeleteUser - Delete an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
