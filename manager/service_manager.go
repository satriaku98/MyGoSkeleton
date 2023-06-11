package manager

import "myGoSkeleton/services"

type ServiceManager interface {
	TesterService() services.TesterService
}
type serviceManager struct {
	repo RepoManager
}

func (u *serviceManager) TesterService() services.TesterService {
	return services.NewTesterService(u.repo.TesterRepo())
}

func NewServiceManager(repo RepoManager) ServiceManager {
	return &serviceManager{
		repo,
	}
}
