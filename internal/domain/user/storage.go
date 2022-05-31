package user

type Storage interface {
	GetOne(uuid string) *User
	GetAll(limit, offset int) []*User
	Create(user *User) *User
	Delete(user *User) error
}
