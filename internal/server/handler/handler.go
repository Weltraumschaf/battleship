package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

func (h *ArticleHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func (h *ArticleHandler) ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(h.articles.ReturnAllArticles())
}

func (h *ArticleHandler) ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
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

func (h *ArticleHandler) CreateNewArticle(w http.ResponseWriter, r *http.Request) {
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
