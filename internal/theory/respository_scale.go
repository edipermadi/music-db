package theory

import (
	"context"
	"database/sql"
	"errors"
	"math"

	"github.com/edipermadi/music-db/internal/platform/api"
)

type scaleRepository interface {
	ListScales(ctx context.Context, pagination api.Pagination) ([]DetailedScale, *api.Pagination, error)
	ListScaleKeys(ctx context.Context, scaleID int64) ([]DetailedKey, error)
	GetScale(ctx context.Context, scaleID int64) (*DetailedScale, error)
}

func (r theoryRepository) ListScales(ctx context.Context, pagination api.Pagination) ([]DetailedScale, *api.Pagination, error) {
	pagination.Sanitize()

	queryCount := `
		SELECT
			COUNT(1)
		FROM scales;`

	var total int
	if err := r.db.GetContext(ctx, &total, queryCount); err != nil {
		return nil, nil, err
	}

	pagination.TotalItems = total
	pagination.TotalPages = int(math.Ceil(float64(total) / float64(pagination.PerPage)))

	entries := make([]DetailedScale, 0)
	if pagination.Offset() > total {
		return entries, nil, nil
	}

	pagination.NextPage = (pagination.Page + 1) % (pagination.TotalPages + 1)

	queryList := `
		SELECT
			id,
			name,
			cardinality,
			number,
			perfection,
			imperfection,
			pitch_class,
			interval_pattern,
			rotational_symmetric,
			rotational_symmetry_level,
			palindromic,
			reflectional_symmetric,
			reflectional_symmetry_axes,
			balanced
		FROM scales
		ORDER BY
			id
		OFFSET $1
		LIMIT  $2;`

	if err := r.db.SelectContext(ctx, &entries, queryList, pagination.Offset(), pagination.PerPage); err != nil {
		return nil, nil, err
	}

	return entries, &pagination, nil
}

func (r theoryRepository) ListScaleKeys(ctx context.Context, scaleID int64) ([]DetailedKey, error) {
	query := `
		SELECT
			k.id,
			s.id   AS "scale.id",
			s.name AS "scale.name",
			p.id   AS "tonic.id",
			p.name AS "tonic.name",
			k.name,
			k.number,
			k.balanced,
			k.center_x,
			k.center_y
		FROM keys k
			JOIN scales s ON k.scale_id = s.id
			JOIN pitches p ON k.tonic_id = p.id
		WHERE
			scale_id = $1
		ORDER BY
			k.id;`

	entries := make([]DetailedKey, 0)
	if err := r.db.SelectContext(ctx, &entries, query, scaleID); err != nil {
		return nil, err
	}

	return entries, nil
}

func (r theoryRepository) GetScale(ctx context.Context, scaleID int64) (*DetailedScale, error) {
	query := `
		SELECT
			id,
			name,
			cardinality,
			number,
			perfection,
			imperfection,
			pitch_class,
			interval_pattern,
			rotational_symmetric,
			rotational_symmetry_level,
			palindromic,
			reflectional_symmetric,
			reflectional_symmetry_axes,
			balanced
		FROM scales
		WHERE
			id = $1;`

	var scale DetailedScale
	if err := r.db.GetContext(ctx, &scale, query, scaleID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrScaleNotFound
		}
		return nil, err
	}

	return &scale, nil
}
