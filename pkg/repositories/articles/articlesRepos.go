package articles

import "github.com/jmoiron/sqlx"

type ArticlesRepos interface {
	getAll() (*sqlx.Rows, error)
}
