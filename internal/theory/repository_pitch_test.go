package theory_test

import (
	"context"
	"database/sql"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
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
			number,
			frequency
		FROM pitches
		WHERE id = $1;`

	getPitchColumns := []string{
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

			if tc.GetPitchError != nil {
				sqlMock.ExpectQuery(getPitchQuery).
					WillReturnError(tc.GetPitchError)
			} else {
				sqlMock.ExpectQuery(getPitchQuery).
					WillReturnRows(sqlmock.NewRows(getPitchColumns).
						AddRow(1, "name", 2, 3.0))
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
