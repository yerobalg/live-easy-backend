package usecase

import (
	"live-easy-backend/src/repository"
)

type Usecase struct {
	User UserInterface
	Medicine MedicineInterface
}

func Init(repo *repository.Repository) *Usecase {
	return &Usecase{
		User: InitUser(repo.User),
		Medicine: InitMedicine(repo.Medicine),
	}
}
