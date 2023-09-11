package usecase

import (
	"big-boss-7/domain"
	"fmt"
	"sort"
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
	// Begin a database transaction
	tx := usecase.repository.GetDB().Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	votesLeft, err := usecase.GetUserVotes(UserVotesPayload.DeviceID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if votesLeft <= 0 {
		//prepare a customized error
		tx.Rollback()
		return 0, err
	}

	if len(UserVotesPayload.UserVote) > 0 {
		for i := 0; i < len(UserVotesPayload.UserVote); i++ {
			// change here if there is any change in number of votes
			// 1 in the below line indicates that one vote will be incremented in contestants_votes table
			err := usecase.repository.VoteContestant(tx, UserVotesPayload.UserVote[i].ContestantID, 1)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return 0, nil
}

// cache it for 5 months.
func (usecase *usecase) GetAllContestants() ([]domain.Contestants, error) {
	fmt.Println("IN USECASE")
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
func (usecase *usecase) GetVotesInPercentages() (domain.VotesPercentages, error) {
	var votesPercentages domain.VotesPercentages
	// call getAllContestantVotes
	votes, err := usecase.repository.GetAllContestantsVotes()
	if err != nil {
		return votesPercentages, err
	}
	fmt.Println(votes)
	sort.Slice(votes, func(i int, j int) bool {
		return votes[i].Votes > votes[j].Votes
	})
	fmt.Println("AFTER SORTING :", votes)
	var totalVotes int64
	for i := range votes {
		votesPercentages.Name = append(votesPercentages.Name, votes[i].Name)
		totalVotes += int64(votes[i].Votes)
	}

	for votes := range votesPercentages.Percentages {
		tempPercentage := float32(votes) / float32(totalVotes)
		votesPercentages.Percentages = append(votesPercentages.Percentages, tempPercentage)
	}

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
