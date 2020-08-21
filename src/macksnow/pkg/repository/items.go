package repository

import (
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
