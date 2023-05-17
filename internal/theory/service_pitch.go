package theory

import (
	"context"

	"github.com/edipermadi/music-db/internal/platform/api"
)

type pitchService interface {
	GetPitch(ctx context.Context, pitchID int64) (*DetailedPitch, error)
	ListPitchChords(ctx context.Context, pitchID int64, filter ChordFilter, pagination api.Pagination) ([]SimplifiedChord, *api.Pagination, error)
	ListPitches(ctx context.Context) ([]DetailedPitch, error)
}

func (s theoryService) GetPitch(ctx context.Context, pitchID int64) (*DetailedPitch, error) {
	return s.repository.GetPitch(ctx, pitchID)
}

func (s theoryService) ListPitchChords(ctx context.Context, pitchID int64, filter ChordFilter, pagination api.Pagination) ([]SimplifiedChord, *api.Pagination, error) {
	return s.repository.ListPitchChords(ctx, pitchID, filter, pagination)
}

func (s theoryService) ListPitches(ctx context.Context) ([]DetailedPitch, error) {
	return s.repository.ListPitches(ctx)
}
