package theory

import (
	"context"

	"github.com/edipermadi/music-db/internal/platform/api"
)

type keyService interface {
	ListKeys(ctx context.Context, filter KeyFilter, pagination api.Pagination) ([]DetailedKey, *api.Pagination, error)
	GetKey(ctx context.Context, keyID int64) (*DetailedKey, error)
	ListKeyModes(ctx context.Context, keyID int64, filter KeyFilter) ([]DetailedKey, error)
	ListKeyChords(ctx context.Context, keyID int64, filter ChordFilter, pagination api.Pagination) ([]DetailedChord, *api.Pagination, error)
	ListKeyPitches(ctx context.Context, keyID int64) ([]DetailedPitch, error)
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
