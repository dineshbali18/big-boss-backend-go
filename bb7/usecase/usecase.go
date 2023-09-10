package usecase

import (
	"big-boss-7/domain"
	"fmt"
)

type usecase struct {
	repository domain.BBRepository
}

func NewUser(repository domain.BBRepository) domain.BBUsecase {
	return &usecase{repository: repository}
}

func (useCase *usecase) RegisterUserUsingDeviceID(userRegisterationPayload domain.UserRegisterationPayload) (userRegisterationResponse domain.UserRegisterationResponse, err error) {
	// user registeration
	userRegisterationResponse, err = useCase.repository.RegisterWithDeviceID(userRegisterationPayload)
	if err != nil {
		userRegisterationResponse.Status = "config.Failure"
		userRegisterationResponse.Description = "Failed to register user"
		return
	}

	userRegisterationResponse.Status = "config.Success"
	userRegisterationResponse.Description = "User registered successfully"
	return
}

func (usecase *usecase) VoteContestant(UserVotesPayload domain.UserVotesPayload) (int, error) {
	votesLeft, err := usecase.GetUserVotes(UserVotesPayload.DeviceID)
	if err != nil {
		return 0, err
	}
	if votesLeft <= 0 {
		//prepare a customized error
		return 0, err
	}

	if len(UserVotesPayload.UserVote) > 0 {
		for i := 0; i < len(UserVotesPayload.UserVote); i++ {
			usecase.repository.VoteContestant(UserVotesPayload.UserVote[i].ContestantID, 1)
		}
	}
	return 0, err
}

// cache it for 5 months.
func (usecase *usecase) GetAllContestants() ([]domain.Contestants, error) {
	response, err := usecase.repository.GetAllContestants()
	if err != nil {
		return response, err
	}
	return response, err
}

// cache them for 1 week from monday to sunday to monday
func (usecase *usecase) GetNominatedContestants() ([]domain.Contestants, error) {
	response, err := usecase.repository.GetNominatedContestants()
	if err != nil {
		return response, err
	}
	return response, err
}

// cache it for 15 min
func (usecase *usecase) GetVotesInPercentages() ([]domain.VotesPercentages, error) {
	var votesPercentages []domain.VotesPercentages
	// call getAllContestantVotes
	votes, err := usecase.repository.GetAllContestantsVotes()
	if err != nil {
		return votesPercentages, err
	}
	fmt.Println(votes)

	return votesPercentages, err
}

// to get the unused votes count left for the user
// if it's zero keep it in cache until that day ends.(don't make any additional requests)
func (usecase *usecase) GetUserVotes(deviceID string) (int, error) {
	votesLeft, err := usecase.repository.GetUserVotes(deviceID)
	if err != nil {
		return votesLeft, err
	}
	return votesLeft, err
}
