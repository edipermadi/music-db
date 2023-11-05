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
			name
		FROM pitches
		ORDER BY id;`

	listPitchesColumns := []string{
		"id",
		"name",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger := mock.Logger()

			db, sqlMock, err := mockDatabase()
			require.NoError(t, err)

			if tc.ListPitchesError != nil {
				sqlMock.ExpectQuery(listPitchesQuery).
					WillReturnError(tc.ListPitchesError)
			} else {
				sqlMock.ExpectQuery(listPitchesQuery).
					WillReturnRows(sqlmock.NewRows(listPitchesColumns).
						AddRow(1, "name"))
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

func TestTheoryRepository_ListPitchChords(t *testing.T) {
	type testCase struct {
		Title                 string
		CountPitchChordsError error
		ListPitchChordsError  error
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
		},
		{
			Title:                "ReturnsErrorWhenFailed",
			ListPitchChordsError: sql.ErrConnDone,
		},
	}

	countPitchChordsQuery := `
		SELECT 
			COUNT(DISTINCT c.id)
		FROM chord_pitches cp
			JOIN chords c ON cp.chord_id = c.id
			JOIN chord_qualities cq ON c.chord_quality_id = cq.id
		WHERE cp.pitch_id = $1
		GROUP BY
		    cp.pitch_id;`

	listPitchChordsQuery := `
		SELECT DISTINCT
			c.id, 
			c.name
		FROM chord_pitches cp
			JOIN chords c ON cp.chord_id = c.id
			JOIN chord_qualities cq ON c.chord_quality_id = cq.id
		WHERE
		    cp.pitch_id = $1
		ORDER
			BY c.id
		OFFSET $2
		LIMIT $3;`

	listPitchChordsColumns := []string{
		"id",
		"name",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger := mock.Logger()

			db, sqlMock, err := mockDatabase()
			require.NoError(t, err)

			if tc.CountPitchChordsError != nil {
				sqlMock.ExpectQuery(countPitchChordsQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(tc.ListPitchChordsError)
			} else {
				sqlMock.ExpectQuery(countPitchChordsQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				if tc.ListPitchChordsError != nil {
					sqlMock.ExpectQuery(listPitchChordsQuery).
						WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnError(tc.ListPitchChordsError)
				} else {
					sqlMock.ExpectQuery(listPitchChordsQuery).
						WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnRows(sqlmock.NewRows(listPitchChordsColumns).
							AddRow(1, "name"))
				}
			}

			repository := theory.NewRepository(logger, db)
			chords, _, err := repository.ListPitchChords(context.Background(), 1, theory.ChordFilter{}, api.Pagination{})
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

func TestTheoryRepository_ListPitchKeys(t *testing.T) {
	type testCase struct {
		Title               string
		CountPitchKeysError error
		ListPitchKeysError  error
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
		},
		{
			Title:               "ReturnsErrorWhenCountingFailed",
			CountPitchKeysError: sql.ErrConnDone,
		},
		{
			Title:              "ReturnsErrorWhenListingFailed",
			ListPitchKeysError: sql.ErrConnDone,
		},
	}

	countPitchKeysQuery := `
		SELECT 
			COUNT(DISTINCT k.id)
		FROM key_pitches kp
			JOIN keys k ON kp.key_id = k.id
			JOIN scales s ON k.scale_id = s.id
		WHERE kp.pitch_id = $1
		GROUP BY
		    kp.pitch_id;`

	listPitchKeysQuery := `
		SELECT DISTINCT
			k.id, 
			k.name
		FROM key_pitches kp
			JOIN keys k ON kp.key_id = k.id
			JOIN scales s ON k.scale_id = s.id
		WHERE
		    kp.pitch_id = $1
		ORDER
			BY k.id
		OFFSET $2
		LIMIT $3;`

	listPitchKeysColumns := []string{
		"id",
		"name",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger := mock.Logger()

			db, sqlMock, err := mockDatabase()
			require.NoError(t, err)

			if tc.CountPitchKeysError != nil {
				sqlMock.ExpectQuery(countPitchKeysQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(tc.CountPitchKeysError)
			} else {
				sqlMock.ExpectQuery(countPitchKeysQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				if tc.ListPitchKeysError != nil {
					sqlMock.ExpectQuery(listPitchKeysQuery).
						WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnError(tc.ListPitchKeysError)
				} else {
					sqlMock.ExpectQuery(listPitchKeysQuery).
						WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnRows(sqlmock.NewRows(listPitchKeysColumns).
							AddRow(1, "name"))
				}
			}

			repository := theory.NewRepository(logger, db)
			chords, _, err := repository.ListPitchKeys(context.Background(), 1, theory.KeyFilter{}, api.Pagination{})
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

func TestTheoryRepository_ListPitchScales(t *testing.T) {
	type testCase struct {
		Title               string
		CountPitchKeysError error
		ListPitchKeysError  error
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
		},
		{
			Title:               "ReturnsErrorWhenCountingFailed",
			CountPitchKeysError: sql.ErrConnDone,
		},
		{
			Title:              "ReturnsErrorWhenListingFailed",
			ListPitchKeysError: sql.ErrConnDone,
		},
	}

	countQuery := `
		SELECT 
			COUNT(DISTINCT s.id)
		FROM key_pitches kp
			JOIN keys k ON kp.key_id = k.id
			JOIN scales s ON k.scale_id = s.id
		WHERE kp.pitch_id = $1
		GROUP BY
		    kp.pitch_id;`

	listQuery := `
		SELECT DISTINCT
			s.id, 
			s.name
		FROM key_pitches kp
			JOIN keys k ON kp.key_id = k.id
			JOIN scales s ON k.scale_id = s.id
		WHERE
		    kp.pitch_id = $1
		ORDER
			BY s.id
		OFFSET $2
		LIMIT $3;`

	listPitchKeysColumns := []string{
		"id",
		"name",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger := mock.Logger()

			db, sqlMock, err := mockDatabase()
			require.NoError(t, err)

			if tc.CountPitchKeysError != nil {
				sqlMock.ExpectQuery(countQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(tc.CountPitchKeysError)
			} else {
				sqlMock.ExpectQuery(countQuery).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				if tc.ListPitchKeysError != nil {
					sqlMock.ExpectQuery(listQuery).
						WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnError(tc.ListPitchKeysError)
				} else {
					sqlMock.ExpectQuery(listQuery).
						WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
						WillReturnRows(sqlmock.NewRows(listPitchKeysColumns).
							AddRow(1, "name"))
				}
			}

			repository := theory.NewRepository(logger, db)
			chords, _, err := repository.ListPitchScales(context.Background(), 1, theory.ScaleFilter{}, api.Pagination{})
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

func TestTheoryRepository_GetPitch(t *testing.T) {
	type testCase struct {
		Title         string
		GetPitchError error
	}

	testCases := []testCase{
		{
			Title: "ReturnsPitchWhenSucceeded",
		},
		{
			Title:         "ReturnsErrorWhenFailed",
			GetPitchError: sql.ErrConnDone,
		},
	}

	getPitchQuery := `
		SELECT
			id,
			name,
			zeitler_number,
			ring_number,
			frequency
		FROM pitches
		WHERE id = $1;`

	getPitchColumns := []string{
		"id",
		"name",
		"zeitler_number",
		"ring_number",
		"frequency",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger := mock.Logger()

			db, sqlMock, err := mockDatabase()
			require.NoError(t, err)

			if tc.GetPitchError != nil {
				sqlMock.ExpectQuery(getPitchQuery).
					WillReturnError(tc.GetPitchError)
			} else {
				sqlMock.ExpectQuery(getPitchQuery).
					WillReturnRows(sqlmock.NewRows(getPitchColumns).
						AddRow(1, "name", 2, 3, 4.0))
			}

			repository := theory.NewRepository(logger, db)
			pitch, err := repository.GetPitch(context.Background(), 1)
			if strings.HasPrefix(tc.Title, "ReturnsError") {
				require.Error(t, err)
				require.Empty(t, pitch)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, pitch)
			}
		})
	}
}
