package interfaces

import (
	"go-Expense/pkg/models"
	"time"
)

type UserUsecase interface {
	AddSection(userName string, buget float64, description string, createdAt time.Time) (models.Section, error)
}
