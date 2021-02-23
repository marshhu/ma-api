package repository

//用户写仓储实现
type UserRepository struct {
}

func (r *UserRepository) Add(username string, password string) (id int64, err error) {
	return 0, nil
}

func (r *UserRepository) UpdateNickName(id int64, nickName string) error {
	return nil
}

func (r *UserRepository) UpdateAvatar(id int64, url string) error {
	return nil
}

func (r *UserRepository) Delete(id int64) error {
	return nil
}
