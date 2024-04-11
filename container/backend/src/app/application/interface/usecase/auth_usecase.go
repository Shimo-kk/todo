package usecase

import "todo/app/application/schema"

type IAuthUsecase interface {
	SignUp(data schema.SignUpModel) error
	SignIn(data schema.SignInModel) (int, error)
}
