package api

import (
	"context"
	"errors"
	"time"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

// Question forms the basis of our data structure for
// appengine Datastore documents
type Question struct {
	Key          *datastore.Key `json:"id" datastore:"-"`
	CTime        time.Time      `json:"created"`
	Question     string         `json:"question"`
	User         UserCard       `json:"user"`
	AnswersCount int            `json:"answers_count"`
}

// OK ensures that before we store any questions
// we valiate that they're at least long enough
func (q Question) OK() error {
	if len(q.Question) < 10 {
		return errors.New("question is too short")
	}
	return nil
}

func (q *Question) Create(ctx context.Context) error {
	log.Debugf(ctx, "Saving question: %s", q.Question, nil)
	if q.Key == nil {
		q.Key = datastore.NewIncompleteKey(ctx, "Question", nil)
	}

	user, err := UserFromAEUser(ctx)
	if err != nil {
		return err
	}

	q.User = user.Card()
	q.CTime = time.Now()
	q.Key, err = datastore.Put(ctx, q.Key, q)
	if err != nil {
		return err
	}

	return nil
}

func (q *Question) Update(ctx context.Context) error {
	if q.Key == nil {
		q.Key = datastore.NewIncompleteKey(ctx, "Question", nil)
	}

	var err error
	q.Key, err = datastore.Put(ctx, q.Key, q)
	if err != nil {
		return err
	}

	return nil
}

func GetQuestion(ctx context.Context, key *datastore.Key) (*Question, error) {
	var q Question
	err := datastore.Get(ctx, key, &q)
	if err != nil {
		return nil, err
	}

	q.Key = key
	return &q, nil
}
