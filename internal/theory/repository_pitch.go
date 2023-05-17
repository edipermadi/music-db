package theory

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/edipermadi/music-db/internal/platform/api"
)

type pitchRepository interface {
	GetPitch(ctx context.Context, pitchID int64) (*DetailedPitch, error)
	ListPitchChords(ctx context.Context, pitchID int64, filter ChordFilter, pagination api.Pagination) ([]SimplifiedChord, *api.Pagination, error)
	ListPitches(ctx context.Context) ([]DetailedPitch, error)
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

func (r theoryRepository) ListPitchChords(ctx context.Context, pitchID int64, filter ChordFilter, pagination api.Pagination) ([]SimplifiedChord, *api.Pagination, error) {
	pagination.Sanitize()

	where := []string{"cp.pitch_id = ?"}
	args := []interface{}{pitchID}

	if filter.ChordQualityID > 0 {
		where = append(where, "c.chord_quality_id = ?")
		args = append(args, filter.ChordQualityID)
	}

	if filter.RootID > 0 {
		where = append(where, "c.root_id = ?")
		args = append(args, filter.RootID)
	}

	if filter.Number > 0 {
		where = append(where, "c.number = ?")
		args = append(args, filter.Number)
	}

	queryCount := fmt.Sprintf(`
		SELECT 
			COUNT(DISTINCT c.id)
		FROM chord_pitches cp
			JOIN chords c on cp.chord_id = c.id
		WHERE %s
		GROUP BY
		    cp.pitch_id;`, strings.Join(where, " AND "))

	var total int
	if err := r.db.GetContext(ctx, &total, r.db.Rebind(queryCount), args...); err != nil {
		return nil, nil, err
	}

	pagination.TotalItems = total
	pagination.TotalPages = int(math.Ceil(float64(total) / float64(pagination.PerPage)))

	entries := make([]SimplifiedChord, 0)
	if pagination.Offset() > total {
		return entries, nil, nil
	}

	pagination.NextPage = (pagination.Page + 1) % (pagination.TotalPages + 1)

	queryList := fmt.Sprintf(`
		SELECT DISTINCT
			c.id, 
			c.name
		FROM chord_pitches cp
			JOIN chords c on cp.chord_id = c.id
		WHERE
		    %s
		ORDER BY
		    c.id
		OFFSET ?
		LIMIT  ?;`, strings.Join(where, " AND "))

	args = append(args, pagination.Offset(), pagination.PerPage)
	if err := r.db.SelectContext(ctx, &entries, r.db.Rebind(queryList), args...); err != nil {
		return nil, nil, err
	}

	return entries, &pagination, nil
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
