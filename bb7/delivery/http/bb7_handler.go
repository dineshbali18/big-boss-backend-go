package http

import (
	"big-boss-7/domain"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type delivery struct {
	BBUsecase domain.BBUsecase
}

func NewBBHandler(e *echo.Echo, useCase domain.BBUsecase) {
	handler := &delivery{BBUsecase: useCase}

	e.POST("/v1/bb7/users/register/device", handler.RegisterUserUsingDeviceId)
}

func (delivery *delivery) RegisterUserUsingDeviceId(context echo.Context) error {
	context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	var userRegisterationPayload domain.UserRegisterationPayload
	err := json.NewDecoder(context.Request().Body).Decode(&userRegisterationPayload)
	if err != nil {
		log.Fatal(err.Error())
		return context.JSON(http.StatusBadRequest, domain.InvalidUserRegisterationPayload)
	}

	if userRegisterationPayload.DeviceID == nil || (userRegisterationPayload.DeviceID != nil && len(*userRegisterationPayload.DeviceID) == 0) {
		log.Fatal(err.Error())
		return context.JSON(http.StatusBadRequest, domain.InvalidDeviceIDPayload)
	}

	response, err := delivery.BBUsecase.RegisterUserUsingDeviceID(userRegisterationPayload)
	if err != nil {
		log.Fatal(err.Error())
		return context.JSON(http.StatusInternalServerError, domain.ErrUnexpectedError)
	}

	// Generate token
	// JWTtoken, err := auth.GenerateJWT(*userRegisterationPayload.DeviceID)
	// if err != nil {
	// 	return context.JSON(http.StatusInternalServerError, domain.ErrUnexpectedError)
	// }

	JWTtoken := "dinesh"

	// Return token in the response
	responseWithToken := domain.UserRegisterationResponsewithJWT{
		JWTToken:    JWTtoken,
		Status:      response.Status,
		Description: response.Description,
		DeviceID:    response.DeviceID,
	}

	return context.JSON(http.StatusOK, responseWithToken)
}
