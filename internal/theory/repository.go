package theory

import (
	"context"
	"database/sql"
	"errors"
	"math"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Repository is music theory data repository
type Repository interface {
	ListPitches(ctx context.Context) ([]DetailedPitch, error)

	ListChords(ctx context.Context, pagination api.Pagination) ([]DetailedChord, *api.Pagination, error)
	ListChordPitches(ctx context.Context, chordID int64) ([]DetailedPitch, error)
	ListChordKeys(ctx context.Context, chordID int64, pagination api.Pagination) ([]DetailedKey, *api.Pagination, error)
	GetChord(ctx context.Context, chordID int64) (*DetailedChord, error)
	GetChordQuality(ctx context.Context, chordID int64) (*DetailedChordQuality, error)

	ListScales(ctx context.Context, pagination api.Pagination) ([]DetailedScale, *api.Pagination, error)
	ListScaleKeys(ctx context.Context, scaleID int64) ([]DetailedKey, error)
	GetScale(ctx context.Context, scaleID int64) (*DetailedScale, error)

	ListKeys(ctx context.Context, pagination api.Pagination) ([]DetailedKey, *api.Pagination, error)
	ListKeyModes(ctx context.Context, keyID int64) ([]DetailedKey, error)
	ListKeyChords(ctx context.Context, keyID int64, pagination api.Pagination) ([]DetailedChord, *api.Pagination, error)
	ListKeyPitches(ctx context.Context, keyID int64) ([]DetailedPitch, error)
	GetKey(ctx context.Context, keyID int64) (*DetailedKey, error)
}

// NewRepository returns an implementation of music theory data repository
func NewRepository(logger *zap.Logger, db *sqlx.DB) Repository {
	return theoryRepository{
		logger: logger,
		db:     db,
	}
}

type theoryRepository struct {
	logger *zap.Logger
	db     *sqlx.DB
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

func (r theoryRepository) ListChords(ctx context.Context, pagination api.Pagination) ([]DetailedChord, *api.Pagination, error) {
	queryCount := `
		SELECT
			COUNT(1)
		FROM chords;`

	var total int
	if err := r.db.GetContext(ctx, &total, queryCount); err != nil {
		return nil, nil, err
	}

	pagination.TotalItems = total
	pagination.TotalPages = int(math.Ceil(float64(total) / float64(pagination.PerPage)))

	entries := make([]DetailedChord, 0)
	if pagination.Offset() > total {
		return entries, nil, nil
	}

	pagination.NextPage = (pagination.Page + 1) % (pagination.TotalPages + 1)

	queryList := `
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
		ORDER BY
			c.id
		OFFSET $1
		LIMIT  $2;`

	if err := r.db.SelectContext(ctx, &entries, queryList, pagination.Offset(), pagination.PerPage); err != nil {
		return nil, nil, err
	}

	return entries, &pagination, nil
}

func (r theoryRepository) ListChordPitches(ctx context.Context, chordID int64) ([]DetailedPitch, error) {
	query := `
		SELECT
			p.id,
			p.name,
			p.number,
			p.frequency
		FROM music.public.chord_pitches cp
			JOIN pitches p ON cp.pitch_id = p.id
		WHERE
			cp.chord_id = $1
		ORDER BY
			p.id;`

	pitches := make([]DetailedPitch, 0)
	if err := r.db.SelectContext(ctx, &pitches, query, chordID); err != nil {
		return nil, err
	}

	return pitches, nil
}

func (r theoryRepository) ListChordKeys(ctx context.Context, chordID int64, pagination api.Pagination) ([]DetailedKey, *api.Pagination, error) {
	queryCount := `
		SELECT
			COUNT(1)
		FROM key_pitch_chords
		WHERE chord_id = $1;`

	var total int
	if err := r.db.GetContext(ctx, &total, queryCount, chordID); err != nil {
		return nil, nil, err
	}

	pagination.TotalItems = total
	pagination.TotalPages = int(math.Ceil(float64(total) / float64(pagination.PerPage)))

	entries := make([]DetailedKey, 0)
	if pagination.Offset() > total {
		return entries, nil, nil
	}

	pagination.NextPage = (pagination.Page + 1) % (pagination.TotalPages + 1)

	queryList := `
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
		FROM key_pitch_chords kpc
			JOIN keys k ON kpc.key_id = k.id
			JOIN scales s ON k.scale_id = s.id
			JOIN pitches p ON k.tonic_id = p.id
		WHERE
			kpc.chord_id = $1
		ORDER BY
			k.id
		OFFSET $2
		LIMIT  $3;`

	if err := r.db.SelectContext(ctx, &entries, queryList, chordID, pagination.Offset(), pagination.PerPage); err != nil {
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

	// get scale keys
	pitches, err := r.ListChordPitches(ctx, chordID)
	if err != nil {
		return nil, err
	}

	// map to simplified pitch
	mapped := make([]SimplifiedPitch, 0)
	for _, v := range pitches {
		mapped = append(mapped, SimplifiedPitch{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	chord.Pitches = mapped
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

func (r theoryRepository) ListScales(ctx context.Context, pagination api.Pagination) ([]DetailedScale, *api.Pagination, error) {
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

	// get scale keys
	keys, err := r.ListScaleKeys(ctx, scaleID)
	if err != nil {
		return nil, err
	}

	scale.Keys = keys
	return &scale, nil
}

func (r theoryRepository) ListKeys(ctx context.Context, pagination api.Pagination) ([]DetailedKey, *api.Pagination, error) {
	queryCount := `
		SELECT
			COUNT(1)
		FROM keys;`

	var total int
	if err := r.db.GetContext(ctx, &total, queryCount); err != nil {
		return nil, nil, err
	}

	pagination.TotalItems = total
	pagination.TotalPages = int(math.Ceil(float64(total) / float64(pagination.PerPage)))

	entries := make([]DetailedKey, 0)
	if pagination.Offset() > total {
		return entries, nil, nil
	}

	pagination.NextPage = (pagination.Page + 1) % (pagination.TotalPages + 1)

	queryList := `
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
		ORDER BY
			k.id
		OFFSET $1
		LIMIT  $2;`

	if err := r.db.SelectContext(ctx, &entries, queryList, pagination.Offset(), pagination.PerPage); err != nil {
		return nil, nil, err
	}

	return entries, &pagination, nil
}

func (r theoryRepository) ListKeyModes(ctx context.Context, keyID int64) ([]DetailedKey, error) {
	query := `
		WITH numbers AS(
			SELECT
				number
			FROM keys
			WHERE id = $1
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
			k.number IN (SELECT number FROM numbers)
		ORDER BY
			k.id;`

	entries := make([]DetailedKey, 0)
	if err := r.db.SelectContext(ctx, &entries, query, keyID); err != nil {
		return nil, err
	}

	return entries, nil
}

func (r theoryRepository) ListKeyChords(ctx context.Context, keyID int64, pagination api.Pagination) ([]DetailedChord, *api.Pagination, error) {
	queryCount := `
		SELECT
			COUNT(1)
		FROM key_pitch_chords
		WHERE key_id = $1;`

	var total int
	if err := r.db.GetContext(ctx, &total, queryCount, keyID); err != nil {
		return nil, nil, err
	}

	pagination.TotalItems = total
	pagination.TotalPages = int(math.Ceil(float64(total) / float64(pagination.PerPage)))

	entries := make([]DetailedChord, 0)
	if pagination.Offset() > total {
		return entries, nil, nil
	}

	pagination.NextPage = (pagination.Page + 1) % (pagination.TotalPages + 1)

	queryList := `
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
		WHERE kpc.key_id = $1
		ORDER BY
			c.id
		OFFSET $2
		LIMIT $3;`

	if err := r.db.SelectContext(ctx, &entries, queryList, keyID, pagination.Offset(), pagination.PerPage); err != nil {
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

	pitches, err := r.ListKeyPitches(ctx, keyID)
	if err != nil {
		return nil, err
	}

	// map to simplified pitches
	mapped := make([]SimplifiedPitch, 0)
	for _, v := range pitches {
		mapped = append(mapped, SimplifiedPitch{ID: v.ID, Name: v.Name})
	}

	entry.Pitches = mapped
	return &entry, nil
}
