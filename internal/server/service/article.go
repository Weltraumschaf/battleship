package service

import (
	"log"
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

func (s *ArticleService) AllArticles() []model.Article {
	return s.data.FindAll()
}

func (s *ArticleService) SingleArticle(id uuid.UUID) (model.Article, bool) {
	return s.data.FindById(id)
}

var nullId = uuid.UUID{}

func (s *ArticleService) CreateArticle(article model.Article) model.Article {
	if article.Id == nullId {
		article.Id = uuid.Must(uuid.NewRandom())
	}

	s.data.Save(article)

	return article
}

func (s *ArticleService) DeleteArticle(id uuid.UUID) model.Article {
	return s.data.Delete(id)
}

func (s *ArticleService) UpdateArticle(toUpdate model.Article) (model.Article, bool) {
	log.Println("Update model: ", toUpdate)
	persisted, exists := s.data.FindById(toUpdate.Id)

	if !exists {
		log.Println("No persistent data found!")
		return model.Article{}, false
	}

	if toUpdate.Title != "" {
		log.Println("Update title.")
		persisted.Title = toUpdate.Title
	}

	if toUpdate.Description != "" {
		log.Println("Update description.")
		persisted.Description = toUpdate.Description
	}

	if toUpdate.Content != "" {
		log.Println("Update content.")
		persisted.Content = toUpdate.Content
	}

	s.data.Save(persisted)
	return persisted, true
}
