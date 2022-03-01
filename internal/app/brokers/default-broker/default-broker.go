package default_broker

import (
	"github.com/ivanmalyi/broker/internal/app/repository"
)

type DefaultBroker struct {
	Repo repository.Repository
}

func New(repo repository.Repository) *DefaultBroker {
	return &DefaultBroker{repo}
}

func (broker *DefaultBroker) Auth() {

}

func (broker *DefaultBroker) CreateApp() string {
	return `{"id": "broker_id"}`
}
