package models

import "database/sql"

type Items struct {
	ItemID         int
	Description    string
	Sugar          int
	Protein        int
	Carbohydrates  int
	AvgCookingTime int
}

type ItemModel struct {
	DB *sql.DB
}
