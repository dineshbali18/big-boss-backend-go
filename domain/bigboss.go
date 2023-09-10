package domain

// BBUsecase represents as a interface for BBUsecase
type BBUsecase interface {
	RegisterUserUsingDeviceID(userRegisterationPayload UserRegisterationPayload) (userRegisterationResponse UserRegisterationResponse, err error)
	VoteContestant(UserVotesPayload domain.UserVotesPayload)
	GetAllContestants() error
	GetNominatedContestants() error
	GetVotesInPercentages() error
	GetUserVotes(deviceID int) (votesLeft, error)
}

// BBRepository represents as a interface for BBRepository
type BBRepository interface {
	RegisterWithDeviceID(userRegisterationPayload UserRegisterationPayload) (userRegisterationResponse UserRegisterationResponse, err error)
	VoteContestant(contestantID int) error
	GetContestants() error
	GetContestants() error
	GetNominatedContestants() error
	GetAllContestantVotes()
	GetUserVotes(deviceID int)
}
