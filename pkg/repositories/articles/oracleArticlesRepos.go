package articles

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/qustavo/dotsql"
)

type OracleArticlesRepos struct {
	db  *sqlx.DB
	dot *dotsql.DotSql
}

func NewOracleArticlesRepos(db *sqlx.DB, dot *dotsql.DotSql) *OracleArticlesRepos {
	return &OracleArticlesRepos{db: db, dot: dot}
}

// getAll implements ArticlesRepos.
func (o *OracleArticlesRepos) getAll() (*sqlx.Rows, error) {
	stmt, err := o.dot.Raw("GetAll")
	if err != nil {
		return nil, fmt.Errorf("[Oracle-getAll]: %w", err)
	}

	return o.db.Queryx(stmt)
}

var _ ArticlesRepos = &OracleArticlesRepos{}
