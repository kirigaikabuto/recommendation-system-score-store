package recommendation_system_score_store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var Queries = []string{
	`CREATE TABLE IF NOT EXISTS score (
		id serial primary key,
		user_id text,
		movie_id integer,
		rating float
	);`,
}

type postgreStore struct {
	db *sql.DB
}

func NewPostgreStore(cfg Config) (ScoreStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range Queries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	return &postgreStore{db: db}, err
}

func (p *postgreStore) List() ([]Score, error) {
	var scores []Score
	data, err := p.db.Query("select id, user_id, movie_id, rating from score")
	if err != nil {
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		score := Score{}
		err = data.Scan(&score.Id, &score.UserId, &score.MovieId, &score.Rating)
		if err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}
	return scores, nil
}

func (p *postgreStore) Create(score *Score) (*Score, error) {
	err := p.db.QueryRow("insert into score (user_id, movie_id, rating) values ($1,$2,$3) RETURNING id", score.UserId, score.MovieId, score.Rating).Scan(&score.Id)
	if err != nil {
		return nil, err
	}
	return score, nil
}
