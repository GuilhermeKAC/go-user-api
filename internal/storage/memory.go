package storage

import "github.com/GuilhermeKAC/go-user-api/internal/models"

var users = make(map[string]models.User)

func SaveUser(user models.User) {
	users[user.ID] = user
}

func GetAllUsers() []models.User {
	list := []models.User{}
	for _, user := range users {
		list = append(list, user)
	}
	return list
}

func GetUserByID(id string) (models.User, bool) {
	user, exists := users[id]
	return user, exists
}

func UpdateUser(id string, user models.User) {
	users[id] = user
}

func DeleteUser(id string) {
	delete(users, id)
}
