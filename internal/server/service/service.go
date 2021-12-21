package service

import (
	"encoding/json"
	"weltraumschaf.de/battleship/internal/server/model"
)

type ArticleService struct {
	data []model.Article
}

func NewArticleService() *ArticleService {
	return &ArticleService{
		data: []model.Article{
			{Id: 1, Title: "Hello", Description: "Article Description", Content: "Article Content"},
			{Id: 2, Title: "Hello 2", Description: "Article Description", Content: "Article Content"},
		},
	}
}

func (as *ArticleService) ReturnAllArticles() []model.Article {
	return as.data
}

func (as *ArticleService) ReturnSingleArticle(id int) model.Article {
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
	as.data = append(as.data, article)

	return article
}

func (as *ArticleService) DeleteArticle(id int) model.Article {
	for index, article := range as.data {
		if article.Id == id {
			as.data = append(as.data[:index], as.data[index + 1])
			return article
		}
	}

	return model.Article{}
}
