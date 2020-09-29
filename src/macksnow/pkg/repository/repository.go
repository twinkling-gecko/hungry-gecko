package repository

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"macksnow/pkg/model"
)

var dsn = os.Getenv("DSN")

// DBへの接続を管理する実体
type repository struct {
	db *sqlx.DB
}

// 実際に振る舞いを提供するinterface
type Repository interface {
	AllItems() ([]*model.Item, error)
	FindItem(id int) (*model.Item, error)
	CreateItem(name string, summary string, uri string) (*model.Item, error)
	UpdateItem(name string, summary string, uri string, id int) (*model.Item, error)
	DeleteItem(id int) error
	Close() error
}

// コンストラクタに当たるもの（repositoryの生成を行う）
func New() (Repository, error) {
	db, err := sqlx.Connect("mysql", dsn)
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
