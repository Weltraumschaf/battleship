package service

import (
	"encoding/json"
	"weltraumschaf.de/battleship/internal/server/repo"

	"github.com/google/uuid"
	"weltraumschaf.de/battleship/internal/server/model"
)

type ArticleService struct {
	data *repo.ArticleRepo
}

func NewArticleService() *ArticleService {
	return &ArticleService{
		data: repo.NewArticleRepo(),
	}
}

func (as *ArticleService) ReturnAllArticles() []model.Article {
	return as.data.FindAll()
}

func (as *ArticleService) ReturnSingleArticle(id uuid.UUID) (model.Article, bool) {
	return as.data.FindById(id)
}

func (as *ArticleService) CreateNewArticle(body []byte) model.Article {
	var article model.Article

	json.Unmarshal(body, &article)

	return as.data.Save(article)
}

func (as *ArticleService) DeleteArticle(id uuid.UUID) model.Article {
	return as.data.Delete(id)
}
