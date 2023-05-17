package theory

import (
	"context"

	"github.com/edipermadi/music-db/internal/platform/api"
)

type chordService interface {
	ListChords(ctx context.Context, filter ChordFilter, pagination api.Pagination) ([]DetailedChord, *api.Pagination, error)
	ListChordPitches(ctx context.Context, chordID int64) ([]DetailedPitch, error)
	ListChordKeys(ctx context.Context, chordID int64, filter KeyFilter, pagination api.Pagination) ([]DetailedKey, *api.Pagination, error)
	GetChord(ctx context.Context, chordID int64) (*DetailedChord, error)
	GetChordQuality(ctx context.Context, chordID int64) (*DetailedChordQuality, error)
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
