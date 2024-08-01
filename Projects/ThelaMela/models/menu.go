package models

import "database/sql"

type Menu struct {
	ItemID   int
	ItemName string
	Price    int
}

type MenuModel struct {
	DB *sql.DB
}

func (m MenuModel) CreateItem(name string, price int) error {
	return nil
}
func (m MenuModel) DeleteItem(id int) error {
	return nil
}
func (m MenuModel) UpdateItem(id, price int) error {
	return nil
}
