package usecase

import (
	"go-Expense/pkg/models"
	"go-Expense/pkg/repository/interfaces"
	services "go-Expense/pkg/usecase/interfaces"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (uu UserUsecase) DeleteSection(id string) error {
	if err := uu.userRepo.DeleteSection(id); err != nil {
		log.Println("deleting section failed", err)
		return err
	}
	log.Println("section deleted")
	return nil
}

func (uu UserUsecase) GetAllSection() ([]primitive.M, error) {
	alldatas, err := uu.userRepo.GetAllSection()
	if err != nil {
		return []primitive.M{}, err
	}
	log.Println(alldatas)
	return alldatas, nil
}
