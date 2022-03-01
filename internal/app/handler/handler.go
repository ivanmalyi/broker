package handler

import (
	"github.com/gorilla/mux"
	"github.com/ivanmalyi/broker/internal/app/repository"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Handler for http requests
type Handler struct {
	mux    *mux.Router
	logger *logrus.Logger
	repo   *repository.Repository
}

// New http handler
func New(mux *mux.Router, logger *logrus.Logger, repo *repository.Repository) *Handler {
	handler := Handler{mux, logger, repo}
	handler.registerRoutes()

	return &handler
}

// RegisterRoutes for all http endpoints
func (h *Handler) registerRoutes() {
	h.mux.HandleFunc("/", h.hello()).Methods(http.MethodGet)
}

func (h *Handler) hello() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		writer.WriteHeader(200)
		writer.Write([]byte("Hello World"))
	}
}
