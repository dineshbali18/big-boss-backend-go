package mysql

import (
	"big-boss-7/domain"
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) domain.UserRepository {
	return &repository{db: db}
}

func (repository *repository) UpdateMysqlUserRepository(db *gorm.DB) {
	repository.db = db
}

func (repository *repository) RegisterUserUsingDeviceId(userRegisterationPayload domain.UserRegisterationPayload) (userRegisterationResponse domain.UserRegisterationResponse, err error) {

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
