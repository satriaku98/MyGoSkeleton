package services

import (
	"myGoSkeleton/controllers/request"
	"myGoSkeleton/controllers/responses"
	"myGoSkeleton/repository"
)

type TesterService interface {
	Tester(LoginData request.TesterRequest) (responses.TesterResponse, string, error)
	Scheduler()
}

type testerService struct {
	repo repository.TesterRepo
}

func (l *testerService) Tester(TesterData request.TesterRequest) (responses.TesterResponse, string, error) {
	return l.repo.Tester(TesterData)
}
func (l *testerService) Scheduler() {
	l.repo.Scheduler()
}

func NewTesterService(repo repository.TesterRepo) TesterService {
	return &testerService{
		repo,
	}
}
