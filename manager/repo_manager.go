package manager

import (
	"myGoSkeleton/repository"

	"github.com/jmoiron/sqlx"
)

type RepoManager interface {
	TesterRepo() repository.TesterRepo
}

type repoManager struct {
	SqlxDb *sqlx.DB
}

func (r *repoManager) TesterRepo() repository.TesterRepo {
	return repository.NewTesterRepo(r.SqlxDb)
}

func NewRepoManager(sqlxDb *sqlx.DB) RepoManager {
	return &repoManager{
		sqlxDb,
	}
}
