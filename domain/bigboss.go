package domain

// BBUsecase represents as a interface for BBUsecase
type BBUsecase interface {
	RegisterUserUsingDeviceID(userRegisterationPayload UserRegisterationPayload) (userRegisterationResponse UserRegisterationResponse, err error)
}

// BBRepository represents as a interface for BBRepository
type BBRepository interface {
	RegisterWithDeviceID(userRegisterationPayload UserRegisterationPayload) (userRegisterationResponse UserRegisterationResponse, err error)
}
