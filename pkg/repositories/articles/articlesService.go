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
		return nil, fmt.Errorf("articles-getall: %w", err)
	}
	defer rows.Close()

	var articles []models.Article

	for rows.Next() {
		var article models.Article

		err := rows.StructScan(&article)

		if err != nil {
			return nil, fmt.Errorf("article-getall-scan: %w", err)
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (a *ArticlesService) New(art *models.Article) error {
	_, err := a.repos.new(art)
	if err != nil {
		return fmt.Errorf("article-new: %w", err)
	}

	return nil
}

func (a *ArticlesService) DeleteById(id int) error {
	_, err := a.repos.deleteById(id)
	if err != nil {
		return fmt.Errorf("article-deleteById: %w", err)
	}

	return nil
}
