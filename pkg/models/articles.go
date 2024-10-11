package models

type Article struct {
	ID    string `db:"ID" json:"id"`
	Title string `db:"TITLE" json:"title"`
	URL   string `db:"URL" json:"url"`
	Point int    `db:"POINT" json:"point"`
}
