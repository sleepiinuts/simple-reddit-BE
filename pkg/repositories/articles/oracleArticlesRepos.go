package articles

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/qustavo/dotsql"
	"github.com/sleepiinuts/simple-reddit-BE/pkg/models"
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

func (o *OracleArticlesRepos) new(a *models.Article) (sql.Result, error) {
	stmt, err := o.dot.Raw("New")
	if err != nil {
		return nil, fmt.Errorf("[Oracle-new]: %w", err)
	}

	return o.db.Exec(stmt, a.Title, a.URL, a.Point)
}

// deleteById implements ArticlesRepos.
func (o *OracleArticlesRepos) deleteById(id int) (sql.Result, error) {
	stmt, err := o.dot.Raw("DeleteById")
	if err != nil {
		return nil, fmt.Errorf("[Oracle-deleteById]: %w", err)
	}

	return o.db.Exec(stmt, id)
}

// vote implements ArticlesRepos.
func (o *OracleArticlesRepos) vote(id int, vote int) (sql.Result, error) {
	stmt, err := o.dot.Raw("Vote")
	if err != nil {
		return nil, fmt.Errorf("[Oracle-vote]: %w", err)
	}

	return o.db.Exec(stmt, vote, id)
}

var _ ArticlesRepos = &OracleArticlesRepos{}
