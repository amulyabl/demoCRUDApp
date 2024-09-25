package request

type CreateUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUser struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}
