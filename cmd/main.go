package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/qustavo/dotsql"
	articlesAPI "github.com/sleepiinuts/simple-reddit-BE/api/articles"
	"github.com/sleepiinuts/simple-reddit-BE/middleware"
	articlesServ "github.com/sleepiinuts/simple-reddit-BE/pkg/repositories/articles"
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

	articlesServ := articlesServ.NewArticlesService(articlesServ.NewOracleArticlesRepos(db, dots["articles"]))
	articlesAPI := articlesAPI.NewArticleAPI(articlesServ)

	router := gin.Default()
	router.Use(middleware.ErrorHandler)
	router.GET("/articles", articlesAPI.GetAll)

	router.Run("localhost:8080")
}
