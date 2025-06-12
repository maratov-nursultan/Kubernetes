package model

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type GetUserRequest struct {
	ID uint `json:"id"`
}

type UpdateUserRequest struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type DeleteUserRequest struct {
	ID uint `json:"id"`
}
