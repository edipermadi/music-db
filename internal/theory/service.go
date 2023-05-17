package theory

import (
	"go.uber.org/zap"
)

// Service is music theory service
type Service interface {
	chordService
	keyService
	pitchService
	scaleService
}

// NewService returns an implementation of music theory service
func NewService(logger *zap.Logger, repository Repository) Service {
	return theoryService{
		logger:     logger,
		repository: repository,
	}
}

type theoryService struct {
	logger     *zap.Logger
	repository Repository
}
