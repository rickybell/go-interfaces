package repositories

import (
	"context"
	"errors"

	"github.com/rickybell/go-interfaces/app/interfaces"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExist     = errors.New("row does not exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type Repository[T interfaces.Model] interface {
	Migrate(ctx context.Context) error
	ClearAll(ctx context.Context) error
	Create(ctx context.Context, model T) (*T, error)
	All(ctx context.Context) ([]T, error)
	GetByName(ctx context.Context, name string) (*T, error)
	Update(ctx context.Context, id int64, updated T) (*T, error)
	Delete(ctx context.Context, id int64) error
}

type UserRepository interface {
	Repository[interfaces.User]
	GetById(id int64) (*interfaces.User, error)
}
