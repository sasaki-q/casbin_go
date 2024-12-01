package repositories

type UserType string

const (
	Member UserType = "member"
	Guest  UserType = "guest"
)

type User struct {
	ID   int
	Name string
	Type UserType
}

type IUserRepository interface {
	GetUsers() []User
	GetUser(ID int) *User
	CreateUser(user User) error
	DeleteUser(ID int) error
}

type UserRepository struct {
	users []User
}

func NewUserRepository() IUserRepository {
	return &UserRepository{
		users: []User{
			{ID: 1, Name: "member 1", Type: Member},
			{ID: 2, Name: "member 2", Type: Member},
			{ID: 3, Name: "guest 1", Type: Guest},
			{ID: 4, Name: "member 3", Type: Member},
			{ID: 5, Name: "guest 2", Type: Guest},
		},
	}
}

func (r *UserRepository) GetUsers() []User {
	return r.users
}

func (r *UserRepository) GetUser(ID int) *User {
	for _, user := range r.users {
		if user.ID == ID {
			return &user
		}
	}
	return nil
}

func (r *UserRepository) CreateUser(user User) error {
	user.ID = len(r.users) + 1
	r.users = append(r.users, user)
	return nil
}

func (r *UserRepository) DeleteUser(ID int) error {
	for i, user := range r.users {
		if user.ID == ID {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return nil
}
