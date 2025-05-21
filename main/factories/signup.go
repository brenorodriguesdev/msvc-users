package factories

import (
	"msvc-users/application/services"
	"msvc-users/infra"
	repositories "msvc-users/infra/repositories"
	"msvc-users/presentation/contracts"
	"msvc-users/presentation/controllers"
	"msvc-users/utils"
)

// SignUpController.Handle
// @Summary Cadastra um novo usuário
// @Description Cria um usuário novo com email, senha, etc.
// @Tags users
// @Accept json
// @Produce json
// @Router /signup [post]
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
