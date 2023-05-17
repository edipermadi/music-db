package theory

import "context"

type pitchService interface {
	ListPitches(ctx context.Context) ([]DetailedPitch, error)
	GetPitch(ctx context.Context, pitchID int64) (*DetailedPitch, error)
}

func (s theoryService) ListPitches(ctx context.Context) ([]DetailedPitch, error) {
	return s.repository.ListPitches(ctx)
}

func (s theoryService) GetPitch(ctx context.Context, pitchID int64) (*DetailedPitch, error) {
	return s.repository.GetPitch(ctx, pitchID)
}
