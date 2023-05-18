package theory

import (
	"context"

	"github.com/edipermadi/music-db/internal/platform/api"
)

type scaleService interface {
	GetScale(ctx context.Context, scaleID int64) (*DetailedScale, error)
	ListScaleChords(ctx context.Context, scaleID int64, filter ChordFilter, pagination api.Pagination) ([]SimplifiedChord, *api.Pagination, error)
	ListScaleKeys(ctx context.Context, scaleID int64) ([]SimplifiedKey, error)
	ListScalePitches(ctx context.Context, scaleID int64) ([]SimplifiedPitch, error)
	ListScales(ctx context.Context, filter ScaleFilter, pagination api.Pagination) ([]SimplifiedScale, *api.Pagination, error)
}

func (s theoryService) ListScales(ctx context.Context, filter ScaleFilter, pagination api.Pagination) ([]SimplifiedScale, *api.Pagination, error) {
	return s.repository.ListScales(ctx, filter, pagination)
}

func (s theoryService) ListScaleKeys(ctx context.Context, scaleID int64) ([]SimplifiedKey, error) {
	return s.repository.ListScaleKeys(ctx, scaleID)
}

func (s theoryService) ListScalePitches(ctx context.Context, scaleID int64) ([]SimplifiedPitch, error) {
	return s.repository.ListScalePitches(ctx, scaleID)
}

func (s theoryService) ListScaleChords(ctx context.Context, scaleID int64, filter ChordFilter, pagination api.Pagination) ([]SimplifiedChord, *api.Pagination, error) {
	return s.repository.ListScaleChords(ctx, scaleID, filter, pagination)
}

func (s theoryService) GetScale(ctx context.Context, scaleID int64) (*DetailedScale, error) {
	return s.repository.GetScale(ctx, scaleID)
}
