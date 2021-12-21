package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
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

var articles []model.Article = []model.Article{
	{Id: 1, Title: "Hello", Description: "Article Description", Content: "Article Content"},
	{Id: 2, Title: "Hello 2", Description: "Article Description", Content: "Article Content"},
}

func (h *ArticleHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func (h *ArticleHandler) ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func (h *ArticleHandler) ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, article := range articles {
		if article.Id == id {
			json.NewEncoder(w).Encode(article)
			break
		}
	}
}

func (h *ArticleHandler) CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article model.Article
	json.Unmarshal(reqBody, &article)
	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)
}

func (h *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, article := range articles {
		if article.Id == id {
			articles = append(articles[:index], articles[index+1:]...)
		}
	}
}
