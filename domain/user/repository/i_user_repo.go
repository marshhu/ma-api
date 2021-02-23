package repository

//用户写仓储
type IUserRepo interface {
	Add(username string, password string) (id int64, err error)
	UpdateNickName(id int64, nickName string) error
	UpdateAvatar(id int64, url string) error
	Delete(id int64) error
}
