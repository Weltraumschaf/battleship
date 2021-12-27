package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
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
	json.NewEncoder(w).Encode(h.articles.ReturnAllArticles())
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

	json.NewEncoder(w).Encode(h.articles.ReturnSingleArticle(id))
}

func (h *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.NewEncoder(w).Encode(h.articles.CreateNewArticle(reqBody))
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

	json.NewEncoder(w).Encode(h.articles.DeleteArticle(id))
}

func (h *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented!")
}
