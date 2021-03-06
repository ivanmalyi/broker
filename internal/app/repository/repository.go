package repository

import "github.com/ivanmalyi/broker/internal/app/model"

type Repository interface {
	Application() ApplicationRepository
	Close()
}

type ApplicationRepository interface {
	Create(*model.Application) error
	Find(int) (*model.Application, error)
}
