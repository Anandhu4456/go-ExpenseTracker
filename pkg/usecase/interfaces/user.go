package interfaces

import (
	"go-Expense/pkg/models"
)

type UserUsecase interface {
	AddSection(userName string, buget float64, description string) (models.Section, error)
}
