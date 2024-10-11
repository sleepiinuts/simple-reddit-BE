package main

import (
	"fmt"
	"log"

	"github.com/qustavo/dotsql"
	"github.com/sleepiinuts/simple-reddit-BE/pkg/repositories/articles"
)

var dots map[string]*dotsql.DotSql

func main() {

	fmt.Println("Starting Simple-Reddit-BE")
	db := conn()
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("PING error: ", err)
	}

	prepSqlLoader()

	articlesServ := articles.NewArticlesService(articles.NewOracleArticlesRepos(db, dots["articles"]))
	fmt.Println(articlesServ.GetAll())
}
