package api

import (
	"ServerAndDB2/internal/app/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

// Full API Handlers init file
func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func (api *API) GetAllArticles(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get all articles GET /api/v1/articles")
	articles, err := api.store.Article().SelectAll()
	if err != nil {
		api.logger.Info("Error while GetAllArticles : ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing to DB",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(articles)
}

func (api *API) GetArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get article by ID /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing ID param: ", err)
		msg := Message{
			Message:    "Invalid form of data",
			StatusCode: 400,
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	article, ok, err := api.store.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing to DB")
		msg := Message{
			Message:    "We have some troubles to accessing to DB",
			StatusCode: 500,
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Cant find article by this ID: ", id)
		msg := Message{
			Message:    "Cant find articles",
			StatusCode: 404,
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(article)

}

func (api *API) DeleteArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete Article by ID DELETE api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing ID param: ", err)
		msg := Message{
			Message:    "Invalid form of data",
			StatusCode: 400,
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok, err := api.store.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing to DB")
		msg := Message{
			Message:    "We have some troubles to accessing to DB",
			StatusCode: 500,
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Cant find article by this ID: ", id)
		msg := Message{
			Message:    "Cant find articles",
			StatusCode: 404,
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, err = api.store.Article().DeleteById(id)
	if err != nil {
		api.logger.Info("Troubles with deleting")
		msg := Message{
			Message:    "Some problems with deleting",
			StatusCode: 501,
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(202)
	msg := Message{
		Message:    "Kaif ydalili)))",
		StatusCode: 202,
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}

func (api *API) PostArticle(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Article POST /api/v1/articles")
	var article models.Article
	err := json.NewDecoder(req.Body).Decode(&article)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	a, err := api.store.Article().Create(&article)
	if err != nil {
		api.logger.Info("Error while creating new Article : ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some errors with accesssing DB",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)

}

func (api *API) PostUserRegister(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post USER register POST /api/v1/user/register")
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid data from user")
		msg := Message{
			Message:    "Invalid data for user",
			StatusCode: 400,
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, ok, err := api.store.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Troubles while accessing to DB")
		msg := Message{
			Message:    "We have some troubles to accessing to DB",
			StatusCode: 500,
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if ok {
		api.logger.Info("User with that already exists")
		msg := Message{
			Message:    "User with that ID alredy exists in DB",
			StatusCode: 400,
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	userAdded, err := api.store.User().Create(&user)
	if err != nil {
		api.logger.Info("Troubles while accessing to DB")
		msg := Message{
			Message:    "We have some troubles to accessing to DB",
			StatusCode: 500,
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("User {login:%s} successfully added", userAdded.Login),
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)

}
