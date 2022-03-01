package pgrepository

import (
	"github.com/ivanmalyi/broker/internal/app/model"
)

type ApplicationRepository struct {
	store *PgRepository
}

func (applicationRepository *ApplicationRepository) Create(application *model.Application) error {
	return nil
}

func (applicationRepository *ApplicationRepository) Find(id int) (*model.Application, error) {
	return nil, nil
}
