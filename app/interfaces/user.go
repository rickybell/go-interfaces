package interfaces

type User interface {
	Model
	IsValid() bool
	GetById(Id int64) (*User, error)
}
