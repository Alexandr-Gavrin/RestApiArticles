package api

import (
	"ServerAndDB2/storage"

	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
)

// Пытаемся отконфигурировать наш api instance (а конкретнее - поле logger)
func (a *API) configureLoggerField() error {
	logLevel, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(logLevel)
	return nil
}

// Пытаемся сконфигурировать маршрутизатор (а конктретнее - поле routyer)

func (a *API) configureRouterField() {
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleById).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticleById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/articles", a.PostArticle).Methods("POST")
	a.router.HandleFunc(prefix+"/user/register", a.PostUserRegister).Methods("POST")
}

// Пытаемся конфигурировать наше хранилище
func (a *API) cofigureStorage() error {
	storage := storage.New(a.config.Storage)
	// Пытаемся установить соединение. если невозможно - возвращаем ошибку
	if err := storage.Open(); err != nil {
		return err
	}
	a.store = storage
	return nil
}
