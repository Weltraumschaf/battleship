package service

import (
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

func (as *ArticleService) AllArticles() []model.Article {
	return as.data.FindAll()
}

func (as *ArticleService) SingleArticle(id uuid.UUID) (model.Article, bool) {
	return as.data.FindById(id)
}

var nullId = uuid.UUID{}

func (as *ArticleService) CreateArticle(article model.Article) model.Article {
	if article.Id == nullId {
		article.Id = uuid.Must(uuid.NewRandom())
	}

	return as.data.Save(article)
}

func (as *ArticleService) DeleteArticle(id uuid.UUID) model.Article {
	return as.data.Delete(id)
}
