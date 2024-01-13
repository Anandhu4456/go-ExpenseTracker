package usecase

import (
	"go-Expense/pkg/models"
	"go-Expense/pkg/repository/interfaces"
	services "go-Expense/pkg/usecase/interfaces"
)

type UserUsecase struct {
	userRepo interfaces.UserRepo
}

func NewUserUsecase(userRepo interfaces.UserRepo) services.UserUsecase {
	return UserUsecase{
		userRepo: userRepo,
	}
}

func (uu UserUsecase) AddSection(userName string, buget float64, description string) (models.Section, error) {
	insertedSection, err := uu.userRepo.AddSection(userName, buget, description)
	if err != nil {
		return models.Section{}, err
	}
	return insertedSection, nil
}
