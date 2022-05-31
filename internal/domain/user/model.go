package user

type User struct {
	UUID     string `json:"uuid"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     int    `json:"role"`
}
