package theory

import (
	"context"

	"github.com/edipermadi/music-db/internal/platform/api"
	"go.uber.org/zap"
)

// Service is music theory service
type Service interface {
	ListPitches(ctx context.Context) ([]DetailedPitch, error)

	ListChords(ctx context.Context, filter ChordFilter, pagination api.Pagination) ([]DetailedChord, *api.Pagination, error)
	ListChordPitches(ctx context.Context, chordID int64) ([]DetailedPitch, error)
	ListChordKeys(ctx context.Context, chordID int64, filter KeyFilter, pagination api.Pagination) ([]DetailedKey, *api.Pagination, error)
	GetChord(ctx context.Context, chordID int64) (*DetailedChord, error)
	GetChordQuality(ctx context.Context, chordID int64) (*DetailedChordQuality, error)

	ListScales(ctx context.Context, pagination api.Pagination) ([]DetailedScale, *api.Pagination, error)
	ListScaleKeys(ctx context.Context, scaleID int64) ([]DetailedKey, error)
	GetScale(ctx context.Context, scaleID int64) (*DetailedScale, error)

	ListKeys(ctx context.Context, filter KeyFilter, pagination api.Pagination) ([]DetailedKey, *api.Pagination, error)
	ListKeyModes(ctx context.Context, keyID int64, filter KeyFilter) ([]DetailedKey, error)
	ListKeyChords(ctx context.Context, keyID int64, filter ChordFilter, pagination api.Pagination) ([]DetailedChord, *api.Pagination, error)
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

func (s theoryService) ListChords(ctx context.Context, filter ChordFilter, pagination api.Pagination) ([]DetailedChord, *api.Pagination, error) {
	return s.repository.ListChords(ctx, filter, pagination)
}

func (s theoryService) ListChordPitches(ctx context.Context, chordID int64) ([]DetailedPitch, error) {
	return s.repository.ListChordPitches(ctx, chordID)
}

func (s theoryService) ListChordKeys(ctx context.Context, chordID int64, filter KeyFilter, pagination api.Pagination) ([]DetailedKey, *api.Pagination, error) {
	return s.repository.ListChordKeys(ctx, chordID, filter, pagination)
}

func (s theoryService) GetChord(ctx context.Context, chordID int64) (*DetailedChord, error) {
	return s.repository.GetChord(ctx, chordID)
}

func (s theoryService) GetChordQuality(ctx context.Context, chordID int64) (*DetailedChordQuality, error) {
	return s.repository.GetChordQuality(ctx, chordID)
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

func (s theoryService) ListKeys(ctx context.Context, filter KeyFilter, pagination api.Pagination) ([]DetailedKey, *api.Pagination, error) {
	return s.repository.ListKeys(ctx, filter, pagination)
}

func (s theoryService) ListKeyModes(ctx context.Context, keyID int64, filter KeyFilter) ([]DetailedKey, error) {
	return s.repository.ListKeyModes(ctx, keyID, filter)
}

func (s theoryService) ListKeyChords(ctx context.Context, keyID int64, filter ChordFilter, pagination api.Pagination) ([]DetailedChord, *api.Pagination, error) {
	return s.repository.ListKeyChords(ctx, keyID, filter, pagination)
}

func (s theoryService) ListKeyPitches(ctx context.Context, keyID int64) ([]DetailedPitch, error) {
	return s.repository.ListKeyPitches(ctx, keyID)
}

func (s theoryService) GetKey(ctx context.Context, keyID int64) (*DetailedKey, error) {
	return s.repository.GetKey(ctx, keyID)
}
