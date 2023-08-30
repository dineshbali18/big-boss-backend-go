package usecase

import "big-boss-7/domain"

type usecase struct {
	repository domain.BBRepository
}

func NewUser(repository domain.BBRepository) domain.BBUsecase {
	return &usecase{repository: repository}
}

func (useCase *usecase) RegisterUserUsingDeviceID(userRegisterationPayload domain.UserRegisterationPayload) (userRegisterationResponse domain.UserRegisterationResponse, err error) {
	// user registeration
	userRegisterationResponse, err = useCase.repository.RegisterWithDeviceID(userRegisterationPayload)
	if err != nil {
		userRegisterationResponse.Status = "config.Failure"
		userRegisterationResponse.Description = "Failed to register user"
		return
	}

	userRegisterationResponse.Status = "config.Success"
	userRegisterationResponse.Description = "User registered successfully"
	return
}
