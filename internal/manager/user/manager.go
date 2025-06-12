package user

import (
	"context"
	"github.com/maratov-nursultan/Kubernetes/internal/model"
	"github.com/maratov-nursultan/Kubernetes/internal/repository"
)

type User struct {
	userRepo repository.UserRepository
}

type UserSDK interface {
	GetUser(ctx context.Context, req *model.GetUserRequest) (*model.GetUserResponse, error)
	CreateUser(ctx context.Context, req *model.CreateUserRequest) (uint, error)
	Delete(ctx context.Context, req *model.DeleteUserRequest) error
	Update(ctx context.Context, req *model.UpdateUserRequest) error
}

func NewUserManager(userRepo repository.UserRepository) UserSDK {
	return &User{
		userRepo: userRepo,
	}
}

func (u *User) GetUser(ctx context.Context, req *model.GetUserRequest) (*model.GetUserResponse, error) {
	user, err := u.userRepo.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &model.GetUserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}

func (u *User) CreateUser(ctx context.Context, req *model.CreateUserRequest) (uint, error) {
	id, err := u.userRepo.Create(ctx, &repository.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *User) Delete(ctx context.Context, req *model.DeleteUserRequest) error {
	err := u.userRepo.Delete(ctx, req.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Update(ctx context.Context, req *model.UpdateUserRequest) error {
	err := u.userRepo.Update(ctx, &repository.User{
		ID:        req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		return err
	}

	return nil
}
