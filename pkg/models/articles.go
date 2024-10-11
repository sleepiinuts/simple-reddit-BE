package models

type Article struct {
	ID    string `db:"ID"`
	Title string `db:"TITLE"`
	URL   string `db:"URL"`
	Point int    `db:"POINT"`
}
