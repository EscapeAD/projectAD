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

// CreatePost - GET post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		_response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	post.Prepare()
	err = post.Validate()
	if err != nil {
		_response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		_response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)
	func(postRepository repository.PostsRepository) {
		post, err = postRepository.Save(post)
		if err != nil {
			_response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, post.ID))
		_response.JSON(w, http.StatusCreated, post)
	}(repo)
}

// GetPosts - GET all posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		_response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)

	func(postRepository repository.PostsRepository) {
		posts, err := postRepository.FindAll()
		if err != nil {
			_response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, len(posts)))
		_response.JSON(w, http.StatusCreated, posts)
	}(repo)
}

// GetPost - GET an post
func GetPost(w http.ResponseWriter, r *http.Request) {
	postIDParam := chi.URLParam(r, "id") // from a route like /posts/{id}
	postID, err := strconv.ParseUint(postIDParam, 10, 32)

	if err != nil {
		_response.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		_response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)

	func(postRepository repository.PostsRepository) {
		post, err := postRepository.FindByID(uint64(postID))
		if err != nil {
			_response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, post))
		_response.JSON(w, http.StatusCreated, post)
	}(repo)
}

// UpdatePost - Update an Post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	postIDParam := chi.URLParam(r, "id") // from a route like /posts/{id}
	postID, err := strconv.ParseUint(postIDParam, 10, 32)
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
	post := models.Post{}
	err = json.Unmarshal(body, &post)
	repo := crud.NewRepositoryPostsCRUD(db)

	func(postRepository repository.PostsRepository) {
		rows, err := postRepository.Update(uint64(postID), post)
		if err != nil {
			_response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, rows))
		_response.JSON(w, http.StatusCreated, rows)
	}(repo)
}

// DeletePost - Delete an post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	postIDParam := chi.URLParam(r, "id") // from a route like /posts/{id}
	postID, err := strconv.ParseUint(postIDParam, 10, 32)
	if err != nil {
		_response.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		_response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)

	func(postRepository repository.PostsRepository) {
		_, err := postRepository.Delete(uint64(postID))
		if err != nil {
			_response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Entity", fmt.Sprintf("%d", postID))
		_response.JSON(w, http.StatusNoContent, "")
	}(repo)
}
