package service

import (
	"github.com/maratov-nursultan/Kubernetes/internal/manager/user"
	"github.com/maratov-nursultan/Kubernetes/internal/repository"
	"github.com/uptrace/bun"
)

type Service struct {
	userManager user.UserSDK
}

func NewService(db bun.IDB) *Service {
	userRepo := repository.NewUserRepo(db)

	return &Service{
		userManager: user.NewUserManager(userRepo),
	}
}

func (s *Service) GetUserManager() user.UserSDK {
	return s.userManager
}
