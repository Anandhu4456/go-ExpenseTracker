package interfaces

import (
	"go-Expense/pkg/models"
)

type UserRepo interface {
	AddSection(userName string, buget float64, description string) (models.Section, error)
	DeleteSection(id string)error
}
