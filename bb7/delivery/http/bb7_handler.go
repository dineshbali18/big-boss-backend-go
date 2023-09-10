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

	e.POST("/v1/bb/users/register/device", handler.RegisterUserUsingDeviceId)
	e.POST("/v1/bb/users/votes", handler.VoteContestant)
	e.GET("/v1/bb/contestants", handler.GetAllContestants)
	e.GET("/v1/bb/nominated/contestants", handler.GetNominatedContestants)
	e.GET("/v1/bb/contestants/results", handler.GetVotesInPercentages)
	e.GET("/v1/bb/user/:id/votes", handler.GetUserVotes)
}

func (delivery *delivery) RegisterUserUsingDeviceId(context echo.Context) error {
	context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	var userRegisterationPayload domain.UserRegisterationPayload
	err := json.NewDecoder(context.Request().Body).Decode(&userRegisterationPayload)
	if err != nil {
		log.Fatal(err.Error())
		return context.JSON(http.StatusBadRequest, domain.InvalidUserRegisterationPayload)
	}

	if len(userRegisterationPayload.ApiToken) > 0 && userRegisterationPayload.ApiToken != "dineshbali91210850445@" {
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

	// Return status in the response
	responseWithStatus := domain.UserRegisterationResponsewithJWT{
		Status: response.Status,
	}

	return context.JSON(http.StatusOK, responseWithStatus.Status)
}

func (delivery *delivery) VoteContestant(context echo.Context) error {
	context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	var userRegisterationPayload domain.UserVotesPayload
	err := json.NewDecoder(context.Request().Body).Decode(&userRegisterationPayload)
	if err != nil {
		log.Fatal(err.Error())
		return context.JSON(http.StatusBadRequest, domain.InvalidUserVotesPayload)
	}
	if len(userRegisterationPayload.ApiToken) > 0 && userRegisterationPayload.ApiToken != "dineshbali91210850445@" {
		return context.JSON(http.StatusBadRequest, domain.InvalidUserVotesPayload)
	}
	result, err := delivery.BBUsecase.VoteContestant(userRegisterationPayload)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}
	return context.JSON(http.StatusOK, result)
}

func (delivery *delivery) GetAllContestants(context echo.Context) error {
	response, err := delivery.BBUsecase.GetAllContestants()
	if err != nil {
		return context.JSON(http.StatusBadRequest, 500)
	}
	return context.JSON(http.StatusOK, response)
}

func (delivery *delivery) GetNominatedContestants(context echo.Context) error {
	response, err := delivery.BBUsecase.GetNominatedContestants()
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}
	return context.JSON(http.StatusOK, response)
}

func (delivery *delivery) GetVotesInPercentages(context echo.Context) error {
	response, err := delivery.BBUsecase.GetVotesInPercentages()
	if err != nil {
		log.Fatal("Get votes in percentages", err)
		return context.JSON(http.StatusBadRequest, err)
	}
	return context.JSON(http.StatusOK, response)
}

func (delivery *delivery) GetUserVotes(context echo.Context) error {
	deviceID := context.Param("id")
	if len(deviceID) == 0 {
		log.Fatal("Error in Fetching user votes")
		return context.JSON(http.StatusBadRequest, "Invalid Device ID")
	}
	votesAvailable, err := delivery.BBUsecase.GetUserVotes(deviceID)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}
	return context.JSON(http.StatusOK, votesAvailable)
}
