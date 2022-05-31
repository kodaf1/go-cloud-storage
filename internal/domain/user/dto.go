package user

type SignUpDTO struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"repeat_password"`
}

type SignInDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
