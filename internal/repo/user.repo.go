package repo

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetInfoUser(id int) string {
	return "user info"
}
