package auto

import "escapead/projectAD/backend/api/models"

var users = []models.User{
	models.User{Nickname: "user 1", Email: "oops@youseemypassword.com", Password: "123456"},
	models.User{Nickname: "user 2", Email: "2@batman.com", Password: "123456"},
	models.User{Nickname: "user 3", Email: "3@batman.com", Password: "3333"},
}
