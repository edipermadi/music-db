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

type chordRepository interface {
	GetChord(ctx context.Context, chordID int64) (*DetailedChord, error)
	GetChordQuality(ctx context.Context, chordID int64) (*DetailedChordQuality, error)
	ListChordKeys(ctx context.Context, chordID int64, filter KeyFilter, pagination api.Pagination) ([]SimplifiedKey, *api.Pagination, error)
	ListChordPitches(ctx context.Context, chordID int64) ([]SimplifiedPitch, error)
	ListChordScales(ctx context.Context, chordID int64, filter ScaleFilter, pagination api.Pagination) ([]SimplifiedScale, *api.Pagination, error)
	ListChords(ctx context.Context, filter ChordFilter, pagination api.Pagination) ([]SimplifiedChord, *api.Pagination, error)
}

func (r theoryRepository) ListChords(ctx context.Context, filter ChordFilter, pagination api.Pagination) ([]SimplifiedChord, *api.Pagination, error) {
	pagination.Sanitize()

	args := make([]interface{}, 0)
	clauses := make([]string, 0)

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

	condition := "TRUE"
	if len(clauses) > 0 {
		condition = strings.Join(clauses, " AND ")
	}

	queryCount := fmt.Sprintf(`
		SELECT
			COUNT(1)
		FROM chords c
			JOIN chord_qualities cq ON c.chord_quality_id = cq.id
		WHERE
			%s;`, condition)

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
		FROM chords c 
			JOIN chord_qualities cq ON c.chord_quality_id = cq.id
			JOIN pitches p ON c.root_id = p.id
		WHERE
			%s
		ORDER BY
			c.id
		OFFSET ?
		LIMIT  ?;`, condition)

	args = append(args, pagination.Offset(), pagination.PerPage)
	if err := r.db.SelectContext(ctx, &entries, r.db.Rebind(queryList), args...); err != nil {
		return nil, nil, err
	}

	return entries, &pagination, nil
}

func (r theoryRepository) ListChordPitches(ctx context.Context, chordID int64) ([]SimplifiedPitch, error) {
	query := `
		SELECT DISTINCT 
			p.id,
			p.name
		FROM music.public.chord_pitches cp
			JOIN pitches p ON cp.pitch_id = p.id
		WHERE
			cp.chord_id = $1
		ORDER BY
			p.id;`

	pitches := make([]SimplifiedPitch, 0)
	if err := r.db.SelectContext(ctx, &pitches, query, chordID); err != nil {
		return nil, err
	}

	return pitches, nil
}

func (r theoryRepository) ListChordKeys(ctx context.Context, chordID int64, filter KeyFilter, pagination api.Pagination) ([]SimplifiedKey, *api.Pagination, error) {
	pagination.Sanitize()

	args := []interface{}{chordID}
	clauses := []string{"kpc.chord_id = ?"}

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
			COUNT(1)
		FROM key_pitch_chords kpc
			JOIN keys k ON kpc.key_id = k.id
			JOIN scales s ON k.scale_id = s.id
		WHERE
			%s;`, strings.Join(clauses, " AND "))

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
		SELECT
			k.id,
			k.name
		FROM key_pitch_chords kpc
			JOIN keys k ON kpc.key_id = k.id
			JOIN scales s ON k.scale_id = s.id
			JOIN pitches p ON k.tonic_id = p.id
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

func (r theoryRepository) ListChordScales(ctx context.Context, chordID int64, filter ScaleFilter, pagination api.Pagination) ([]SimplifiedScale, *api.Pagination, error) {
	pagination.Sanitize()

	args := []interface{}{chordID}
	clauses := []string{"kpc.chord_id = ?"}

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
			COUNT(s.id)
		FROM key_pitch_chords kpc
			JOIN keys k ON kpc.key_id = k.id
			JOIN scales s ON k.scale_id = s.id
		WHERE
			%s;`, strings.Join(clauses, " AND "))

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
		FROM key_pitch_chords kpc
			JOIN keys k ON kpc.key_id = k.id
			JOIN scales s ON k.scale_id = s.id
			JOIN pitches p ON k.tonic_id = p.id
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

func (r theoryRepository) GetChord(ctx context.Context, chordID int64) (*DetailedChord, error) {
	query := `
		SELECT
			c.id,
			cq.id   AS "quality.id",
			cq.name AS "quality.name",
			p.id    AS "root.id",
			p.name  AS "root.name",
			c.name,
			c.number
		FROM chords c 
			JOIN chord_qualities cq ON c.chord_quality_id = cq.id
			JOIN pitches p ON c.root_id = p.id
		WHERE
			c.id = $1;`

	var chord DetailedChord
	if err := r.db.GetContext(ctx, &chord, query, chordID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrChordNotFound
		}
		return nil, err
	}

	return &chord, nil
}

func (r theoryRepository) GetChordQuality(ctx context.Context, chordID int64) (*DetailedChordQuality, error) {
	query := `
		SELECT
			cq.id,
			cq.name,
			cq.cardinality,
			cq.number,
			cq.pitch_class,
			cq.interval_pattern
		FROM chords c 
			JOIN chord_qualities cq ON c.chord_quality_id = cq.id
			JOIN pitches p ON c.root_id = p.id
		WHERE
			c.id = $1;`

	var quality DetailedChordQuality
	if err := r.db.GetContext(ctx, &quality, query, chordID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrChordQualityNotFound
		}
		return nil, err
	}

	return &quality, nil
}
