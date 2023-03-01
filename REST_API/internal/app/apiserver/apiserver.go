package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Apiserver
type apiserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *apiserver {
	return &apiserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Функция "Старт" сервера.
func (s *apiserver) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *apiserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil

}

func (s *apiserver) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())

}

func (s *apiserver) handleHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}

}
