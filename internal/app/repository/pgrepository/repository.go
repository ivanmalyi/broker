package pgrepository

import (
	"database/sql"
	"github.com/ivanmalyi/broker/internal/app/appserver"
	"github.com/ivanmalyi/broker/internal/app/repository"
	_ "github.com/lib/pq"
)

type PgRepository struct {
	db                    *sql.DB
	applicationRepository *ApplicationRepository
}

func New(config *appserver.TomlConfig) *PgRepository {
	var repo *PgRepository
	db, err := sql.Open("postgres", config.DatabaseUrl)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	repo.db = db

	return repo
}

func (repository *PgRepository) Application() repository.ApplicationRepository {
	if repository.applicationRepository != nil {
		return repository.applicationRepository
	}

	repository.applicationRepository = &ApplicationRepository{
		store: repository,
	}

	return repository.applicationRepository
}
