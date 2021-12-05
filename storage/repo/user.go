package repo

import "github.com/Shakhrik/inha/bus_tracking/api/models"

type UserRepoI interface {
	Create(req models.UserCreateRequest) (res models.ResponseWithID, err error)
	Delete(id int32) (res models.ResponseWithID, err error)
	GetAll(alias string, limit, page int32) (models.AllUsers, error)
	Login(req models.UserLogin) (res models.UserInfo, err error)
}
