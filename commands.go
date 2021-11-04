package recommendation_system_score_store

type CreateScoreCommand struct {
	UserId  string  `json:"user_id"`
	MovieId int64   `json:"movie_id"`
	Rating  float64 `json:"rating"`
}

func (cmd *CreateScoreCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(ScoreService).CreateScore(cmd)
}

type ListScoreCommand struct {
}

func (cmd *ListScoreCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(ScoreService).ListScore(cmd)
}
