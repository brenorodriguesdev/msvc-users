package dtos

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
