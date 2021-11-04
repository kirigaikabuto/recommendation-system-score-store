package recommendation_system_score_store

type ScoreStore interface {
	Create(score *Score) (*Score, error)
	List() ([]Score, error)
}
