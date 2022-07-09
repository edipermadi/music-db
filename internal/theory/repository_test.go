package theory_test

import (
	"context"
	"database/sql"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestNewRepository(t *testing.T) {
	logger, err := zap.NewProduction()
	require.NoError(t, err)

	db, _, err := mockDatabase()
	require.NoError(t, err)

	repository := theory.NewRepository(logger, db)
	require.Implements(t, (*theory.Repository)(nil), repository)
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
		CountChordKeysError error
		ListChordKeysError  error
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordKeysWhenSucceeded",
		},
		{
			Title:               "ReturnsErrorWhenCountChordKeysFailed",
			CountChordKeysError: sql.ErrConnDone,
		},
		{
			Title:              "ReturnsErrorWhenListChordKeysFailed",
			ListChordKeysError: sql.ErrConnDone,
		},
	}

	countChordKeysQuery := `
		SELECT
			COUNT(1)
		FROM key_pitch_chords
		WHERE chord_id = $1;`

	countChordKeysColumns := []string{"count"}

	listChordKeysQuery := `
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
				sqlMock.ExpectQuery(countChordKeysQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(tc.CountChordKeysError)
			} else {
				sqlMock.ExpectQuery(countChordKeysQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows(countChordKeysColumns).
						AddRow(1))

				if tc.ListChordKeysError != nil {
					sqlMock.ExpectQuery(listChordKeysQuery).
						WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnError(tc.ListChordKeysError)
				} else {
					sqlMock.ExpectQuery(listChordKeysQuery).
						WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnRows(sqlmock.NewRows(listChordKeysColumns).
							AddRow(1, 2, "name1", 3, "name2", "name3", 4, true, 5.0, 6.0))
				}
			}

			var pagination api.Pagination
			repository := theory.NewRepository(logger, db)
			keys, _, err := repository.ListChordKeys(context.Background(), 1, pagination)
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

func TestTheoryRepository_GetKey(t *testing.T) {
	type testCase struct {
		Title               string
		GetKeyError         error
		ListKeyPitchesError error
	}

	testCases := []testCase{
		{
			Title: "ReturnsKeyWhenSucceeded",
		},
		{
			Title:       "ReturnsErrorWhenKeyNotFound",
			GetKeyError: sql.ErrNoRows,
		},
		{
			Title:       "ReturnsErrorWhenGetKeyFailed",
			GetKeyError: sql.ErrConnDone,
		},
		{
			Title:               "ReturnsErrorWhenListKeyPitchesFailed",
			ListKeyPitchesError: sql.ErrConnDone,
		},
	}

	getKeyQuery := `
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

	getKeyColumns := []string{
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

	listKeyPitchesQuery := `
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

	listKeyPitchesColumns := []string{
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

			if tc.GetKeyError != nil {
				sqlMock.ExpectQuery(getKeyQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(tc.GetKeyError)
			} else {
				sqlMock.ExpectQuery(getKeyQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows(getKeyColumns).
						AddRow(1, 2, "name1", 3, "name2", "name3", 4, true, 5.0, 6.0))

				if tc.ListKeyPitchesError != nil {
					sqlMock.ExpectQuery(listKeyPitchesQuery).
						WithArgs(sqlmock.AnyArg()).
						WillReturnError(tc.ListKeyPitchesError)
				} else {
					sqlMock.ExpectQuery(listKeyPitchesQuery).
						WithArgs(sqlmock.AnyArg()).
						WillReturnRows(sqlmock.NewRows(listKeyPitchesColumns).
							AddRow(1, "name", 2, 3.0))
				}
			}

			repository := theory.NewRepository(logger, db)
			key, err := repository.GetKey(context.Background(), 1)
			if strings.HasPrefix(tc.Title, "ReturnsError") {
				require.Error(t, err)
				require.Empty(t, key)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, key)
			}
		})
	}
}

func TestTheoryRepository_GetScale(t *testing.T) {
	type testCase struct {
		Title              string
		GetScaleError      error
		ListScaleKeysError error
	}

	testCases := []testCase{
		{
			Title: "ReturnsScaleWhenSucceeded",
		},
		{
			Title:         "ReturnsErrorWhenScaleNotFound",
			GetScaleError: sql.ErrNoRows,
		},
		{
			Title:         "ReturnsErrorWhenGetScaleFailed",
			GetScaleError: sql.ErrConnDone,
		},
		{
			Title:              "ReturnsErrorWhenListScaleKeysFailed",
			ListScaleKeysError: sql.ErrConnDone,
		},
	}

	getScaleQuery := `
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

	getScaleColumns := []string{
		"id",
		"name",
		"cardinality",
		"number",
		"perfection",
		"imperfection",
		"pitch_class",
		"interval_pattern",
		"rotational_symmetric",
		"rotational_symmetry_level",
		"palindromic",
		"reflectional_symmetric",
		"reflectional_symmetry_axes",
		"balanced",
	}

	listScaleKeysQuery := `
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

	listScaleKeysColumns := []string{
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

			if tc.GetScaleError != nil {
				sqlMock.ExpectQuery(getScaleQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(tc.GetScaleError)
			} else {
				sqlMock.ExpectQuery(getScaleQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows(getScaleColumns).
						AddRow(1, "name", 2, 3, 4, 5, []byte("[6,7]"), []byte("[8,9]"), true, 10, true, true, []byte("[10,11]"), true))

				if tc.ListScaleKeysError != nil {
					sqlMock.ExpectQuery(listScaleKeysQuery).
						WithArgs(sqlmock.AnyArg()).
						WillReturnError(tc.ListScaleKeysError)
				} else {
					sqlMock.ExpectQuery(listScaleKeysQuery).
						WithArgs(sqlmock.AnyArg()).
						WillReturnRows(sqlmock.NewRows(listScaleKeysColumns).
							AddRow(1, 2, "name1", 3, "name2", "name2", 4, true, 5.0, 6.0))
				}
			}

			repository := theory.NewRepository(logger, db)
			scale, err := repository.GetScale(context.Background(), 1)
			if strings.HasPrefix(tc.Title, "ReturnsError") {
				require.Error(t, err)
				require.Empty(t, scale)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, scale)
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

func TestTheoryRepository_ListChords(t *testing.T) {
	type testCase struct {
		Title            string
		CountChordsError error
		ListChordsError  error
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
		},
		{
			Title:            "ReturnsErrorWhenCountChordsFailed",
			CountChordsError: sql.ErrConnDone,
		},
		{
			Title:           "ReturnsErrorWhenListChordsFailed",
			ListChordsError: sql.ErrConnDone,
		},
	}

	countChordsQuery := `
		SELECT
			COUNT(1)
		FROM chords;`

	countChordsColumns := []string{"count"}

	listChordQuery := `
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
				sqlMock.ExpectQuery(countChordsQuery).
					WithArgs().
					WillReturnError(tc.CountChordsError)
			} else {
				sqlMock.ExpectQuery(countChordsQuery).
					WithArgs().
					WillReturnRows(sqlmock.NewRows(countChordsColumns).
						AddRow(1))

				if tc.ListChordsError != nil {
					sqlMock.ExpectQuery(listChordQuery).
						WithArgs(sqlmock.AnyArg()).
						WillReturnError(tc.ListChordsError)
				} else {
					sqlMock.ExpectQuery(listChordQuery).
						WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnRows(sqlmock.NewRows(listChordsColumns).AddRow(1, 2, "name1", 3, "name2", "name3", 4))
				}
			}

			var pagination api.Pagination
			repository := theory.NewRepository(logger, db)
			chords, _, err := repository.ListChords(context.Background(), pagination)
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

func TestTheoryRepository_ListKeyChords(t *testing.T) {
	type testCase struct {
		Title               string
		CountKeyChordsError error
		ListKeyChordsError  error
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
		},
		{
			Title:               "ReturnsErrorWhenCountKeyChordsFailed",
			CountKeyChordsError: sql.ErrConnDone,
		},
		{
			Title:              "ReturnsErrorWhenListKeyChordsFailed",
			ListKeyChordsError: sql.ErrConnDone,
		},
	}

	countChordsQuery := `
		SELECT
			COUNT(1)
		FROM key_pitch_chords
		WHERE key_id = $1;`

	countChordsColumns := []string{"count"}

	listChordQuery := `
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
			kpc.key_id = $1
		ORDER BY
			c.id
		OFFSET $2
		LIMIT  $3;`

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

			if tc.CountKeyChordsError != nil {
				sqlMock.ExpectQuery(countChordsQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(tc.CountKeyChordsError)
			} else {
				sqlMock.ExpectQuery(countChordsQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows(countChordsColumns).
						AddRow(1))

				if tc.ListKeyChordsError != nil {
					sqlMock.ExpectQuery(listChordQuery).
						WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnError(tc.ListKeyChordsError)
				} else {
					sqlMock.ExpectQuery(listChordQuery).
						WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnRows(sqlmock.NewRows(listChordsColumns).AddRow(1, 2, "name1", 3, "name2", "name3", 4))
				}
			}

			var pagination api.Pagination
			repository := theory.NewRepository(logger, db)
			chords, _, err := repository.ListKeyChords(context.Background(), 1, pagination)
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

func TestTheoryRepository_ListKeyModes(t *testing.T) {
	type testCase struct {
		Title             string
		ListKeyModesError error
	}

	testCases := []testCase{
		{
			Title: "ReturnsKeysWhenSucceeded",
		},
		{
			Title:             "ReturnsErrorWhenFailed",
			ListKeyModesError: sql.ErrConnDone,
		},
	}

	listKeyModesQuery := `
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

	listKeyModesColumns := []string{
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

			if tc.ListKeyModesError != nil {
				sqlMock.ExpectQuery(listKeyModesQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(tc.ListKeyModesError)
			} else {
				sqlMock.ExpectQuery(listKeyModesQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows(listKeyModesColumns).
						AddRow(1, 2, "name1", 3, "name2", "name3", 4, true, 5.0, 6.0))
			}

			repository := theory.NewRepository(logger, db)
			keys, err := repository.ListKeyModes(context.Background(), 1)
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

func TestTheoryRepository_ListKeyPitches(t *testing.T) {
	type testCase struct {
		Title               string
		ListKeyPitchesError error
	}

	testCases := []testCase{
		{
			Title: "ReturnsPitchesWhenSucceeded",
		},
		{
			Title:               "ReturnsErrorWhenFailed",
			ListKeyPitchesError: sql.ErrConnDone,
		},
	}

	listKeyPitchesQuery := `
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

	listKeyPitchesColumns := []string{
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

			if tc.ListKeyPitchesError != nil {
				sqlMock.ExpectQuery(listKeyPitchesQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(tc.ListKeyPitchesError)
			} else {
				sqlMock.ExpectQuery(listKeyPitchesQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows(listKeyPitchesColumns).
						AddRow(1, "name", 2, 3.0))
			}

			repository := theory.NewRepository(logger, db)
			pitches, err := repository.ListKeyPitches(context.Background(), 1)
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

func TestTheoryRepository_ListKeys(t *testing.T) {
	type testCase struct {
		Title          string
		CountKeysError error
		ListKeysError  error
	}

	testCases := []testCase{
		{
			Title: "ReturnsKeysWhenSucceeded",
		},
		{
			Title:          "ReturnsErrorWhenCountKeysFailed",
			CountKeysError: sql.ErrConnDone,
		},
		{
			Title:         "ReturnsErrorWhenListKeysFailed",
			ListKeysError: sql.ErrConnDone,
		},
	}

	countKeysQuery := `
		SELECT
			COUNT(1)
		FROM keys;`

	countKeysColumns := []string{"count"}

	listKeysQuery := `
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

	listKeysColumns := []string{
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

			if tc.CountKeysError != nil {
				sqlMock.ExpectQuery(countKeysQuery).
					WillReturnError(tc.CountKeysError)
			} else {
				sqlMock.ExpectQuery(countKeysQuery).
					WillReturnRows(sqlmock.NewRows(countKeysColumns).
						AddRow(1))

				if tc.ListKeysError != nil {
					sqlMock.ExpectQuery(listKeysQuery).
						WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnError(tc.ListKeysError)
				} else {
					sqlMock.ExpectQuery(listKeysQuery).
						WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnRows(sqlmock.NewRows(listKeysColumns).
							AddRow(1, 2, "name1", 3, "name2", "name3", 4, true, 5.0, 6.0))
				}
			}

			var pagination api.Pagination
			repository := theory.NewRepository(logger, db)
			keys, _, err := repository.ListKeys(context.Background(), pagination)
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

func TestTheoryRepository_ListPitches(t *testing.T) {
	type testCase struct {
		Title            string
		ListPitchesError error
	}

	testCases := []testCase{
		{
			Title: "ReturnsPitchesWhenSucceeded",
		},
		{
			Title:            "ReturnsErrorWhenFailed",
			ListPitchesError: sql.ErrConnDone,
		},
	}

	listPitchesQuery := `
		SELECT
			id,
			name,
			number,
			frequency
		FROM pitches
		ORDER BY id;`

	listPitchesColumns := []string{
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

			if tc.ListPitchesError != nil {
				sqlMock.ExpectQuery(listPitchesQuery).
					WillReturnError(tc.ListPitchesError)
			} else {
				sqlMock.ExpectQuery(listPitchesQuery).
					WillReturnRows(sqlmock.NewRows(listPitchesColumns).
						AddRow(1, "name", 2, 3.0))
			}

			repository := theory.NewRepository(logger, db)
			pitches, err := repository.ListPitches(context.Background())
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

func TestTheoryRepository_ListScaleKeys(t *testing.T) {
	type testCase struct {
		Title              string
		ListScaleKeysError error
	}

	testCases := []testCase{
		{
			Title: "ReturnsKeysWhenSucceeded",
		},
		{
			Title:              "ReturnsErrorWhenListScaleKeysFailed",
			ListScaleKeysError: sql.ErrConnDone,
		},
	}

	listScaleKeysQuery := `
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

	listScaleKeysColumns := []string{
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

			if tc.ListScaleKeysError != nil {
				sqlMock.ExpectQuery(listScaleKeysQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(tc.ListScaleKeysError)
			} else {
				sqlMock.ExpectQuery(listScaleKeysQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows(listScaleKeysColumns).
						AddRow(1, 2, "name1", 3, "name2", "name3", 4, true, 5.0, 6.0))
			}

			repository := theory.NewRepository(logger, db)
			keys, err := repository.ListScaleKeys(context.Background(), 1)
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

func TestTheoryRepository_ListScales(t *testing.T) {
	type testCase struct {
		Title            string
		CountScalesError error
		ListScalesError  error
	}

	testCases := []testCase{
		{
			Title: "ReturnsScalesWhenSucceeded",
		},
		{
			Title:            "ReturnsErrorWhenCountScalesFailed",
			CountScalesError: sql.ErrConnDone,
		},
		{
			Title:           "ReturnsErrorWhenListScalesFailed",
			ListScalesError: sql.ErrConnDone,
		},
	}

	countScalesQuery := `
		SELECT
			COUNT(1)
		FROM scales;`

	countScalesColumns := []string{"count"}

	listScalesQuery := `
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

	listScalesColumns := []string{
		"id",
		"name",
		"cardinality",
		"number",
		"perfection",
		"imperfection",
		"pitch_class",
		"interval_pattern",
		"rotational_symmetric",
		"rotational_symmetry_level",
		"palindromic",
		"reflectional_symmetric",
		"reflectional_symmetry_axes",
		"balanced",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			db, sqlMock, err := mockDatabase()
			require.NoError(t, err)

			if tc.CountScalesError != nil {
				sqlMock.ExpectQuery(countScalesQuery).
					WillReturnError(tc.CountScalesError)
			} else {
				sqlMock.ExpectQuery(countScalesQuery).
					WillReturnRows(sqlmock.NewRows(countScalesColumns).
						AddRow(1))

				if tc.ListScalesError != nil {
					sqlMock.ExpectQuery(listScalesQuery).
						WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnError(tc.ListScalesError)
				} else {
					sqlMock.ExpectQuery(listScalesQuery).
						WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnRows(sqlmock.NewRows(listScalesColumns).
							AddRow(1, "name", 2, 3, 4, 5, []byte("[4,5]"), []byte("[6,7]"), true, 8, true, true, []byte("[9,10]"), true))
				}
			}

			var pagination api.Pagination
			repository := theory.NewRepository(logger, db)
			scales, _, err := repository.ListScales(context.Background(), pagination)
			if strings.HasPrefix(tc.Title, "ReturnsError") {
				require.Error(t, err)
				require.Empty(t, scales)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, scales)
			}
		})
	}
}

func mockDatabase() (*sqlx.DB, sqlmock.Sqlmock, error) {
	db, sqlMock, err := sqlmock.New(
		sqlmock.MonitorPingsOption(true),
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual),
	)
	if err != nil {
		return nil, nil, err
	}

	dbx := sqlx.NewDb(db, "postgres")
	return dbx, sqlMock, err
}

// MockRepositoryReturnValues stores return value of mocked theory.Repository
type MockRepositoryReturnValues struct {
	ListPitches []interface{}

	ListChords       []interface{}
	ListChordPitches []interface{}
	ListChordKeys    []interface{}
	GetChord         []interface{}
	GetChordQuality  []interface{}

	ListScales    []interface{}
	ListScaleKeys []interface{}
	GetScale      []interface{}

	ListKeys       []interface{}
	ListKeyModes   []interface{}
	ListKeyChords  []interface{}
	ListKeyPitches []interface{}
	GetKey         []interface{}
}

// MockRepository mocks theory.Repository
type MockRepository struct {
	mock.Mock
}

// ListPitches mocks theory.Repository#ListPitches
func (m *MockRepository) ListPitches(ctx context.Context) ([]theory.DetailedPitch, error) {
	args := m.Called(ctx)

	var entries []theory.DetailedPitch
	if v, ok := args.Get(0).([]theory.DetailedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListChords mocks theory.Repository#ListChords
func (m *MockRepository) ListChords(ctx context.Context, pagination api.Pagination) ([]theory.DetailedChord, *api.Pagination, error) {
	args := m.Called(ctx, pagination)

	var entries []theory.DetailedChord
	if v, ok := args.Get(0).([]theory.DetailedChord); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(0).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// ListChordPitches mocks theory.Repository#ListChordPitches
func (m *MockRepository) ListChordPitches(ctx context.Context, chordID int64) ([]theory.DetailedPitch, error) {
	args := m.Called(ctx, chordID)

	var entries []theory.DetailedPitch
	if v, ok := args.Get(0).([]theory.DetailedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListChordKeys mocks theory.Repository#ListChordKeys
func (m *MockRepository) ListChordKeys(ctx context.Context, chordID int64, pagination api.Pagination) ([]theory.DetailedKey, *api.Pagination, error) {
	args := m.Called(ctx, chordID, pagination)

	var entries []theory.DetailedKey
	if v, ok := args.Get(0).([]theory.DetailedKey); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(0).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// GetChord mocks theory.Repository#GetChord
func (m *MockRepository) GetChord(ctx context.Context, chordID int64) (*theory.DetailedChord, error) {
	args := m.Called(ctx, chordID)

	var entry *theory.DetailedChord
	if v, ok := args.Get(0).(*theory.DetailedChord); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// GetChordQuality mocks theory.Repository#GetChordQuality
func (m *MockRepository) GetChordQuality(ctx context.Context, chordID int64) (*theory.DetailedChordQuality, error) {
	args := m.Called(ctx, chordID)

	var entry *theory.DetailedChordQuality
	if v, ok := args.Get(0).(*theory.DetailedChordQuality); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// ListScales mocks theory.Repository#ListScales
func (m *MockRepository) ListScales(ctx context.Context, pagination api.Pagination) ([]theory.DetailedScale, *api.Pagination, error) {
	args := m.Called(ctx, pagination)

	var entries []theory.DetailedScale
	if v, ok := args.Get(0).([]theory.DetailedScale); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(0).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// ListScaleKeys mocks theory.Repository#ListScaleKeys
func (m *MockRepository) ListScaleKeys(ctx context.Context, scaleID int64) ([]theory.DetailedKey, error) {
	args := m.Called(ctx, scaleID)

	var entries []theory.DetailedKey
	if v, ok := args.Get(0).([]theory.DetailedKey); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// GetScale mocks theory.Repository#GetScale
func (m *MockRepository) GetScale(ctx context.Context, scaleID int64) (*theory.DetailedScale, error) {
	args := m.Called(ctx, scaleID)

	var entry *theory.DetailedScale
	if v, ok := args.Get(0).(*theory.DetailedScale); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// ListKeys mocks theory.Repository#ListKeys
func (m *MockRepository) ListKeys(ctx context.Context, pagination api.Pagination) ([]theory.DetailedKey, *api.Pagination, error) {
	args := m.Called(ctx, pagination)

	var entries []theory.DetailedKey
	if v, ok := args.Get(0).([]theory.DetailedKey); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(0).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// ListKeyModes mocks theory.Repository#ListKeyModes
func (m *MockRepository) ListKeyModes(ctx context.Context, keyID int64) ([]theory.DetailedKey, error) {
	args := m.Called(ctx, keyID)

	var entries []theory.DetailedKey
	if v, ok := args.Get(0).([]theory.DetailedKey); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListKeyChords mocks theory.Repository#ListKeyChords
func (m *MockRepository) ListKeyChords(ctx context.Context, keyID int64, pagination api.Pagination) ([]theory.DetailedChord, *api.Pagination, error) {
	args := m.Called(ctx, keyID, pagination)

	var entries []theory.DetailedChord
	if v, ok := args.Get(0).([]theory.DetailedChord); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(0).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// ListKeyPitches mocks theory.Repository#ListKeyPitches
func (m *MockRepository) ListKeyPitches(ctx context.Context, keyID int64) ([]theory.DetailedPitch, error) {
	args := m.Called(ctx, keyID)

	var entries []theory.DetailedPitch
	if v, ok := args.Get(0).([]theory.DetailedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// GetKey mocks theory.Repository#GetKey
func (m *MockRepository) GetKey(ctx context.Context, keyID int64) (*theory.DetailedKey, error) {
	args := m.Called(ctx, keyID)

	var entry *theory.DetailedKey
	if v, ok := args.Get(0).(*theory.DetailedKey); ok {
		entry = v
	}

	return entry, args.Error(1)
}
