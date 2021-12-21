package service

import "weltraumschaf.de/battleship/internal/server/model"

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
