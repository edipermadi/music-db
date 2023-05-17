package theory

import (
	"context"
	"database/sql"
	"errors"
)

type pitchRepository interface {
	ListPitches(ctx context.Context) ([]DetailedPitch, error)
	GetPitch(ctx context.Context, pitchID int64) (*DetailedPitch, error)
}

func (r theoryRepository) ListPitches(ctx context.Context) ([]DetailedPitch, error) {
	query := `
		SELECT
			id,
			name,
			number,
			frequency
		FROM pitches
		ORDER BY id;`

	entries := make([]DetailedPitch, 0)
	if err := r.db.SelectContext(ctx, &entries, query); err != nil {
		return nil, err
	}

	return entries, nil
}

func (r theoryRepository) GetPitch(ctx context.Context, pitchID int64) (*DetailedPitch, error) {
	query := `
		SELECT
			id,
			name,
			number,
			frequency
		FROM pitches
		WHERE id = $1;`

	var entry DetailedPitch
	if err := r.db.GetContext(ctx, &entry, query, pitchID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPitchNotFound
		}
		return nil, err
	}

	return &entry, nil
}
