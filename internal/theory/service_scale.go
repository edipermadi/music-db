package theory

import (
	"context"

	"github.com/edipermadi/music-db/internal/platform/api"
)

type scaleService interface {
	ListScales(ctx context.Context, pagination api.Pagination) ([]DetailedScale, *api.Pagination, error)
	ListScaleKeys(ctx context.Context, scaleID int64) ([]DetailedKey, error)
	GetScale(ctx context.Context, scaleID int64) (*DetailedScale, error)
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
