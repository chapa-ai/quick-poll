package pg

import (
	"context"
	"encoding/json"
	"quick-poll/internal/models"
	"quick-poll/pkg/errors"
)

func (r *Repository) Create(ctx context.Context, question string, options []string) (*models.Poll, error) {
	r.logger.Info("creating poll")

	opts := make(map[string]int)
	for _, opt := range options {
		opts[opt] = 0
	}

	optionsJSON, err := json.Marshal(opts)
	if err != nil {
		r.logger.Errorf("failed to marshal options: %v", err)
		return nil, errors.Wrapf(err, "failed to encode options")
	}

	var id string
	err = r.db.QueryRow(ctx, `INSERT INTO polls (question, options) VALUES ($1, $2) RETURNING id`, question, optionsJSON).Scan(&id)
	if err != nil {
		r.logger.Errorf("failed creating poll: %v", err)
		return nil, errors.Wrapf(err, "failed to create poll")
	}

	poll := &models.Poll{
		ID:       id,
		Question: question,
	}

	return poll, nil
}

func (r *Repository) Vote(ctx context.Context, pollID, option string) error {
	r.logger.Infof("voting for pollID=%s, option=%s", pollID, option)
	_, err := r.db.Exec(ctx, `UPDATE polls SET options = jsonb_set(options, array[$2], 
	to_jsonb(COALESCE((options->>$2)::int, 0) + 1)) WHERE id = $1`, pollID, option)
	if err != nil {
		r.logger.Errorf("voting failed: %v", err)
		return errors.Wrapf(err, "failed to vote poll")
	}
	return nil
}

func (r *Repository) GetByID(ctx context.Context, pollID string) (*models.Poll, error) {
	var poll models.Poll
	var optionsJSON []byte

	err := r.db.QueryRow(ctx, `SELECT id, question, options FROM polls 
	 WHERE id = $1`, pollID).Scan(&poll.ID, &poll.Question, &optionsJSON)
	if err != nil {
		r.logger.Errorf("failed to get poll: %v", err)
		return nil, errors.Wrapf(err, "failed to get poll")
	}
	if err = json.Unmarshal(optionsJSON, &poll.Options); err != nil {
		r.logger.Errorf("unmarshalling failed: %v", err)
		return nil, errors.Wrapf(err, "failed unmarshalling")
	}

	return &poll, nil
}
