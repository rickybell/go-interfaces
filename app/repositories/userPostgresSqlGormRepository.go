package repositories

import (
	"context"
	"errors"

	"github.com/rickybell/go-interfaces/app/interfaces"
	"gorm.io/gorm"
)

type UserPostgresSqlGormRepository struct {
	PostgreSQLGORMRepositoryGenerics[interfaces.User]
	ctx context.Context
	u   interfaces.User
	db  *gorm.DB
}

func NewUserPostgresSqlGormRepository(db *gorm.DB, ctx context.Context) UserPostgresSqlGormRepository {
	return UserPostgresSqlGormRepository{
		ctx: ctx,
		db:  db,
		PostgreSQLGORMRepositoryGenerics: PostgreSQLGORMRepositoryGenerics[interfaces.User]{
			db: db,
		},
	}
}

func (u *UserPostgresSqlGormRepository) GetById(Id int64) (*interfaces.User, error) {
	respository := PostgreSQLGORMRepositoryGenerics[interfaces.User]{
		db: u.db,
	}

	err := respository.db.WithContext(u.ctx).Where("id = ?", Id).First(&u.u).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotExist
		}
		return nil, err
	}

	result := interfaces.User(u.u)
	return &result, nil
}

// func (u *UserPostgresSqlGormRepository) GetByName(name string) (*interfaces.User, error) {
// 	respository := PostgreSQLGORMRepositoryGenerics[entities.User]

// func (u *UserPostgresSqlGormRepository) All() ([]interfaces.User, error) {
// 	var users []interfaces.User
// 	err := u.db.WithContext(c).Find(&users).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	result := make([]interfaces.User, len(users))

// 	for i, user := range users {
// 		result[i] = user
// 	}

// 	return result, nil
// }

// func (u *UserPostgresSqlGormRepository) Create(c context.Context, user interfaces.User) (*interfaces.User, error) {
// 	err := u.db.WithContext(c).Create(&user).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	result := interfaces.User(user)

// 	return &result, nil
// }

// func (u *UserPostgresSqlGormRepository) Update(c context.Context, id int64, updated interfaces.User) (*interfaces.User, error) {
// 	err := u.db.WithContext(c).Where("id = ?", id).Save(&updated).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	result := interfaces.User(updated)

// 	return &result, nil
// }

// func (u *UserPostgresSqlGormRepository) Delete(c context.Context, id int64) error {
// 	err := u.db.WithContext(c).Where("id = ?", id).Delete(&entities.User{}).Error

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (u *UserPostgresSqlGormRepository) getContext() context.Context {
// 	return u.ctx
// }

// func (u *UserPostgresSqlGormRepository) Migrate(ctx context.Context) error {

// 	return nil
// }
