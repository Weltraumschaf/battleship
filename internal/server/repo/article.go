package repo

import (
	"github.com/google/uuid"
	"log"
	"weltraumschaf.de/battleship/internal/server/model"
)

type ArticleRepo struct {
	data map[uuid.UUID]model.Article
}

func NewArticleRepo() *ArticleRepo {
	r := &ArticleRepo{
		data: make(map[uuid.UUID]model.Article),
	}

	for _, article := range generateTestData() {
		r.Save(article)
	}

	return r
}

func (r *ArticleRepo) Save(article model.Article) model.Article {
	log.Println("Saving article for id", article.Id, "with data:", article)
	r.data[article.Id] = article
	return article
}

func (r *ArticleRepo) Delete(id uuid.UUID) model.Article {
	article, exists := r.FindById(id)

	if exists {
		delete(r.data, id)
	}

	return article
}

func (r *ArticleRepo) FindAll() []model.Article {
	articles := make([]model.Article, 0, len(r.data))

	for _, article := range r.data {
		articles = append(articles, article)
	}

	return articles
}

func (r *ArticleRepo) FindById(id uuid.UUID) (model.Article, bool) {
	article, exists := r.data[id]
	return article, exists
}

func generateTestData() []model.Article {
	return []model.Article{
		{
			Id: uuid.Must(uuid.NewRandom()),
			Title: "Hello, World!",
			Description: "Lorem ipsum dolor sit amet",
			Content: "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor.",
		},
		{
			Id: uuid.Must(uuid.NewRandom()),
			Title: "Another One",
			Description: "Aenean commodo ligula eget",
			Content: "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor.",
		},
	}
}
