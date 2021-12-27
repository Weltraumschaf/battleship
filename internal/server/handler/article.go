package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
	literalId, ok := vars["id"]

	if !ok {
		http.Error(w, "Path parameter 'id' missing!", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(literalId)
	if err != nil {
		http.Error(w, "Malformed parameter 'id'!", http.StatusBadRequest)
		return
	}

	article, exists := h.articles.SingleArticle(id)

	if !exists {
		http.Error(w, "Not found", http.StatusNotFound)
	}

	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		http.Error(w, "Can't encode data to JSON!", http.StatusInternalServerError)
		return
	}
}

func (h *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read request body!", http.StatusInternalServerError)
		return
	}

	var article model.Article
	err = json.Unmarshal(body, &article)
	if err != nil {
		http.Error(w, "Can't unmarshal request body to JSON!", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(h.articles.CreateArticle(article))
	if err != nil {
		http.Error(w, "Can't marshal data to JSON!", http.StatusInternalServerError)
		return
	}
}

func (h *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	literalId, ok := vars["id"]

	if !ok {
		http.Error(w, "Get parameter 'id' missing!", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(literalId)
	if err != nil {
		http.Error(w, "Malformed parameter 'id'!", http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(h.articles.DeleteArticle(id))
	if err != nil {
		http.Error(w, "Can't encode data to JSON!", http.StatusInternalServerError)
		return
	}
}

func (h *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented!")
}
