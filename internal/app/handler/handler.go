package handler

import (
	"github.com/gorilla/mux"
	"github.com/ivanmalyi/broker/internal/app/brokers"
	"github.com/ivanmalyi/broker/internal/app/repository"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Handler for http requests
type Handler struct {
	mux    *mux.Router
	logger *logrus.Logger
	repo   repository.Repository
}

// New http handler
func New(mux *mux.Router, logger *logrus.Logger, repo repository.Repository) *Handler {
	handler := Handler{mux, logger, repo}
	handler.registerRoutes()

	return &handler
}

// RegisterRoutes for all http endpoints
func (h *Handler) registerRoutes() {
	h.mux.HandleFunc("/broker/orders/{broker_key}", h.createApplication).Methods(http.MethodGet)
	h.mux.HandleFunc("/default/{broker_key}", h.createApplication).Methods(http.MethodGet)
}

func (h *Handler) createApplication(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	brokerKey, ok := vars["broker_key"]
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	broker := brokers.GetBroker(brokerKey, h.repo)
	response := broker.CreateApp()
	writer.Write([]byte(response))
}
