package domain

const (
	ContestantsTable      = "contestants"
	ContestantsVotesTable = "contestants_votes"
)

type ContestantVotes struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Votes int    `json:"votes"`
}

type Contestants struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
}

type VotesPercentages struct {
	Name        []string  `json:"name"`
	Percentages []float32 `json:"percentages"`
}
