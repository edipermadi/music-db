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
	ListPitchKeys(ctx context.Context, pitchID int64, filter KeyFilter, pagination api.Pagination) ([]SimplifiedKey, *api.Pagination, error)
	ListPitchScales(ctx context.Context, pitchID int64, filter ScaleFilter, pagination api.Pagination) ([]SimplifiedScale, *api.Pagination, error)
	ListPitches(ctx context.Context) ([]SimplifiedPitch, error)
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

	clauses := []string{"cp.pitch_id = ?"}
	args := []interface{}{pitchID}

	if filter.ChordQualityID > 0 {
		clauses = append(clauses, "c.chord_quality_id = ?")
		args = append(args, filter.ChordQualityID)
	}

	if filter.RootID > 0 {
		clauses = append(clauses, "c.root_id = ?")
		args = append(args, filter.RootID)
	}

	if filter.Number > 0 {
		clauses = append(clauses, "c.number = ?")
		args = append(args, filter.Number)
	}

	if filter.Cardinality > 0 {
		clauses = append(clauses, "cq.cardinality = ?")
		args = append(args, filter.Cardinality)
	}

	queryCount := fmt.Sprintf(`
		SELECT 
			COUNT(DISTINCT c.id)
		FROM chord_pitches cp
			JOIN chords c ON cp.chord_id = c.id
			JOIN chord_qualities cq ON c.chord_quality_id = cq.id
		WHERE %s
		GROUP BY
		    cp.pitch_id;`, strings.Join(clauses, " AND "))

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
			JOIN chords c ON cp.chord_id = c.id
			JOIN chord_qualities cq ON c.chord_quality_id = cq.id
		WHERE
		    %s
		ORDER BY
		    c.id
		OFFSET ?
		LIMIT  ?;`, strings.Join(clauses, " AND "))

	args = append(args, pagination.Offset(), pagination.PerPage)
	if err := r.db.SelectContext(ctx, &entries, r.db.Rebind(queryList), args...); err != nil {
		return nil, nil, err
	}

	return entries, &pagination, nil
}

func (r theoryRepository) ListPitchKeys(ctx context.Context, pitchID int64, filter KeyFilter, pagination api.Pagination) ([]SimplifiedKey, *api.Pagination, error) {
	pagination.Sanitize()

	clauses := []string{"kp.pitch_id = ?"}
	args := []interface{}{pitchID}

	if filter.ScaleID > 0 {
		args = append(args, filter.ScaleID)
		clauses = append(clauses, "k.scale_id = ?")
	}

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

	queryCount := fmt.Sprintf(`
		SELECT 
			COUNT(DISTINCT k.id)
		FROM key_pitches kp
			JOIN keys k ON kp.key_id = k.id
			JOIN scales s ON k.scale_id = s.id
		WHERE %s
		GROUP BY
		    kp.pitch_id;`, strings.Join(clauses, " AND "))

	var total int
	if err := r.db.GetContext(ctx, &total, r.db.Rebind(queryCount), args...); err != nil {
		return nil, nil, err
	}

	pagination.TotalItems = total
	pagination.TotalPages = int(math.Ceil(float64(total) / float64(pagination.PerPage)))

	entries := make([]SimplifiedKey, 0)
	if pagination.Offset() > total {
		return entries, nil, nil
	}

	pagination.NextPage = (pagination.Page + 1) % (pagination.TotalPages + 1)

	queryList := fmt.Sprintf(`
		SELECT DISTINCT
			k.id, 
			k.name
		FROM key_pitches kp
			JOIN keys k ON kp.key_id = k.id
			JOIN scales s ON k.scale_id = s.id
		WHERE
		    %s
		ORDER BY
		    k.id
		OFFSET ?
		LIMIT  ?;`, strings.Join(clauses, " AND "))

	args = append(args, pagination.Offset(), pagination.PerPage)
	if err := r.db.SelectContext(ctx, &entries, r.db.Rebind(queryList), args...); err != nil {
		return nil, nil, err
	}

	return entries, &pagination, nil
}

func (r theoryRepository) ListPitchScales(ctx context.Context, pitchID int64, filter ScaleFilter, pagination api.Pagination) ([]SimplifiedScale, *api.Pagination, error) {
	pagination.Sanitize()

	clauses := []string{"kp.pitch_id = ?"}
	args := []interface{}{pitchID}

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

	queryCount := fmt.Sprintf(`
		SELECT 
			COUNT(DISTINCT s.id)
		FROM key_pitches kp
			JOIN keys k ON kp.key_id = k.id
			JOIN scales s ON k.scale_id = s.id
		WHERE %s
		GROUP BY
		    kp.pitch_id;`, strings.Join(clauses, " AND "))

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
		SELECT DISTINCT
			s.id, 
			s.name
		FROM key_pitches kp
			JOIN keys k ON kp.key_id = k.id
			JOIN scales s ON k.scale_id = s.id
		WHERE
		    %s
		ORDER BY
		    k.id
		OFFSET ?
		LIMIT  ?;`, strings.Join(clauses, " AND "))

	args = append(args, pagination.Offset(), pagination.PerPage)
	if err := r.db.SelectContext(ctx, &entries, r.db.Rebind(queryList), args...); err != nil {
		return nil, nil, err
	}

	return entries, &pagination, nil
}

func (r theoryRepository) ListPitches(ctx context.Context) ([]SimplifiedPitch, error) {
	query := `
		SELECT
			id,
			name
		FROM pitches
		ORDER BY id;`

	entries := make([]SimplifiedPitch, 0)
	if err := r.db.SelectContext(ctx, &entries, query); err != nil {
		return nil, err
	}

	return entries, nil
}
