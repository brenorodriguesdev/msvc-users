package factories

import (
	"msvc-users/application/services"
	"msvc-users/infra"
	repositories "msvc-users/infra/repositories"
	"msvc-users/presentation/contracts"
	"msvc-users/presentation/controllers"
)

func MakeSignInController() contracts.Controller {
	repo := repositories.NewUserRepository()
	hasherCompare := infra.NewBcryptHasherCompare()
	encrypter := infra.NewJWTEncrypter()

	signInService := services.NewSignInService(repo, hasherCompare, encrypter)
	signInValidator := MakeSignInValidator()
	signInController := controllers.NewSignInController(signInService, signInValidator)

	return signInController
}
