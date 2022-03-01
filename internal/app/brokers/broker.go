package brokers

import (
	"github.com/ivanmalyi/broker/internal/app/brokers/default-broker"
	"github.com/ivanmalyi/broker/internal/app/brokers/technodom"
	"github.com/ivanmalyi/broker/internal/app/repository"
)

type Broker interface {
	Auth()
	CreateApp() string
}

var brokers = map[string]Broker{}

const (
	Technodom = "technodom"
)

func GetBroker(brokerKey string, repo repository.Repository) Broker {
	broker, ok := brokers[brokerKey]
	if !ok {
		broker = brokerFactory(brokerKey, repo)
	}

	return broker
}

func brokerFactory(brokerKey string, repo repository.Repository) Broker {
	var broker Broker
	switch brokerKey {
	case Technodom:
		broker = technodom.New(repo)
		brokers[brokerKey] = broker
		break
	default:
		broker = default_broker.New(repo)
		brokers[brokerKey] = broker
		break
	}

	return broker
}
