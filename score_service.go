package recommendation_system_score_store

type ScoreService interface {
	CreateScore(cmd *CreateScoreCommand) (*Score, error)
	ListScore(cmd *ListScoreCommand) ([]Score, error)
}

type scoreService struct {
	store ScoreStore
}

func NewScoreService(s ScoreStore) ScoreService {
	return &scoreService{store: s}
}

func (s *scoreService) CreateScore(cmd *CreateScoreCommand) (*Score, error) {
	score := &Score{
		UserId:  cmd.UserId,
		MovieId: cmd.MovieId,
		Rating:  cmd.Rating,
	}
	return s.store.Create(score)
}

func (s *scoreService) ListScore(cmd *ListScoreCommand) ([]Score, error) {
	return s.store.List()
}
