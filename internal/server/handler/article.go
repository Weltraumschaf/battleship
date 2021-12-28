package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"weltraumschaf.de/battleship/internal/server/model"
	"weltraumschaf.de/battleship/internal/server/service"
)

type ArticleHandler struct {
	articles service.ArticleService
}

func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{
		articles: *service.NewArticleService(),
	}
}

func (h *ArticleHandler) AllArticles(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(h.articles.AllArticles())
	if err != nil {
		http.Error(w, "Can't encode data to JSON!", http.StatusInternalServerError)
		return
	}
}

func (h *ArticleHandler) SingleArticle(w http.ResponseWriter, r *http.Request) {
	id, err := retrieveId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	article, exists := h.articles.SingleArticle(id)

	if !exists {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		http.Error(w, "can't encode data to JSON", http.StatusInternalServerError)
		return
	}
}

func (h *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read request body", http.StatusInternalServerError)
		return
	}

	var article model.Article
	err = json.Unmarshal(body, &article)
	if err != nil {
		http.Error(w, "can't unmarshal request body to JSON", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(h.articles.CreateArticle(article))
	if err != nil {
		http.Error(w, "can't marshal data to JSON", http.StatusInternalServerError)
		return
	}
}

func (h *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	id, err := retrieveId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(h.articles.DeleteArticle(id))
	if err != nil {
		http.Error(w, "can't encode data to JSON", http.StatusInternalServerError)
		return
	}
}

func (h *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	id, err := retrieveId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var article model.Article
	err = json.Unmarshal(body, &article)
	if err != nil {
		http.Error(w, "can't unmarshal request body to JSON", http.StatusInternalServerError)
		return
	}

	article.Id = id
	updated, exists := h.articles.UpdateArticle(article)

	if !exists {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(updated)

	if err != nil {
		http.Error(w, "can't encode data to JSON", http.StatusInternalServerError)
		return
	}
}

