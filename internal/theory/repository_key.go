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

type keyRepository interface {
	ListKeys(ctx context.Context, filter KeyFilter, pagination api.Pagination) ([]DetailedKey, *api.Pagination, error)
	ListKeyModes(ctx context.Context, keyID int64, filter KeyFilter) ([]DetailedKey, error)
	ListKeyChords(ctx context.Context, keyID int64, filter ChordFilter, pagination api.Pagination) ([]DetailedChord, *api.Pagination, error)
	ListKeyPitches(ctx context.Context, keyID int64) ([]DetailedPitch, error)
	GetKey(ctx context.Context, keyID int64) (*DetailedKey, error)
}

func (r theoryRepository) ListKeys(ctx context.Context, filter KeyFilter, pagination api.Pagination) ([]DetailedKey, *api.Pagination, error) {
	pagination.Sanitize()

	args := make([]interface{}, 0)
	clauses := make([]string, 0)

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

	condition := "TRUE"
	if len(clauses) > 0 {
		condition = strings.Join(clauses, " AND ")
	}

	queryCount := fmt.Sprintf(`
		SELECT
			COUNT(1)
		FROM keys k
			JOIN scales s ON k.scale_id = s.id
		WHERE
			%s;`, condition)

	var total int
	if err := r.db.GetContext(ctx, &total, r.db.Rebind(queryCount), args...); err != nil {
		return nil, nil, err
	}

	pagination.TotalItems = total
	pagination.TotalPages = int(math.Ceil(float64(total) / float64(pagination.PerPage)))

	entries := make([]DetailedKey, 0)
	if pagination.Offset() > total {
		return entries, nil, nil
	}

	pagination.NextPage = (pagination.Page + 1) % (pagination.TotalPages + 1)

	queryList := fmt.Sprintf(`
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
			%s
		ORDER BY
			k.id
		OFFSET ?
		LIMIT  ?;`, condition)

	args = append(args, pagination.Offset(), pagination.PerPage)
	if err := r.db.SelectContext(ctx, &entries, r.db.Rebind(queryList), args...); err != nil {
		return nil, nil, err
	}

	return entries, &pagination, nil
}

func (r theoryRepository) ListKeyModes(ctx context.Context, keyID int64, filter KeyFilter) ([]DetailedKey, error) {
	args := []interface{}{keyID}
	clauses := []string{"k.number IN (SELECT number FROM numbers)"}

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

	query := fmt.Sprintf(`
		WITH numbers AS(
			SELECT
				number
			FROM keys
			WHERE id = ?
		)
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
			%s
		ORDER BY
			k.id;`, strings.Join(clauses, " AND "))

	entries := make([]DetailedKey, 0)
	if err := r.db.SelectContext(ctx, &entries, r.db.Rebind(query), args...); err != nil {
		return nil, err
	}

	return entries, nil
}

func (r theoryRepository) ListKeyChords(ctx context.Context, keyID int64, filter ChordFilter, pagination api.Pagination) ([]DetailedChord, *api.Pagination, error) {
	pagination.Sanitize()

	args := []interface{}{keyID}
	clauses := []string{"kpc.key_id = ?"}

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
			COUNT(1)
		FROM key_pitch_chords kpc
			JOIN chords c ON kpc.chord_id = c.id
			JOIN chord_qualities cq ON c.chord_quality_id = cq.id
		WHERE
			%s;`, strings.Join(clauses, " AND "))

	var total int
	if err := r.db.GetContext(ctx, &total, r.db.Rebind(queryCount), args...); err != nil {
		return nil, nil, err
	}

	pagination.TotalItems = total
	pagination.TotalPages = int(math.Ceil(float64(total) / float64(pagination.PerPage)))

	entries := make([]DetailedChord, 0)
	if pagination.Offset() > total {
		return entries, nil, nil
	}

	pagination.NextPage = (pagination.Page + 1) % (pagination.TotalPages + 1)

	queryList := fmt.Sprintf(`
		SELECT
			c.id,
			cq.id   AS "quality.id",
			cq.name AS "quality.name",
			p.id    AS "root.id",
			p.name  AS "root.name",
			c.name,
			c.number
		FROM key_pitch_chords kpc
			JOIN chords c ON kpc.chord_id = c.id
			JOIN chord_qualities cq ON c.chord_quality_id = cq.id
			JOIN pitches p ON c.root_id = p.id
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

func (r theoryRepository) ListKeyPitches(ctx context.Context, keyID int64) ([]DetailedPitch, error) {
	query := `
		SELECT
			p.id,
			p.name,
			p.number,
			p.frequency
		FROM key_pitches k
			JOIN pitches p ON k.pitch_id = p.id
		WHERE
			k.key_id = $1
		ORDER BY
			p.id;`

	entries := make([]DetailedPitch, 0)
	if err := r.db.SelectContext(ctx, &entries, query, keyID); err != nil {
		return nil, err
	}

	return entries, nil
}

func (r theoryRepository) GetKey(ctx context.Context, keyID int64) (*DetailedKey, error) {
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
		WHERE k.id = $1;`

	var entry DetailedKey
	if err := r.db.GetContext(ctx, &entry, query, keyID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrKeyNotFound
		}

		return nil, err
	}

	return &entry, nil
}
