package articles

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/sleepiinuts/simple-reddit-BE/pkg/models"
)

type ArticlesRepos interface {
	getAll() (*sqlx.Rows, error)
	new(a *models.Article) (sql.Result, error)
	deleteById(id int) (sql.Result, error)
}
