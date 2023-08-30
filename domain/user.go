package domain

// List of tables being used in users module
const (
	// UsersTable is the name of the table that stores the users
	UsersTable = "users"
)

// UserRegisterationPayload represents the payload for user registration
type UserRegisterationPayload struct {
	ID       uint    `json:"id"`
	DeviceID *string `json:"deviceId"`
}

// UserRegisterationResponse represents the response for user registration
type UserRegisterationResponse struct {
	Status      string  `json:"status"`
	Description string  `json:"description"`
	UserID      uint    `json:"userId"`
	DeviceID    *string `json:"deviceId"`
}

type UserRegisterationResponsewithJWT struct {
	JWTToken    string  `json:"jwtToken,omitempty"`
	Status      string  `json:"status"`
	Description string  `json:"description"`
	DeviceID    *string `json:"deviceId"`
}

// type UserUseCase interface {
// 	RegisterUserUsingDeviceId(userRegisterationPayload UserRegisterationPayload) (userRegisterationResponse UserRegisterationResponse, err error)
// }

// type UserRepository interface {
// 	RegisterUserUsingDeviceId(userRegisterationPayload UserRegisterationPayload) (userRegisterationResponse UserRegisterationResponse, err error)
// }
