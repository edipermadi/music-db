package theory

import (
	"context"

	"github.com/edipermadi/music-db/internal/platform/api"
)

type chordService interface {
	GetChord(ctx context.Context, chordID int64) (*DetailedChord, error)
	GetChordQuality(ctx context.Context, chordID int64) (*DetailedChordQuality, error)
	ListChordKeys(ctx context.Context, chordID int64, filter KeyFilter, pagination api.Pagination) ([]SimplifiedKey, *api.Pagination, error)
	ListChordPitches(ctx context.Context, chordID int64) ([]SimplifiedPitch, error)
	ListChordScales(ctx context.Context, chordID int64, filter ScaleFilter, pagination api.Pagination) ([]SimplifiedScale, *api.Pagination, error)
	ListChords(ctx context.Context, filter ChordFilter, pagination api.Pagination) ([]SimplifiedChord, *api.Pagination, error)
}

func (s theoryService) ListChords(ctx context.Context, filter ChordFilter, pagination api.Pagination) ([]SimplifiedChord, *api.Pagination, error) {
	return s.repository.ListChords(ctx, filter, pagination)
}

func (s theoryService) ListChordPitches(ctx context.Context, chordID int64) ([]SimplifiedPitch, error) {
	return s.repository.ListChordPitches(ctx, chordID)
}

func (s theoryService) ListChordKeys(ctx context.Context, chordID int64, filter KeyFilter, pagination api.Pagination) ([]SimplifiedKey, *api.Pagination, error) {
	return s.repository.ListChordKeys(ctx, chordID, filter, pagination)
}

func (s theoryService) ListChordScales(ctx context.Context, chordID int64, filter ScaleFilter, pagination api.Pagination) ([]SimplifiedScale, *api.Pagination, error) {
	return s.repository.ListChordScales(ctx, chordID, filter, pagination)
}

func (s theoryService) GetChord(ctx context.Context, chordID int64) (*DetailedChord, error) {
	return s.repository.GetChord(ctx, chordID)
}

func (s theoryService) GetChordQuality(ctx context.Context, chordID int64) (*DetailedChordQuality, error) {
	return s.repository.GetChordQuality(ctx, chordID)
}
