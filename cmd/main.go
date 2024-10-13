package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/gin-contrib/cors"
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

	// create logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	router := gin.Default()
	router.Use(middleware.ErrorHandler(logger))
	router.Use(cors.Default())
	router.GET("/articles", articlesAPI.GetAll)
	router.POST("/articles", articlesAPI.New)
	router.DELETE("/articles/:id", articlesAPI.DeleteById)
	router.PATCH("/articles/:id", articlesAPI.Vote)

	router.Run("localhost:8080")
}
