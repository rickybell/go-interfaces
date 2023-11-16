package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rickybell/go-interfaces/app/interfaces"
	"gorm.io/gorm"
)

type PostgreSQLGORMRepositoryGenerics[T interfaces.Model] struct {
	Repository[T]
	m  T
	db *gorm.DB
}

func NewPostgreSQLGORMRepositoryGenerics[T interfaces.Model](db *gorm.DB, repository Repository[T]) *PostgreSQLGORMRepositoryGenerics[T] {
	return &PostgreSQLGORMRepositoryGenerics[T]{
		db:         db,
		Repository: repository,
	}
}

func (r *PostgreSQLGORMRepositoryGenerics[T]) Migrate(ctx context.Context) error {
	return r.db.WithContext(ctx).AutoMigrate(r.m)
}

func (r *PostgreSQLGORMRepositoryGenerics[T]) ClearAll(ctx context.Context) error {
	err := r.db.Where("1 = 1").Delete(r.m).Error
	if err != nil {
		return fmt.Errorf("failed to clear table: %w", err)
	}
	return nil
}

func (r *PostgreSQLGORMRepositoryGenerics[T]) Create(ctx context.Context, m T) (*T, error) {
	err := r.db.WithContext(ctx).Create(&m).Error
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}
	result := T(m)

	return &result, nil
}

func (r *PostgreSQLGORMRepositoryGenerics[T]) All(ctx context.Context) ([]T, error) {
	var localModel []T
	err := r.db.WithContext(ctx).Find(&localModel).Error
	if err != nil {
		return nil, err
	}

	var result []T
	for _, gw := range localModel {
		result = append(result, T(gw))
	}

	return result, nil
}

func (r *PostgreSQLGORMRepositoryGenerics[T]) GetByName(ctx context.Context, name string) (*T, error) {
	var localModel T
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&localModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotExist
		}
		return nil, err
	}

	result := T(localModel)

	return &result, nil
}

func (r *PostgreSQLGORMRepositoryGenerics[T]) Update(ctx context.Context, id int64, updated T) (*T, error) {
	localModel := T(updated)
	updateRes := r.db.WithContext(ctx).Where("id = ?", id).Save(&localModel)
	if err := updateRes.Error; err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	rowsAffected := updateRes.RowsAffected
	if rowsAffected == 0 {
		return nil, ErrUpdateFailed
	}

	return &updated, nil
}

func (r *PostgreSQLGORMRepositoryGenerics[T]) Delete(ctx context.Context, id int64) error {
	deleteRes := r.db.WithContext(ctx).Delete(r.m, id)
	if err := deleteRes.Error; err != nil {
		return err
	}

	rowsAffected := deleteRes.RowsAffected
	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return nil
}
