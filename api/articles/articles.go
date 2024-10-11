package articles

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sleepiinuts/simple-reddit-BE/pkg/repositories/articles"
)

type ArticleAPI struct {
	articleServ *articles.ArticlesService
}

func NewArticleAPI(as *articles.ArticlesService) *ArticleAPI {
	return &ArticleAPI{articleServ: as}
}

func (a *ArticleAPI) GetAll(c *gin.Context) {
	articles, err := a.articleServ.GetAll()
	if err != nil {
		c.Error(fmt.Errorf("[Article API] GetAll: %w", err))
	}

	c.IndentedJSON(http.StatusOK, articles)
}
