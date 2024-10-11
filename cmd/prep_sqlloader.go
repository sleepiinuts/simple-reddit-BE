package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/qustavo/dotsql"
)

func prepSqlLoader() {
	dots = make(map[string]*dotsql.DotSql)
	basePath, _ := os.Getwd()
	basePath = filepath.Dir(basePath) + "/pkg/repositories"

	if dot, err := dotsql.LoadFromFile(basePath + "/articles/articles.sql"); true {
		if err != nil {
			log.Fatal("Articles sql loader error: ", err)
		}

		dots["articles"] = dot
	}

}
