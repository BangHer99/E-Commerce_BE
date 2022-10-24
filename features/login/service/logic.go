package service

import (
	"project/e-commerce/middlewares"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"

	"project/e-commerce/features/login"
)

type authUsecase struct {
	authData login.DataInterface
}

func New(data login.DataInterface) login.UsecaseInterface {
	return &authUsecase{
		authData: data,
	}
}

func (usecase *authUsecase) Login(input login.Core) (login.Core, string, error) {
	res, err := usecase.authData.Login(input)
	if err != nil {
		log.Error(err.Error(), "username not found")
		return login.Core{}, "", err
	}

	pass := login.Core{Password: res.Password}
	check := bcrypt.CompareHashAndPassword([]byte(pass.Password), []byte(input.Password))
	if check != nil {
		log.Error(check, " wrong password")
		return login.Core{}, "", check
	}
	token, err := middlewares.CreateToken(int(res.ID))

	return res, token, err
}
