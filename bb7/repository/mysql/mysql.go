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

func (repository *repository) GetDB() *gorm.DB {
	return repository.db
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

func (repository *repository) VoteContestant(tx *gorm.DB, contestantID int, votes int) error {
	err := tx.Exec("UPDATE contestants_votes SET votes=votes+? where contestant_id=?", votes, contestantID).Error
	if err != nil {
		return err
	}
	return err
}

func (repository *repository) GetAllContestants() ([]domain.Contestants, error) {
	var nominatedContestants []domain.Contestants
	err := repository.db.WithContext(context.Background()).Table(domain.ContestantsTable).Scan(nominatedContestants).Error
	if err != nil {
		return nominatedContestants, err
	}
	return nominatedContestants, err
}

func (repository *repository) GetNominatedContestants() ([]domain.Contestants, error) {
	var nominatedContestants []domain.Contestants
	err := repository.db.WithContext(context.Background()).Table(domain.ContestantsTable).Scan(nominatedContestants).Where("is_nominated=1").Error
	if err != nil {
		return nominatedContestants, err
	}
	return nominatedContestants, err
}

// returns array of objects with contestant name and the number of votes
func (repository *repository) GetAllContestantsVotes() ([]domain.ContestantVotes, error) {
	var contestantVotes []domain.ContestantVotes
	err := repository.db.WithContext(context.Background()).Table(domain.ContestantsTable).Scan(contestantVotes).Error
	if err != nil {
		return contestantVotes, err
	}
	return contestantVotes, err
}

func (repository *repository) GetUserVotes(deviceID string) (int, error) {
	var votes int
	err := repository.db.WithContext(context.Background()).Table(domain.UsersTable).Select("votes").Where("device_id=?", deviceID).Scan(votes).Error
	if err != nil {
		return -1, err
	}
	return votes, err
}
