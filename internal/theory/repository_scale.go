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

type scaleRepository interface {
	GetScale(ctx context.Context, scaleID int64) (*DetailedScale, error)
	ListScaleChords(ctx context.Context, scaleID int64, filter ChordFilter, pagination api.Pagination) ([]SimplifiedChord, *api.Pagination, error)
	ListScaleKeys(ctx context.Context, scaleID int64) ([]SimplifiedKey, error)
	ListScalePitches(ctx context.Context, scaleID int64) ([]SimplifiedPitch, error)
	ListScales(ctx context.Context, filter ScaleFilter, pagination api.Pagination) ([]SimplifiedScale, *api.Pagination, error)
}

func (r theoryRepository) ListScales(ctx context.Context, filter ScaleFilter, pagination api.Pagination) ([]SimplifiedScale, *api.Pagination, error) {
	pagination.Sanitize()

	clauses := make([]string, 0)
	args := make([]interface{}, 0)

	if filter.TonicID > 0 {
		args = append(args, filter.TonicID)
		clauses = append(clauses, "k.tonic_id = ?")
	}

	if filter.Number > 0 {
		args = append(args, filter.Number)
		clauses = append(clauses, "k.number = ?")
	}

	if filter.Perfection != nil && (*filter.Perfection) >= 0 {
		args = append(args, *filter.Perfection)
		clauses = append(clauses, "s.perfection = ?")
	}

	if filter.Imperfection != nil && (*filter.Imperfection) >= 0 {
		args = append(args, *filter.Imperfection)
		clauses = append(clauses, "s.imperfection = ?")
	}

	if filter.Balanced != nil {
		args = append(args, *filter.Balanced)
		clauses = append(clauses, "k.balanced = ?")
	}

	if filter.RotationalSymmetric != nil {
		args = append(args, *filter.RotationalSymmetric)
		clauses = append(clauses, "s.rotational_symmetric = ?")
	}

	if filter.RotationalSymmetryLevel > 0 {
		args = append(args, filter.RotationalSymmetryLevel)
		clauses = append(clauses, "s.rotational_symmetry_level = ?")
	}

	if filter.ReflectionalSymmetric != nil {
		args = append(args, *filter.ReflectionalSymmetric)
		clauses = append(clauses, "s.reflectional_symmetric = ?")
	}

	if filter.Palindromic != nil {
		args = append(args, *filter.Palindromic)
		clauses = append(clauses, "s.palindromic = ?")
	}

	if filter.Cardinality > 0 {
		args = append(args, filter.Cardinality)
		clauses = append(clauses, "s.cardinality = ?")
	}

	condition := "TRUE"
	if len(clauses) > 0 {
		condition = strings.Join(clauses, " AND ")
	}

	queryCount := fmt.Sprintf(`
		SELECT
			COUNT(1)
		FROM scales s
		JOIN keys k ON s.id = k.scale_id
		WHERE
			%s;`, condition)

	var total int
	if err := r.db.GetContext(ctx, &total, r.db.Rebind(queryCount), args...); err != nil {
		return nil, nil, err
	}

	pagination.TotalItems = total
	pagination.TotalPages = int(math.Ceil(float64(total) / float64(pagination.PerPage)))

	entries := make([]SimplifiedScale, 0)
	if pagination.Offset() > total {
		return entries, nil, nil
	}

	pagination.NextPage = (pagination.Page + 1) % (pagination.TotalPages + 1)

	queryList := fmt.Sprintf(`
		SELECT
			s.id,
			s.name
		FROM scales s
			JOIN keys k ON s.id = k.scale_id
		WHERE %s
		ORDER BY
			s.id
		OFFSET ?
		LIMIT  ?;`, condition)

	args = append(args, pagination.Offset(), pagination.PerPage)
	if err := r.db.SelectContext(ctx, &entries, r.db.Rebind(queryList), args...); err != nil {
		return nil, nil, err
	}

	return entries, &pagination, nil
}

func (r theoryRepository) ListScaleKeys(ctx context.Context, scaleID int64) ([]SimplifiedKey, error) {
	query := `
		SELECT
			k.id,
			k.name
		FROM keys k
			JOIN scales s ON k.scale_id = s.id
			JOIN pitches p ON k.tonic_id = p.id
		WHERE
			scale_id = $1
		ORDER BY
			k.id;`

	entries := make([]SimplifiedKey, 0)
	if err := r.db.SelectContext(ctx, &entries, query, scaleID); err != nil {
		return nil, err
	}

	return entries, nil
}

func (r theoryRepository) ListScalePitches(ctx context.Context, scaleID int64) ([]SimplifiedPitch, error) {
	entries := make([]SimplifiedPitch, 0)

	query := `
		SELECT DISTINCT
			p.id,
			p.name
		FROM key_pitch_chords kpc
			JOIN pitches p ON kpc.pitch_id = p.id
			JOIN keys k ON kpc.key_id = k.id
		WHERE
			k.scale_id = $1
		ORDER BY
			p.id;`

	if err := r.db.SelectContext(ctx, &entries, query, scaleID); err != nil {
		return nil, err
	}

	return entries, nil
}

func (r theoryRepository) ListScaleChords(ctx context.Context, scaleID int64, filter ChordFilter, pagination api.Pagination) ([]SimplifiedChord, *api.Pagination, error) {
	pagination.Sanitize()

	args := []interface{}{scaleID}
	clauses := []string{"k.scale_id = ?"}

	if filter.ChordQualityID > 0 {
		args = append(args, filter.ChordQualityID)
		clauses = append(clauses, "c.chord_quality_id = ?")
	}

	if filter.RootID > 0 {
		args = append(args, filter.RootID)
		clauses = append(clauses, "c.root_id = ?")
	}

	if filter.Number > 0 {
		args = append(args, filter.Number)
		clauses = append(clauses, "c.number = ?")
	}

	if filter.Cardinality > 0 {
		args = append(args, filter.Cardinality)
		clauses = append(clauses, "cq.cardinality = ?")
	}

	queryCount := fmt.Sprintf(`
		SELECT
			COUNT(c.id)
		FROM key_pitch_chords kpc
			JOIN chords c ON kpc.chord_id = c.id
			JOIN chord_qualities cq ON c.chord_quality_id = cq.id
			JOIN pitches p ON c.root_id = p.id
			JOIN keys k ON kpc.key_id = k.id
		WHERE
			%s;`, strings.Join(clauses, " AND "))

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
		SELECT
			c.id,
			c.name
		FROM key_pitch_chords kpc
			JOIN chords c ON kpc.chord_id = c.id
			JOIN chord_qualities cq ON c.chord_quality_id = cq.id
			JOIN pitches p ON c.root_id = p.id
			JOIN keys k ON kpc.key_id = k.id
		WHERE
			%s
		ORDER BY
			c.id
		OFFSET ?
		LIMIT ?;`, strings.Join(clauses, " AND "))

	args = append(args, pagination.Offset(), pagination.PerPage)
	if err := r.db.SelectContext(ctx, &entries, r.db.Rebind(queryList), args...); err != nil {
		return nil, nil, err
	}

	return entries, &pagination, nil
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
