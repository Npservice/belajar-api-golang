package response

import (
	"perpustakaan/model/domain"
)

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func ResponseUser(User domain.User) UserResponse {
	formatter := UserResponse{}
	formatter.ID = User.Id
	formatter.Name = User.Name
	formatter.Username = User.Username
	return formatter
}

func ResponseUsers(User []domain.User) []UserResponse {
	var usersResponse []UserResponse

	for _, users := range User {
		formatters := ResponseUser(users)
		usersResponse = append(usersResponse, formatters)
	}
	return usersResponse
}
