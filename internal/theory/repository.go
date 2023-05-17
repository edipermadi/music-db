package theory

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Repository is music theory data repository
type Repository interface {
	chordRepository
	keyRepository
	pitchRepository
	scaleRepository
}

// NewRepository returns an implementation of music theory data repository
func NewRepository(logger *zap.Logger, db *sqlx.DB) Repository {
	return theoryRepository{
		logger: logger,
		db:     db,
	}
}

type theoryRepository struct {
	logger *zap.Logger
	db     *sqlx.DB
}
