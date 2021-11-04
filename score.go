package recommendation_system_score_store

type Score struct {
	Id      string  `json:"id"`
	UserId  string  `json:"user_id"`
	MovieId int64   `json:"movie_id"`
	Rating  float64 `json:"rating"`
}
