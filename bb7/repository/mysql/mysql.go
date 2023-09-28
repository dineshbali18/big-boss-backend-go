package mysql

import (
	"big-boss-7/domain"
	"context"
	"errors"
	"fmt"

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

func (repository *repository) DecrementUserVotes(tx *gorm.DB, deviceID string, votes int) error {
	err := tx.Exec("UPDATE users SET votes=votes-? where device_id=?", votes, deviceID).Error
	if err != nil {
		return err
	}
	return err
}

func (repository *repository) GetAllContestants() ([]domain.Contestants, error) {
	var nominatedContestants []domain.Contestants
	err := repository.db.WithContext(context.Background()).Table(domain.ContestantsTable).Scan(&nominatedContestants).Error
	if err != nil {
		return nominatedContestants, err
	}
	return nominatedContestants, err
}

func (repository *repository) GetNominatedContestants() ([]domain.Contestants, error) {
	var nominatedContestants []domain.Contestants
	err := repository.db.WithContext(context.Background()).Table(domain.ContestantsTable).Select("id", "name", "image").Where("is_nominated=1").Find(&nominatedContestants).Error
	if err != nil {
		return nominatedContestants, err
	}
	return nominatedContestants, err
}

// returns array of objects with contestant name and the number of votes
func (repository *repository) GetAllContestantsVotes() ([]domain.ContestantVotes, error) {
	var contestantVotes []domain.ContestantVotes
	// join contestant and contestant votes
	err := repository.db.WithContext(context.Background()).Table(domain.ContestantsVotesTable).
		Joins("INNER JOIN contestants ON contestants.id = contestants_votes.contestant_id").Where("contestants.is_nominated=1").Select("contestants.id", "contestants.name", "contestants_votes.votes").
		Find(&contestantVotes).Error
	if err != nil {
		return contestantVotes, err
	}
	return contestantVotes, err
}

func (repository *repository) GetUserVotes(deviceID string) (int, error) {
	var votes int
	result := repository.db.WithContext(context.Background()).Table(domain.UsersTable).Select("votes").Where("device_id = ?", deviceID).Scan(&votes)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Return a custom error when the record is not found.
			return -1, fmt.Errorf("deviceID '%s' not found", deviceID)
		}
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		// No rows were affected, indicating the record does not exist.
		return -1, fmt.Errorf("deviceID '%s' not found", deviceID)
	}

	return votes, nil
}
