package theory_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestTheoryRepository_ListChords(t *testing.T) {
	type testCase struct {
		Title              string
		GivenFilter        theory.ChordFilter
		ExpectedCountQuery string
		ExpectedListQuery  string
		ExpectedCountArgs  []driver.Value
		ExpectedListArgs   []driver.Value
		CountChordsError   error
		ListChordsError    error
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordsWhenSucceededWithoutFilter",
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM chords c
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
				WHERE TRUE;`,
			ExpectedListQuery: `
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
				WHERE TRUE
				ORDER BY
					c.id
				OFFSET $1
				LIMIT  $2;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordsWhenSucceededWithChordQualityIDFilter",
			GivenFilter: theory.ChordFilter{ChordQualityID: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM chords c
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
				WHERE
					c.chord_quality_id = $1;`,
			ExpectedListQuery: `
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
					c.chord_quality_id = $1
				ORDER BY
					c.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordsWhenSucceededWithRootIDFilter",
			GivenFilter: theory.ChordFilter{RootID: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM chords c
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
				WHERE
					c.root_id = $1;`,
			ExpectedListQuery: `
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
					c.root_id = $1
				ORDER BY
					c.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordsWhenSucceededWithNumberFilter",
			GivenFilter: theory.ChordFilter{Number: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM chords c
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
				WHERE
					c.number = $1;`,
			ExpectedListQuery: `
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
					c.number = $1
				ORDER BY
					c.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordsWhenSucceededWithCardinalityFilter",
			GivenFilter: theory.ChordFilter{Cardinality: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM chords c
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
				WHERE
					cq.cardinality = $1;`,
			ExpectedListQuery: `
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
					cq.cardinality = $1
				ORDER BY
					c.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title: "ReturnsErrorWhenCountChordsFailed",
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM chords c
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
				WHERE TRUE;`,
			ExpectedListQuery: `
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
				WHERE TRUE
				ORDER BY
					c.id
				OFFSET $1
				LIMIT  $2;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			CountChordsError: sql.ErrConnDone,
		},
		{
			Title: "ReturnsErrorWhenListChordsFailed",
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM chords c
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
				WHERE TRUE;`,
			ExpectedListQuery: `
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
				WHERE TRUE
				ORDER BY
					c.id
				OFFSET $1
				LIMIT  $2;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ListChordsError:  sql.ErrConnDone,
		},
	}

	countChordsColumns := []string{"count"}

	listChordsColumns := []string{
		"id",
		"quality.id",
		"quality.name",
		"root.id",
		"root.name",
		"name",
		"number",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			db, sqlMock, err := mockDatabase()
			require.NoError(t, err)

			if tc.CountChordsError != nil {
				sqlMock.ExpectQuery(tc.ExpectedCountQuery).
					WithArgs(tc.ExpectedCountArgs...).
					WillReturnError(tc.CountChordsError)
			} else {
				sqlMock.ExpectQuery(tc.ExpectedCountQuery).
					WithArgs(tc.ExpectedCountArgs...).
					WillReturnRows(sqlmock.NewRows(countChordsColumns).
						AddRow(1))

				if tc.ListChordsError != nil {
					sqlMock.ExpectQuery(tc.ExpectedListQuery).
						WithArgs(tc.ExpectedListArgs...).
						WillReturnError(tc.ListChordsError)
				} else {
					sqlMock.ExpectQuery(tc.ExpectedListQuery).
						WithArgs(tc.ExpectedListArgs...).
						WillReturnRows(sqlmock.NewRows(listChordsColumns).AddRow(1, 2, "name1", 3, "name2", "name3", 4))
				}
			}

			repository := theory.NewRepository(logger, db)
			chords, _, err := repository.ListChords(context.Background(), tc.GivenFilter, api.Pagination{})
			if strings.HasPrefix(tc.Title, "ReturnsError") {
				require.Error(t, err)
				require.Empty(t, chords)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, chords)
			}
		})
	}
}

func TestTheoryRepository_GetChord(t *testing.T) {
	type testCase struct {
		Title                 string
		GetChordError         error
		ListChordPitchesError error
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordWhenSucceeded",
		},
		{
			Title:         "ReturnsErrorWhenChordNotFound",
			GetChordError: sql.ErrNoRows,
		},
		{
			Title:         "ReturnsErrorWhenGetChordFailed",
			GetChordError: sql.ErrConnDone,
		},
		{
			Title:                 "ReturnsErrorWhenListChordPitchesFailed",
			ListChordPitchesError: sql.ErrConnDone,
		},
	}

	getChordQuery := `
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

	getChordColumns := []string{
		"id",
		"quality.id",
		"quality.name",
		"root.id",
		"root.name",
		"name",
		"number",
	}

	listChordPitchesQuery := `
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

	listChordPitchesColumns := []string{
		"id",
		"name",
		"number",
		"frequency",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			db, sqlMock, err := mockDatabase()
			require.NoError(t, err)

			if tc.GetChordError != nil {
				sqlMock.ExpectQuery(getChordQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(tc.GetChordError)
			} else {
				sqlMock.ExpectQuery(getChordQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows(getChordColumns).
						AddRow(1, 2, "name1", 3, "name2", "name3", 4))

				if tc.ListChordPitchesError != nil {
					sqlMock.ExpectQuery(listChordPitchesQuery).
						WithArgs(sqlmock.AnyArg()).
						WillReturnError(tc.ListChordPitchesError)
				} else {
					sqlMock.ExpectQuery(listChordPitchesQuery).
						WithArgs(sqlmock.AnyArg()).
						WillReturnRows(sqlmock.NewRows(listChordPitchesColumns).
							AddRow(1, "name", 2, 3.0))
				}
			}

			repository := theory.NewRepository(logger, db)
			chord, err := repository.GetChord(context.Background(), 1)
			if strings.HasPrefix(tc.Title, "ReturnsError") {
				require.Error(t, err)
				require.Empty(t, chord)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, chord)
			}
		})
	}
}

func TestTheoryRepository_GetChordQuality(t *testing.T) {
	type testCase struct {
		Title      string
		GivenError error
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordQualityWhenSucceeded",
		},
		{
			Title:      "ReturnsErrorWhenFailed",
			GivenError: sql.ErrNoRows,
		},
	}

	getChordQualityQuery := `
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

	getChordQualityColumns := []string{
		"id",
		"name",
		"cardinality",
		"number",
		"pitch_class",
		"interval_pattern",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			db, sqlMock, err := mockDatabase()
			require.NoError(t, err)

			if tc.GivenError != nil {
				sqlMock.ExpectQuery(getChordQualityQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(tc.GivenError)
			} else {
				sqlMock.ExpectQuery(getChordQualityQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows(getChordQualityColumns).
						AddRow(1, "name", 2, 3, []byte("[1,2,3]"), []byte("[1,2]")))
			}

			repository := theory.NewRepository(logger, db)
			chordQuality, err := repository.GetChordQuality(context.Background(), 1)
			if strings.HasPrefix(tc.Title, "ReturnsError") {
				require.Error(t, err)
				require.Empty(t, chordQuality)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, chordQuality)
			}
		})
	}
}

func TestTheoryRepository_ListChordKeys(t *testing.T) {
	type testCase struct {
		Title               string
		GivenFilter         theory.KeyFilter
		ExpectedCountQuery  string
		ExpectedListQuery   string
		ExpectedCountArgs   []driver.Value
		ExpectedListArgs    []driver.Value
		CountChordKeysError error
		ListChordKeysError  error
	}

	intValue := 1
	boolValue := true
	testCases := []testCase{
		{
			Title: "ReturnsChordKeysWhenSucceededWithoutFilter",
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM key_pitch_chords kpc
					JOIN keys k ON kpc.key_id = k.id
					JOIN scales s ON k.scale_id = s.id
				WHERE kpc.chord_id = $1;`,
			ExpectedListQuery: `
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
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordKeysWhenSucceededWithScaleIDFilter",
			GivenFilter: theory.KeyFilter{ScaleID: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM key_pitch_chords kpc
					JOIN keys k ON kpc.key_id = k.id
					JOIN scales s ON k.scale_id = s.id
				WHERE
					kpc.chord_id = $1 AND
					k.scale_id   = $2;`,
			ExpectedListQuery: `
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
					kpc.chord_id = $1 AND
					k.scale_id   = $2
				ORDER BY
					k.id
				OFFSET $3
				LIMIT  $4;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordKeysWhenSucceededWithTonicIDFilter",
			GivenFilter: theory.KeyFilter{TonicID: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM key_pitch_chords kpc
					JOIN keys k ON kpc.key_id = k.id
					JOIN scales s ON k.scale_id = s.id
				WHERE
					kpc.chord_id = $1 AND
					k.tonic_id   = $2;`,
			ExpectedListQuery: `
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
					kpc.chord_id = $1 AND
					k.tonic_id   = $2
				ORDER BY
					k.id
				OFFSET $3
				LIMIT  $4;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordKeysWhenSucceededWithNumberFilter",
			GivenFilter: theory.KeyFilter{Number: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM key_pitch_chords kpc
					JOIN keys k ON kpc.key_id = k.id
					JOIN scales s ON k.scale_id = s.id
				WHERE
					kpc.chord_id = $1 AND
					k.number     = $2;`,
			ExpectedListQuery: `
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
					kpc.chord_id = $1 AND
					k.number     = $2
				ORDER BY
					k.id
				OFFSET $3
				LIMIT  $4;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordKeysWhenSucceededWithPerfectionFilter",
			GivenFilter: theory.KeyFilter{Perfection: &intValue},
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM key_pitch_chords kpc
					JOIN keys k ON kpc.key_id = k.id
					JOIN scales s ON k.scale_id = s.id
				WHERE
					kpc.chord_id = $1 AND
					s.perfection = $2;`,
			ExpectedListQuery: `
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
					kpc.chord_id  = $1 AND
					s.perfection  = $2
				ORDER BY
					k.id
				OFFSET $3
				LIMIT  $4;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordKeysWhenSucceededWithImperfectionFilter",
			GivenFilter: theory.KeyFilter{Imperfection: &intValue},
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM key_pitch_chords kpc
					JOIN keys k ON kpc.key_id = k.id
					JOIN scales s ON k.scale_id = s.id
				WHERE
					kpc.chord_id = $1 AND
					s.imperfection = $2;`,
			ExpectedListQuery: `
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
					kpc.chord_id  = $1 AND
					s.imperfection  = $2
				ORDER BY
					k.id
				OFFSET $3
				LIMIT  $4;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordKeysWhenSucceededWithBalancedFilter",
			GivenFilter: theory.KeyFilter{Balanced: &boolValue},
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM key_pitch_chords kpc
					JOIN keys k ON kpc.key_id = k.id
					JOIN scales s ON k.scale_id = s.id
				WHERE
					kpc.chord_id = $1 AND
					k.balanced   = $2;`,
			ExpectedListQuery: `
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
					kpc.chord_id = $1 AND
					k.balanced   = $2
				ORDER BY
					k.id
				OFFSET $3
				LIMIT  $4;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordKeysWhenSucceededWithRotationalSymmetricFilter",
			GivenFilter: theory.KeyFilter{RotationalSymmetric: &boolValue},
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM key_pitch_chords kpc
					JOIN keys k ON kpc.key_id = k.id
					JOIN scales s ON k.scale_id = s.id
				WHERE
					kpc.chord_id           = $1 AND
					s.rotational_symmetric = $2;`,
			ExpectedListQuery: `
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
					kpc.chord_id           = $1 AND
					s.rotational_symmetric = $2
				ORDER BY
					k.id
				OFFSET $3
				LIMIT  $4;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordKeysWhenSucceededWithRotationalSymmetryLevelFilter",
			GivenFilter: theory.KeyFilter{RotationalSymmetryLevel: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM key_pitch_chords kpc
					JOIN keys k ON kpc.key_id = k.id
					JOIN scales s ON k.scale_id = s.id
				WHERE
					kpc.chord_id                = $1 AND
					s.rotational_symmetry_level = $2;`,
			ExpectedListQuery: `
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
					kpc.chord_id                = $1 AND
					s.rotational_symmetry_level = $2
				ORDER BY
					k.id
				OFFSET $3
				LIMIT  $4;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordKeysWhenSucceededWithReflectionalSymmetricFilter",
			GivenFilter: theory.KeyFilter{ReflectionalSymmetric: &boolValue},
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM key_pitch_chords kpc
					JOIN keys k ON kpc.key_id = k.id
					JOIN scales s ON k.scale_id = s.id
				WHERE
					kpc.chord_id           = $1 AND
					s.reflectional_symmetric = $2;`,
			ExpectedListQuery: `
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
					kpc.chord_id             = $1 AND
					s.reflectional_symmetric = $2
				ORDER BY
					k.id
				OFFSET $3
				LIMIT  $4;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordKeysWhenSucceededWithPalindromicFilter",
			GivenFilter: theory.KeyFilter{Palindromic: &boolValue},
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM key_pitch_chords kpc
					JOIN keys k ON kpc.key_id = k.id
					JOIN scales s ON k.scale_id = s.id
				WHERE
					kpc.chord_id = $1 AND
					s.palindromic = $2;`,
			ExpectedListQuery: `
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
					kpc.chord_id  = $1 AND
					s.palindromic = $2
				ORDER BY
					k.id
				OFFSET $3
				LIMIT  $4;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordKeysWhenSucceededWithCardinalityFilter",
			GivenFilter: theory.KeyFilter{Cardinality: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM key_pitch_chords kpc
					JOIN keys k ON kpc.key_id = k.id
					JOIN scales s ON k.scale_id = s.id
				WHERE
					kpc.chord_id = $1 AND
					s.cardinality = $2;`,
			ExpectedListQuery: `
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
					kpc.chord_id  = $1 AND
					s.cardinality = $2
				ORDER BY
					k.id
				OFFSET $3
				LIMIT  $4;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title: "ReturnsErrorWhenCountChordKeysFailed",
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM key_pitch_chords kpc
					JOIN keys k ON kpc.key_id = k.id
					JOIN scales s ON k.scale_id = s.id
				WHERE kpc.chord_id = $1;`,
			ExpectedListQuery: `
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
				LIMIT  $3;`,
			ExpectedCountArgs:   []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:    []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
			CountChordKeysError: sql.ErrConnDone,
		},
		{
			Title: "ReturnsErrorWhenListChordKeysFailed",
			ExpectedCountQuery: `
				SELECT
					COUNT(1)
				FROM key_pitch_chords kpc
					JOIN keys k ON kpc.key_id = k.id
					JOIN scales s ON k.scale_id = s.id
				WHERE kpc.chord_id = $1;`,
			ExpectedListQuery: `
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
				LIMIT  $3;`,
			ExpectedCountArgs:  []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:   []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
			ListChordKeysError: sql.ErrConnDone,
		},
	}

	countChordKeysColumns := []string{"count"}

	listChordKeysColumns := []string{
		"id",
		"scale.id",
		"scale.name",
		"tonic.id",
		"tonic.name",
		"name",
		"number",
		"balanced",
		"center_x",
		"center_y",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			db, sqlMock, err := mockDatabase()
			require.NoError(t, err)

			if tc.CountChordKeysError != nil {
				sqlMock.ExpectQuery(tc.ExpectedCountQuery).
					WithArgs(tc.ExpectedCountArgs...).
					WillReturnError(tc.CountChordKeysError)
			} else {
				sqlMock.ExpectQuery(tc.ExpectedCountQuery).
					WithArgs(tc.ExpectedCountArgs...).
					WillReturnRows(sqlmock.NewRows(countChordKeysColumns).
						AddRow(1))

				if tc.ListChordKeysError != nil {
					sqlMock.ExpectQuery(tc.ExpectedListQuery).
						WithArgs(tc.ExpectedListArgs...).
						WillReturnError(tc.ListChordKeysError)
				} else {
					sqlMock.ExpectQuery(tc.ExpectedListQuery).
						WithArgs(tc.ExpectedListArgs...).
						WillReturnRows(sqlmock.NewRows(listChordKeysColumns).
							AddRow(1, 2, "name1", 3, "name2", "name3", 4, true, 5.0, 6.0))
				}
			}

			var pagination api.Pagination
			repository := theory.NewRepository(logger, db)
			keys, _, err := repository.ListChordKeys(context.Background(), 1, tc.GivenFilter, pagination)
			if strings.HasPrefix(tc.Title, "ReturnsError") {
				require.Error(t, err)
				require.Empty(t, keys)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, keys)
			}
		})
	}
}

func TestTheoryRepository_ListChordPitches(t *testing.T) {
	type testCase struct {
		Title                 string
		ListChordPitchesError error
	}

	testCases := []testCase{
		{
			Title: "ReturnsPitchesWhenSucceeded",
		},
		{
			Title:                 "ReturnsErrorWhenFailed",
			ListChordPitchesError: sql.ErrConnDone,
		},
	}

	listChordPitchesQuery := `
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

	listChordPitchesColumns := []string{
		"id",
		"name",
		"number",
		"frequency",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			db, sqlMock, err := mockDatabase()
			require.NoError(t, err)

			if tc.ListChordPitchesError != nil {
				sqlMock.ExpectQuery(listChordPitchesQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(tc.ListChordPitchesError)
			} else {
				sqlMock.ExpectQuery(listChordPitchesQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows(listChordPitchesColumns).
						AddRow(1, "name", 2, 3.0))
			}

			repository := theory.NewRepository(logger, db)
			pitches, err := repository.ListChordPitches(context.Background(), 1)
			if strings.HasPrefix(tc.Title, "ReturnsError") {
				require.Error(t, err)
				require.Empty(t, pitches)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, pitches)
			}
		})
	}
}
