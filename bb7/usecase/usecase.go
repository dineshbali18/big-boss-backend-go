package usecase

import (
	"big-boss-7/domain"
	"fmt"
	"sort"
)

type usecase struct {
	repository domain.BBRepository
	cache      domain.CacheService
}

func NewUser(repository domain.BBRepository, cache domain.CacheService) domain.BBUsecase {
	return &usecase{repository: repository, cache: cache}
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
			err1 := usecase.repository.DecrementUserVotes(tx, UserVotesPayload.DeviceID, 1)
			if err1 != nil {
				tx.Rollback()
				return 0, err1
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
	response, err := usecase.cache.GetAllContestants()
	if err == nil {
		return response, err
	}
	response, err = usecase.repository.GetAllContestants()
	if err != nil {
		return response, err
	}
	saveErr := usecase.cache.SaveAllContestants(response)
	if saveErr != nil {
		fmt.Println("Error while saving data into cache")
	}
	return response, err
}

// cache them for 1 week from monday to sunday to monday
func (usecase *usecase) GetNominatedContestants() ([]domain.Contestants, error) {
	response, err := usecase.cache.GetNominatedContestants()
	if err == nil {
		return response, err
	}
	response, err = usecase.repository.GetNominatedContestants()
	if err != nil {
		return response, err
	}
	saveErr := usecase.cache.SaveNominatedContestants(response)
	if saveErr != nil {
		fmt.Println("Error while saving nominated contestants")
	}
	return response, err
}

// cache it for 15 min
func (usecase *usecase) GetVotesInPercentages() (votesPercentages domain.VotesPercentages, err error) {
	response, cacheErr := usecase.cache.GetPercentagesResults()
	if cacheErr == nil {
		return response, cacheErr
	}
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

	for j := range votes {
		tmpVotes := votes[j].Votes
		tempPercentage := float32(tmpVotes) / float32(totalVotes)
		votesPercentages.Percentages = append(votesPercentages.Percentages, tempPercentage*100)
	}
	saveErr := usecase.cache.SavePercentagesResults(votesPercentages)
	if saveErr != nil {
		fmt.Println("Error in saving percentage results")
	}

	return votesPercentages, err
}

// to get the unused votes count left for the user
// if it's zero keep it in cache until that day ends.(don't make any additional requests)
func (usecase *usecase) GetUserVotes(deviceID string) (int, error) {
	votesLeft, err := usecase.repository.GetUserVotes(deviceID)
	if votesLeft == -1 {
		var newUser domain.UserRegisterationPayload
		newUser.ApiToken = "dineshbali91210850445@"
		newUser.DeviceID = &deviceID
		_, err := usecase.RegisterUserUsingDeviceID(newUser)
		if err != nil {
			votesLeft, err := usecase.repository.GetUserVotes(deviceID)
			if err != nil {
				return votesLeft, err
			}
		}
		return votesLeft, err
	}
	return votesLeft, err
}
