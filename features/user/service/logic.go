package service

import (
	"errors"
	"project/e-commerce/features/user"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userData user.DataInterface
}

func New(data user.DataInterface) user.UsecaseInterface {
	return &userUsecase{
		userData: data,
	}
}

func (usecase *userUsecase) PostData(data user.Core) (int, error) {
	if data.Name == "" {
		return -1, errors.New("masukan nama ")
	}
	if data.Email == "" {
		return -1, errors.New("masukan email ")
	}
	if data.Password == "" {
		return -1, errors.New("masukann password")
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	data.Password = string(hashPass)
	row, err := usecase.userData.InsertData(data)
	if err != nil {
		return -1, err
	}

	return row, nil
}
