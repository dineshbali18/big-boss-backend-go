package domain

import "gorm.io/gorm"

// BBUsecase represents as a interface for BBUsecase
type BBUsecase interface {
	RegisterUserUsingDeviceID(userRegisterationPayload UserRegisterationPayload) (userRegisterationResponse UserRegisterationResponse, err error)
	VoteContestant(UserVotesPayload UserVotesPayload) (int, error)
	GetAllContestants() ([]Contestants, error)
	GetNominatedContestants() ([]Contestants, error)
	GetVotesInPercentages() ([]VotesPercentages, error)
	GetUserVotes(deviceID string) (int, error)
}

// BBRepository represents as a interface for BBRepository
type BBRepository interface {
	RegisterWithDeviceID(userRegisterationPayload UserRegisterationPayload) (userRegisterationResponse UserRegisterationResponse, err error)
	VoteContestant(tx *gorm.DB, contestantID int, votes int) error
	GetAllContestants() ([]Contestants, error)
	GetNominatedContestants() ([]Contestants, error)
	GetAllContestantsVotes() ([]ContestantVotes, error)
	GetUserVotes(deviceID string) (int, error)
	GetDB() *gorm.DB
}
