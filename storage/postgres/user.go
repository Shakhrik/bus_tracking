package postgres

import (
	"github.com/Shakhrik/inha/bus_tracking/api/models"
	"github.com/Shakhrik/inha/bus_tracking/storage/repo"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) repo.UserRepoI {
	return userRepo{db: db}
}

func (u userRepo) Create(req models.UserCreateRequest) (res models.ResponseWithID, err error) {
	query := `INSERT INTO users (name, login, password, bus_id, user_type_id) 
			  VALUES($1, $2, $3, $4, (SELECT id FROM user_type WHERE alias = $5))
			  RETURNING id`
	err = u.db.Get(&res.ID, query, req.Name, req.Login, req.Password, req.BusID, req.Alias)
	return
}

func (u userRepo) Delete(id int32) (res models.ResponseWithID, err error) {
	query := `DELETE FROM users WHERE id = $1 RETURNING id`
	err = u.db.Get(&res.ID, query, id)
	return
}

func (u userRepo) GetAll(alias string, limit, page int32) (models.AllUsers, error) {
	var count int64
	res := []models.User{}

	offset := (page - 1) * limit

	query := `SELECT u.id, u.name, u.login, b.name AS bus_name, ut.alias
			  FROM users u
			  JOIN bus b ON u.bus_id = b.id
			  JOIN user_type ut ON u.user_type_id = ut.id
			  WHERE u.user_type_id = (SELECT id FROM user_type WHERE alias = $1)
			  LIMIT $2 OFFSET $3`
	err := u.db.Select(&res, query, alias, limit, offset)
	if err != nil {
		return models.AllUsers{Users: res}, err
	}

	queryCount := `SELECT count(1)
				   FROM users u
				   JOIN bus b ON u.bus_id = b.id
				   JOIN user_type ut ON u.user_type_id = ut.id
				   WHERE u.user_type_id = (SELECT id FROM user_type WHERE alias = $1)
				   LIMIT $2 OFFSET $3`
	err = u.db.Get(&count, queryCount, alias, limit, offset)
	if err != nil {
		return models.AllUsers{Users: res}, err

	}
	return models.AllUsers{Users: res, Count: count}, nil
}

func (u userRepo) Login(req models.UserLogin) (res models.UserInfo, err error) {
	query := `SELECT u.id, ut.alias as user_type, bus_id
			  FROM users u
			  JOIN user_type ut ON ut.id = u.user_type_id
			  WHERE login = $1 and password = $2
			  `
	err = u.db.Get(&res, query, req.Login, req.Password)
	return
}
