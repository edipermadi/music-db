package theory

import (
	"context"
	"github.com/edipermadi/music-db/internal/platform/api"
	"go.uber.org/zap"
)

// Service is music theory service
type Service interface {
	ListPitches(ctx context.Context) ([]DetailedPitch, error)

	ListChords(ctx context.Context, pagination api.Pagination) ([]DetailedChord, *api.Pagination, error)

	ListScales(ctx context.Context, pagination api.Pagination) ([]DetailedScale, *api.Pagination, error)
	ListScaleKeys(ctx context.Context, scaleID int64) ([]DetailedKey, error)
	GetScale(ctx context.Context, scaleID int64) (*DetailedScale, error)

	ListKeys(ctx context.Context, pagination api.Pagination) ([]DetailedKey, *api.Pagination, error)
	ListKeyModes(ctx context.Context, keyID int64) ([]DetailedKey, error)
	ListKeyChords(ctx context.Context, keyID int64, pagination api.Pagination) ([]DetailedChord, *api.Pagination, error)
	ListKeyPitches(ctx context.Context, keyID int64) ([]DetailedPitch, error)
	GetKey(ctx context.Context, keyID int64) (*DetailedKey, error)
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

func (s theoryService) ListPitches(ctx context.Context) ([]DetailedPitch, error) {
	return s.repository.ListPitches(ctx)
}

func (s theoryService) ListChords(ctx context.Context, pagination api.Pagination) ([]DetailedChord, *api.Pagination, error) {
	return s.repository.ListChords(ctx, pagination)
}
func (s theoryService) ListScales(ctx context.Context, pagination api.Pagination) ([]DetailedScale, *api.Pagination, error) {
	return s.repository.ListScales(ctx, pagination)
}

func (s theoryService) ListScaleKeys(ctx context.Context, scaleID int64) ([]DetailedKey, error) {
	return s.repository.ListScaleKeys(ctx, scaleID)
}

func (s theoryService) GetScale(ctx context.Context, scaleID int64) (*DetailedScale, error) {
	return s.repository.GetScale(ctx, scaleID)
}

func (s theoryService) ListKeys(ctx context.Context, pagination api.Pagination) ([]DetailedKey, *api.Pagination, error) {
	return s.repository.ListKeys(ctx, pagination)
}

func (s theoryService) ListKeyModes(ctx context.Context, keyID int64) ([]DetailedKey, error) {
	return s.repository.ListKeyModes(ctx, keyID)
}

func (s theoryService) ListKeyChords(ctx context.Context, keyID int64, pagination api.Pagination) ([]DetailedChord, *api.Pagination, error) {
	return s.repository.ListKeyChords(ctx, keyID, pagination)
}

func (s theoryService) ListKeyPitches(ctx context.Context, keyID int64) ([]DetailedPitch, error) {
	return s.repository.ListKeyPitches(ctx, keyID)
}

func (s theoryService) GetKey(ctx context.Context, keyID int64) (*DetailedKey, error) {
	return s.repository.GetKey(ctx, keyID)
}
