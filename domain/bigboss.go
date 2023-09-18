package domain

import "gorm.io/gorm"

// BBUsecase represents as a interface for BBUsecase
type BBUsecase interface {
	RegisterUserUsingDeviceID(userRegisterationPayload UserRegisterationPayload) (userRegisterationResponse UserRegisterationResponse, err error)
	VoteContestant(UserVotesPayload UserVotesPayload) (int, error)
	GetAllContestants() ([]Contestants, error)
	GetNominatedContestants() ([]Contestants, error)
	GetVotesInPercentages() (VotesPercentages, error)
	GetUserVotes(deviceID string) (int, error)
}

// BBRepository represents as a interface for BBRepository
type BBRepository interface {
	RegisterWithDeviceID(userRegisterationPayload UserRegisterationPayload) (userRegisterationResponse UserRegisterationResponse, err error)
	VoteContestant(tx *gorm.DB, contestantID int, votes int) error
	DecrementUserVotes(tx *gorm.DB, deviceID string, votes int) error
	GetAllContestants() ([]Contestants, error)
	GetNominatedContestants() ([]Contestants, error)
	GetAllContestantsVotes() ([]ContestantVotes, error)
	GetUserVotes(deviceID string) (int, error)
	GetDB() *gorm.DB
}

// CacheService
type CacheService interface {
	CheckRedisConnection() (result string, err error)
	GetAllContestants() (contestants []Contestants, err error)
	GetNominatedContestants() (contestants []Contestants, err error)
	GetPercentagesResults() (voteData VotesPercentages, err error)

	SaveAllContestants([]Contestants) (err error)
	SaveNominatedContestants([]Contestants) (err error)
	SavePercentagesResults(VotesPercentages) (err error)
}
