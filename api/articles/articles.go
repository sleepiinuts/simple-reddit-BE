package articles

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sleepiinuts/simple-reddit-BE/pkg/models"
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

func (a *ArticleAPI) New(c *gin.Context) {
	var article models.Article
	err := c.BindJSON(&article)
	if err != nil {
		c.Error(fmt.Errorf("[Article API] New - BindJSON: %w", err))
		c.Status(http.StatusBadRequest)
		return
	}

	err = a.articleServ.New(&article)
	if err != nil {
		c.Error(fmt.Errorf("[Article API] New - ArticleServ: %w", err))
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func (a *ArticleAPI) DeleteById(c *gin.Context) {
	var id int

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.Error(fmt.Errorf("[Article API] DeleteById - bad request: %v", id))
		c.Status(http.StatusBadRequest)
		return
	}

	err = a.articleServ.DeleteById(id)
	if err != nil {
		c.Error(fmt.Errorf("[Article API] DeleteById - ArticleServ: %w", err))
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (a *ArticleAPI) Vote(c *gin.Context) {
	var (
		id, vote int
		err      error
	)

	id, err = strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.Error(fmt.Errorf("[Article API] Vote[id param] - bad request: %v", id))
		c.Status(http.StatusBadRequest)
		return
	}

	err = c.BindJSON(&vote)
	if err != nil {
		c.Error(fmt.Errorf("[Article API] Vote[body] - bad request: %w", err))
		c.Status(http.StatusBadRequest)
		return
	}

	err = a.articleServ.Vote(id, vote)
	if err != nil {
		c.Error(fmt.Errorf("[Article API] Vote - ArticleServ: %w", err))
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
