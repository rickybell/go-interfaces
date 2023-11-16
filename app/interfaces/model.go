package interfaces

// import "gorm.io/gorm"

type Model interface {
	// gorm.Model
	TableName() string
}
