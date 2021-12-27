package service

import (
	"encoding/json"

	"github.com/google/uuid"
	"weltraumschaf.de/battleship/internal/server/model"
)

type ArticleService struct {
	data []model.Article
}

func NewArticleService() *ArticleService {
	return &ArticleService{
		data: []model.Article{
			{
				Id: uuid.Must(uuid.NewRandom()),
				Title: "Hello",
				Description: "Article Description",
				Content: "Article Content",
			},
			{
				Id: uuid.Must(uuid.NewRandom()),
				Title: "Hello 2",
				Description: "Article Description",
				Content: "Article Content",
			},
		},
	}
}

func (as *ArticleService) ReturnAllArticles() []model.Article {
	return as.data
}

func (as *ArticleService) ReturnSingleArticle(id uuid.UUID) model.Article {
	for _, article := range as.data {
		if article.Id == id {
			return article
		}
	}

	return model.Article{}
}

func (as *ArticleService) CreateNewArticle(body []byte) model.Article {
	var article model.Article
	json.Unmarshal(body, &article)
	article.Id = uuid.New()
	as.data = append(as.data, article)

	return article
}

func (as *ArticleService) DeleteArticle(id uuid.UUID) model.Article {
	for index, article := range as.data {
		if article.Id == id {
			as.data = append(as.data[:index], as.data[index + 1])
			return article
		}
	}

	return model.Article{}
}
