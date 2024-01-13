package interfaces

import (
	"go-Expense/pkg/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepo interface {
	AddSection(userName string, buget float64, description string) (models.Section, error)
	DeleteSection(id string) error
	GetAllSection() ([]primitive.M,error)
}
