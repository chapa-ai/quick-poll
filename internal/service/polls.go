package service

import (
	"context"
	"github.com/goccy/go-json"
	"quick-poll/internal/models"
	"quick-poll/pkg/errors"
)

func (s *Service) CreatePoll(ctx context.Context, question string, options []string) (*models.Poll, error) {
	if len(options) < 2 {
		s.Logger.Error("at least two options required")
		return nil, errors.New("at least two options required")
	}

	return s.DB.Create(ctx, question, options)
}

func (s *Service) Vote(ctx context.Context, pollID, option string) error {
	poll, err := s.DB.GetByID(ctx, pollID)
	if err != nil {
		s.Logger.Error("poll not found")
		return errors.ErrPollNotFound
	}

	if _, exists := poll.Options[option]; !exists {
		s.Logger.Error("invalid option")
		return errors.ErrInvalidOption
	}

	message := models.Poll{ID: pollID, Options: map[string]int{option: 1}}

	value, err := json.Marshal(message)
	if err != nil {
		s.Logger.Errorf("encoding message failed: %v", err)
		return errors.Wrapf(err, "failed to encode message")
	}

	if err = s.Broker.Publish(ctx, []byte(pollID), value); err != nil {
		s.Logger.Errorf("publishing vote failed: %v", err)
		return errors.Wrapf(err, "failed to publish vote")
	}

	return nil
}

func (s *Service) GetResults(ctx context.Context, pollID string) (*models.Poll, error) {
	poll, err := s.DB.GetByID(ctx, pollID)
	if err != nil {
		s.Logger.Error("poll not found")
		return nil, errors.ErrPollNotFound
	}
	return poll, nil
}
