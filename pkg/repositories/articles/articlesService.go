package articles

import (
	"fmt"

	models "github.com/sleepiinuts/simple-reddit-BE/pkg/models"
)

type ArticlesService struct {
	repos ArticlesRepos
}

func NewArticlesService(repos ArticlesRepos) *ArticlesService {
	return &ArticlesService{
		repos: repos,
	}
}

func (a *ArticlesService) GetAll() ([]models.Article, error) {
	rows, err := a.repos.getAll()
	if err != nil {
		return nil, fmt.Errorf("articles: %w", err)
	}
	defer rows.Close()

	var articles []models.Article

	for rows.Next() {
		var article models.Article

		err := rows.StructScan(&article)

		if err != nil {
			return nil, fmt.Errorf("article-scan: %w", err)
		}
		articles = append(articles, article)
	}
	return articles, nil
}
