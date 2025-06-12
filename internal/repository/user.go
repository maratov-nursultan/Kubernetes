package repository

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type UserRepo struct {
	db bun.IDB
}

type UserRepository interface {
	Get(ctx context.Context, id uint) (*User, error)
	Create(ctx context.Context, user *User) (uint, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, user *User) error
}

func NewUserRepo(db bun.IDB) UserRepository {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) Get(ctx context.Context, id uint) (*User, error) {
	user := new(User)
	_, err := u.db.NewSelect().
		Model(user).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) Create(ctx context.Context, user *User) (uint, error) {
	user.CreatedAt = time.Now()
	_, err := u.db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (u *UserRepo) Delete(ctx context.Context, id uint) error {
	user := new(User)
	_, err := u.db.NewDelete().
		Model(user).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) Update(ctx context.Context, user *User) error {
	_, err := u.db.NewUpdate().Model(user).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

type User struct {
	ID        uint `bun:",pk,autoincrement"`
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
