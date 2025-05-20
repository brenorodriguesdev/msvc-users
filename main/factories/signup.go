package factories

import (
	"msvc-users/application/services"
	"msvc-users/infra"
	repositories "msvc-users/infra/repositories"
	"msvc-users/presentation/contracts"
	"msvc-users/presentation/controllers"
	"msvc-users/utils"
)

func MakeSignUpController() contracts.Controller {
	repo := repositories.NewUserRepository()
	hasher := infra.NewBcryptHasher()
	passGen := utils.NewRandomPasswordGenerator()
	emailSender := infra.NewGmailMailer("your-email@gmail.com", "your-app-password")

	signUpService := services.NewSignUpService(repo, hasher, passGen, emailSender)
	signUpValidator := MakeSignUpValidator()
	signUpController := controllers.NewSignUpController(signUpService, signUpValidator)

	return signUpController
}
