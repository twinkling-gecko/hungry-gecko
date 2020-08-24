package repository

import (
	"database/sql"
	"macksnow/pkg/model"
)

func (repo *repository) AllItems() ([]*model.Item, error) {
	var items []*model.Item

	if err := repo.db.Select(&items,
		"SELECT * FROM items;",
	); err != nil {
		return nil, err
	}

	// 1件もデータがない場合echoはnilを返す
	if items == nil {
		items = []*model.Item{}
	}

	return items, nil
}

func (repo *repository) FindItem(id int) (*model.Item, error) {
	var item model.Item

	if err := repo.db.Get(&item,
		"SELECT * FROM items WHERE id = ?",
		id,
	); err != nil {
		// 存在しないIDの場合、no rows in result setがerrとして返却される
		return nil, err
	}

	return &item, nil
}

func (repo *repository) CreateItem(name string, summary string, uri string) (*model.Item, error) {
	var result sql.Result
	var err error

	if result, err = repo.db.Exec(
		"INSERT INTO items(name, summary, uri) VALUES(?, ?, ?)",
		name, summary, uri,
	); err != nil {
		return nil, err
	}

	var lastInsertedId int64

	if lastInsertedId, err = result.LastInsertId(); err != nil {
		return nil, err
	}

	var id = int(lastInsertedId)
	var item *model.Item

	if item, err = repo.FindItem(int(id)); err != nil {
		return nil, err
	}

	return item, nil
}
