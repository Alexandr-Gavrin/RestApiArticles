package api

import (
	"ServerAndDB2/storage"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Base API server instance discription
type API struct {
	//UNEXPORTED FIELD!
	config *Config
	logger *logrus.Logger
	router *mux.Router
	// Добавление поля для работы с хранилищем
	store *storage.Storage
}

// Api constructor: build base API instance

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start HTTP server/configures/ loggers, router, db connection and etc..
func (api *API) Start() error {
	// trying to configure logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	// подтверждение того, чтобы логгер сконфижен
	api.logger.Info("starting api server at port:", api.config.BindAddr)

	// Конфигурируем марштрктизатор
	api.configureRouterField()
	// Конфигурируем хранилище
	if err := api.cofigureStorage(); err != nil {
		return err
	}
	// На этапе валидного завершения стартурем http сервер

	return http.ListenAndServe(api.config.BindAddr, api.router)
}
