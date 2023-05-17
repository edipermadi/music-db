package theory_test

import (
	"context"
	"database/sql"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
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
