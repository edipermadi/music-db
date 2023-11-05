package theory_test

import (
	"context"
	"database/sql"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/edipermadi/music-db/mock"
	"github.com/stretchr/testify/require"
)

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
		FROM scales s
			JOIN keys k ON s.id = k.scale_id
		WHERE
			TRUE;`

	countScalesColumns := []string{"count"}

	listScalesQuery := `
		SELECT
			s.id,
			s.name
		FROM scales s
			JOIN keys k ON s.id = k.scale_id
		WHERE TRUE
		ORDER BY
			s.id
		OFFSET $1
		LIMIT  $2;`

	listScalesColumns := []string{
		"id",
		"name",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger := mock.Logger()

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
							AddRow(1, "name"))
				}
			}

			var pagination api.Pagination
			var filter theory.ScaleFilter
			repository := theory.NewRepository(logger, db)
			scales, _, err := repository.ListScales(context.Background(), filter, pagination)
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

func TestTheoryRepository_GetScale(t *testing.T) {
	type testCase struct {
		Title         string
		GetScaleError error
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
	}

	getScaleQuery := `
		SELECT
			id,
			name,
			cardinality,
			zeitler_number,
			ring_number,
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
		"zeitler_number",
		"ring_number",
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
			logger := mock.Logger()

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
						AddRow(1, "name", 2, 3, 4, 5, 6, []byte("[7,8]"), []byte("[9,10]"), true, 11, true, true, []byte("[12,13]"), true))
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
			k.name
		FROM keys k
			JOIN scales s ON k.scale_id = s.id
			JOIN pitches p ON k.tonic_id = p.id
		WHERE
			scale_id = $1
		ORDER BY
			k.id;`

	listScaleKeysColumns := []string{
		"id",
		"name",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger := mock.Logger()

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
						AddRow(1, "name"))
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

func TestTheoryRepository_ListScalePitches(t *testing.T) {
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

	listScaleKeysColumns := []string{
		"id",
		"name",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger := mock.Logger()

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
						AddRow(1, "name"))
			}

			repository := theory.NewRepository(logger, db)
			keys, err := repository.ListScalePitches(context.Background(), 1)
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

func TestTheoryRepository_ListScaleChords(t *testing.T) {
	type testCase struct {
		Title      string
		CountError error
		ListError  error
	}

	testCases := []testCase{
		{
			Title: "ReturnsScalesWhenSucceeded",
		},
		{
			Title:      "ReturnsErrorWhenCountScalesFailed",
			CountError: sql.ErrConnDone,
		},
		{
			Title:     "ReturnsErrorWhenListScalesFailed",
			ListError: sql.ErrConnDone,
		},
	}

	countQuery := `
		SELECT
			COUNT(c.id)
		FROM key_pitch_chords kpc
			JOIN chords c ON kpc.chord_id = c.id
			JOIN chord_qualities cq ON c.chord_quality_id = cq.id
			JOIN pitches p ON c.root_id = p.id
			JOIN keys k ON kpc.key_id = k.id
		WHERE
			k.scale_id = $1;`

	listQuery := `
		SELECT
			c.id,
			c.name
		FROM key_pitch_chords kpc
			JOIN chords c ON kpc.chord_id = c.id
			JOIN chord_qualities cq ON c.chord_quality_id = cq.id
			JOIN pitches p ON c.root_id = p.id
			JOIN keys k ON kpc.key_id = k.id
		WHERE
		    k.scale_id = $1
		ORDER BY
			c.id
		OFFSET $2
		LIMIT  $3;`

	listColumns := []string{
		"id",
		"name",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger := mock.Logger()

			db, sqlMock, err := mockDatabase()
			require.NoError(t, err)

			if tc.CountError != nil {
				sqlMock.ExpectQuery(countQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(tc.CountError)
			} else {
				sqlMock.ExpectQuery(countQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).
						AddRow(1))

				if tc.ListError != nil {
					sqlMock.ExpectQuery(listQuery).
						WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnError(tc.ListError)
				} else {
					sqlMock.ExpectQuery(listQuery).
						WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnRows(sqlmock.NewRows(listColumns).
							AddRow(1, "name"))
				}
			}

			repository := theory.NewRepository(logger, db)
			scales, _, err := repository.ListScaleChords(context.Background(), 1, theory.ChordFilter{}, api.Pagination{})
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
