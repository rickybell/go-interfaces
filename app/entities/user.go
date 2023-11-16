package entities

import (
	"github.com/google/uuid"
	"github.com/rickybell/go-interfaces/app/interfaces"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var tablename string = "users"

type User struct {
	gorm.Model
	interfaces.User
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name     string    `gorm:"uniqueIndex; not null"`
	Email    string    `gorm:"not null"`
	Password string    `gorm:"not null"`
}

func (u *User) TableName() string {
	return tablename
}

func (u *User) IsValid() bool {
	return u.Name != "" && u.Email != ""
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password, err = HashPassword(u.Password)
	return
}

func (u *User) GetById(id uuid.UUID) (*interfaces.User, error) {
	return nil, nil
}
