package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"macksnow/pkg/model"
)

// DBへの接続を管理する実態
type repository struct {
	db *sqlx.DB
}

// 実際に振る舞いを提供するinterface
type Repository interface {
	AllItems() ([]*model.Item, error)
	FindItem(id int) (*model.Item, error)
	Close() error
}

// コンストラクタに当たるもの（repositoryの生成を行う）
func New() (Repository, error) {
	db, err := sqlx.Connect("mysql", "giant:leopard@tcp(giant)/macksnow?parseTime=true")
	if err != nil {
		return nil, err
	}

	return &repository{db: db}, nil
}

func (repo *repository) Close() error {
	if err := repo.db.Close(); err != nil {
		return err
	}

	return nil
}
