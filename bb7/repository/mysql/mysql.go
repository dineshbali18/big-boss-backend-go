package mysql

import (
	"big-boss-7/domain"
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) domain.BBRepository {
	return &repository{db: db}
}

func (repository *repository) UpdateMysqlUserRepository(db *gorm.DB) {
	repository.db = db
}

func (repository *repository) RegisterWithDeviceID(userRegisterationPayload domain.UserRegisterationPayload) (userRegisterationResponse domain.UserRegisterationResponse, err error) {

	type userData struct {
		ID       uint
		DeviceID *string
	}

	var user userData
	err = repository.db.WithContext(context.Background()).
		Table(domain.UsersTable).
		Where("device_id=?", userRegisterationPayload.DeviceID).
		Find(&user).Error

	if user.ID == 0 {
		var payload userData
		payload.DeviceID = userRegisterationPayload.DeviceID

		err = repository.db.WithContext(context.Background()).
			Table(domain.UsersTable).
			Create(&payload).Error

		userRegisterationResponse.UserID = payload.ID
		userRegisterationPayload.DeviceID = payload.DeviceID

		return
	}

	userRegisterationResponse.UserID = user.ID
	userRegisterationResponse.DeviceID = user.DeviceID

	return

}

func (repository *repository) VoteContestant(contestantID int) error {

}

func (repository *repository) GetContestants() error {

}

func (repository *repository) GetNominatedContestants() error {

}

func (repository *repository) GetAllContestantVotes() error {

}

func (repository *repository) GetUserVotes(deviceID int) (int, error) {
	var votes int
	err := repository.db.WithContext().Where(deviceID).Scan(votes).Error
	if err != nil {
		return -1, err
	}
	return votes, err
}
