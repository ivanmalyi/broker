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

func New(config *appserver.TomlConfig) repository.Repository {
	db, err := sql.Open("postgres", config.DatabaseUrl)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &PgRepository{db: db}
}

func (repo *PgRepository) Close() {
	_ = repo.db.Close()
}

func (repo *PgRepository) Application() repository.ApplicationRepository {
	if repo.applicationRepository != nil {
		return repo.applicationRepository
	}

	repo.applicationRepository = &ApplicationRepository{
		store: repo,
	}

	return repo.applicationRepository
}
