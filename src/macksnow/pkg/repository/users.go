package repository

import (
	"database/sql"
	"macksnow/pkg/model"
)

func (repo *repository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User

	if err := repo.db.Get(&user,
		"SELECT * FROM users WHERE email = ?",
		email,
	); err != nil {
		// 存在しないEmailの場合、no rows in result setがerrとして返却される
		return nil, err
	}

	return &user, nil
}

func (repo *repository) FindUserById(id int) (*model.User, error) {
	var user model.User

	if err := repo.db.Get(&user,
		"SELECT * FROM users WHERE id = ?",
		id,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *repository) CreateUser(nickName string, screenName string, email string, passwordHash string, confirmationToken string) (*model.User, error) {
	var result sql.Result
	var err error

	if result, err = repo.db.Exec(
		"INSERT INTO users(nickname, screen_name, email, password_hash, confirmation_token) VALUES(?, ?, ?, ?, ?)",
		nickName, screenName, email, passwordHash, confirmationToken,
	); err != nil {
		return nil, err
	}

	var lastInsertedId int64

	if lastInsertedId, err = result.LastInsertId(); err != nil {
		return nil, err
	}

	var id = int(lastInsertedId)
	var user *model.User

	if user, err = repo.FindUserById(int(id)); err != nil {
		return nil, err
	}

	return user, nil
}
