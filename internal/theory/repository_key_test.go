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
	"github.com/edipermadi/music-db/mock"
	"github.com/stretchr/testify/require"
)

func TestTheoryRepository_ListKeys(t *testing.T) {
	type testCase struct {
		Title              string
		GivenFilter        theory.KeyFilter
		ExpectedCountQuery string
		ExpectedListQuery  string
		ExpectedCountArgs  []driver.Value
		ExpectedListArgs   []driver.Value
		CountKeysError     error
		ListKeysError      error
	}

	valueInt := 1
	valueBool := true
	testCases := []testCase{
		{
			Title: "ReturnsKeysWhenSucceededWithoutFilter",
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT k.id)
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
				WHERE TRUE;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					TRUE
				ORDER BY
					k.id
				OFFSET $1
				LIMIT  $2;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithScaleIDFilter",
			GivenFilter: theory.KeyFilter{ScaleID: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT k.id)
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
				WHERE
					k.scale_id = $1;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.scale_id = $1
				ORDER BY
					k.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithTonicIDFilter",
			GivenFilter: theory.KeyFilter{TonicID: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT k.id)
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
				WHERE
					k.tonic_id = $1;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.tonic_id = $1
				ORDER BY
					k.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithZeitlerNumberFilter",
			GivenFilter: theory.KeyFilter{ZeitlerNumber: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT k.id)
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
				WHERE
					k.zeitler_number = $1;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.zeitler_number = $1
				ORDER BY
					k.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithRingNumberFilter",
			GivenFilter: theory.KeyFilter{RingNumber: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT k.id)
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
				WHERE
					k.ring_number = $1;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.ring_number = $1
				ORDER BY
					k.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithPerfectionFilter",
			GivenFilter: theory.KeyFilter{Perfection: &valueInt},
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT k.id)
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
				WHERE
					s.perfection = $1;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					s.perfection = $1
				ORDER BY
					k.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithImperfectionFilter",
			GivenFilter: theory.KeyFilter{Imperfection: &valueInt},
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT k.id)
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
				WHERE
					s.imperfection = $1;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					s.imperfection = $1
				ORDER BY
					k.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithBalancedFilter",
			GivenFilter: theory.KeyFilter{Balanced: &valueBool},
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT k.id)
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
				WHERE
					k.balanced = $1;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.balanced = $1
				ORDER BY
					k.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithRotationalSymmetricFilter",
			GivenFilter: theory.KeyFilter{RotationalSymmetric: &valueBool},
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT k.id)
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
				WHERE
					s.rotational_symmetric = $1;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					s.rotational_symmetric = $1
				ORDER BY
					k.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithRotationalSymmetryLevelFilter",
			GivenFilter: theory.KeyFilter{RotationalSymmetryLevel: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT k.id)
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
				WHERE
					s.rotational_symmetry_level = $1;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					s.rotational_symmetry_level = $1
				ORDER BY
					k.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithReflectionalSymmetricFilter",
			GivenFilter: theory.KeyFilter{ReflectionalSymmetric: &valueBool},
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT k.id)
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
				WHERE
					s.reflectional_symmetric = $1;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					s.reflectional_symmetric = $1
				ORDER BY
					k.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithPalindromicFilter",
			GivenFilter: theory.KeyFilter{Palindromic: &valueBool},
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT k.id)
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
				WHERE
					s.palindromic = $1;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					s.palindromic = $1
				ORDER BY
					k.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithCardinalityFilter",
			GivenFilter: theory.KeyFilter{Cardinality: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT k.id)
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
				WHERE
					s.cardinality = $1;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					s.cardinality = $1
				ORDER BY
					k.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title: "ReturnsErrorWhenCountKeysFailed",
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT k.id)
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
				WHERE TRUE;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					TRUE
				ORDER BY
					k.id
				OFFSET $1
				LIMIT  $2;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			CountKeysError:   sql.ErrConnDone,
		},
		{
			Title: "ReturnsErrorWhenListKeysFailed",
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT k.id)
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
				WHERE TRUE;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					TRUE
				ORDER BY
					k.id
				OFFSET $1
				LIMIT  $2;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ListKeysError:    sql.ErrConnDone,
		},
	}

	countKeysColumns := []string{"count"}

	listKeysColumns := []string{
		"id",
		"name",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger := mock.Logger()

			db, sqlMock, err := mockDatabase()
			require.NoError(t, err)

			if tc.CountKeysError != nil {
				sqlMock.ExpectQuery(tc.ExpectedCountQuery).
					WithArgs(tc.ExpectedCountArgs...).
					WillReturnError(tc.CountKeysError)
			} else {
				sqlMock.ExpectQuery(tc.ExpectedCountQuery).
					WithArgs(tc.ExpectedCountArgs...).
					WillReturnRows(sqlmock.NewRows(countKeysColumns).
						AddRow(1))

				if tc.ListKeysError != nil {
					sqlMock.ExpectQuery(tc.ExpectedListQuery).
						WithArgs(tc.ExpectedListArgs...).
						WillReturnError(tc.ListKeysError)
				} else {
					sqlMock.ExpectQuery(tc.ExpectedListQuery).
						WithArgs(tc.ExpectedListArgs...).
						WillReturnRows(sqlmock.NewRows(listKeysColumns).
							AddRow(1, "name"))
				}
			}

			repository := theory.NewRepository(logger, db)
			keys, _, err := repository.ListKeys(context.Background(), tc.GivenFilter, api.Pagination{})
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
		Title       string
		GetKeyError error
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
	}

	getKeyQuery := `
		SELECT
			k.id,
			s.id   AS "scale.id",
			s.name AS "scale.name",
			p.id   AS "tonic.id",
			p.name AS "tonic.name",
			k.name,
			k.zeitler_number,
			k.ring_number,
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
		"zeitler_number",
		"ring_number",
		"balanced",
		"center_x",
		"center_y",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger := mock.Logger()

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
						AddRow(1, 2, "name1", 3, "name2", "name3", 4, 5, true, 6.0, 7.0))
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

func TestTheoryRepository_ListKeyModes(t *testing.T) {
	type testCase struct {
		Title             string
		GivenFilter       theory.KeyFilter
		ExpectedListQuery string
		ExpectedListArgs  []driver.Value
		ListKeyModesError error
	}

	intValue := 1
	boolValue := true
	testCases := []testCase{
		{
			Title: "ReturnsKeysWhenSucceededWithoutFilter",
			ExpectedListQuery: `
				WITH numbers AS(
					SELECT
						zeitler_number
					FROM keys
					WHERE id = $1
				)
				SELECT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.zeitler_number IN (SELECT zeitler_number FROM numbers)
				ORDER BY
					k.id;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithScaleIDFilter",
			GivenFilter: theory.KeyFilter{ScaleID: 1},
			ExpectedListQuery: `
				WITH numbers AS(
					SELECT
						zeitler_number
					FROM keys
					WHERE id = $1
				)
				SELECT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.zeitler_number IN (SELECT zeitler_number FROM numbers) AND
					k.scale_id = $2
				ORDER BY
					k.id;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithTonicIDFilter",
			GivenFilter: theory.KeyFilter{TonicID: 1},
			ExpectedListQuery: `
				WITH numbers AS(
					SELECT
						zeitler_number
					FROM keys
					WHERE id = $1
				)
				SELECT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.zeitler_number IN (SELECT zeitler_number FROM numbers) AND
					k.tonic_id = $2
				ORDER BY
					k.id;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithZeitlerNumberFilter",
			GivenFilter: theory.KeyFilter{ZeitlerNumber: 1},
			ExpectedListQuery: `
				WITH numbers AS(
					SELECT
						zeitler_number
					FROM keys
					WHERE id = $1
				)
				SELECT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.zeitler_number IN (SELECT zeitler_number FROM numbers) AND
					k.zeitler_number = $2
				ORDER BY
					k.id;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithRingNumberFilter",
			GivenFilter: theory.KeyFilter{RingNumber: 1},
			ExpectedListQuery: `
				WITH numbers AS(
					SELECT
						zeitler_number
					FROM keys
					WHERE id = $1
				)
				SELECT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.zeitler_number IN (SELECT zeitler_number FROM numbers) AND
					k.ring_number = $2
				ORDER BY
					k.id;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithPerfectionFilter",
			GivenFilter: theory.KeyFilter{Perfection: &intValue},
			ExpectedListQuery: `
				WITH numbers AS(
					SELECT
						zeitler_number
					FROM keys
					WHERE id = $1
				)
				SELECT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.zeitler_number IN (SELECT zeitler_number FROM numbers) AND
					s.perfection = $2
				ORDER BY
					k.id;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithImperfectionFilter",
			GivenFilter: theory.KeyFilter{Imperfection: &intValue},
			ExpectedListQuery: `
				WITH numbers AS(
					SELECT
						zeitler_number
					FROM keys
					WHERE id = $1
				)
				SELECT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.zeitler_number IN (SELECT zeitler_number FROM numbers) AND
					s.imperfection = $2
				ORDER BY
					k.id;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithBalancedFilter",
			GivenFilter: theory.KeyFilter{Balanced: &boolValue},
			ExpectedListQuery: `
				WITH numbers AS(
					SELECT
						zeitler_number
					FROM keys
					WHERE id = $1
				)
				SELECT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.zeitler_number IN (SELECT zeitler_number FROM numbers) AND
					k.balanced = $2
				ORDER BY
					k.id;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithRotationalSymmetricFilter",
			GivenFilter: theory.KeyFilter{RotationalSymmetric: &boolValue},
			ExpectedListQuery: `
				WITH numbers AS(
					SELECT
						zeitler_number
					FROM keys
					WHERE id = $1
				)
				SELECT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.zeitler_number IN (SELECT zeitler_number FROM numbers) AND
					s.rotational_symmetric = $2
				ORDER BY
					k.id;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithRotationalSymmetryLevelFilter",
			GivenFilter: theory.KeyFilter{RotationalSymmetryLevel: 1},
			ExpectedListQuery: `
				WITH numbers AS(
					SELECT
						zeitler_number
					FROM keys
					WHERE id = $1
				)
				SELECT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.zeitler_number IN (SELECT zeitler_number FROM numbers) AND
					s.rotational_symmetry_level = $2
				ORDER BY
					k.id;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithReflectiveSymmetricFilter",
			GivenFilter: theory.KeyFilter{ReflectionalSymmetric: &boolValue},
			ExpectedListQuery: `
				WITH numbers AS(
					SELECT
						zeitler_number
					FROM keys
					WHERE id = $1
				)
				SELECT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.zeitler_number IN (SELECT zeitler_number FROM numbers) AND
					s.reflectional_symmetric = $2
				ORDER BY
					k.id;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithPalindromicFilter",
			GivenFilter: theory.KeyFilter{Palindromic: &boolValue},
			ExpectedListQuery: `
				WITH numbers AS(
					SELECT
						zeitler_number
					FROM keys
					WHERE id = $1
				)
				SELECT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.zeitler_number IN (SELECT zeitler_number FROM numbers) AND
					s.palindromic = $2
				ORDER BY
					k.id;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsKeysWhenSucceededWithCardinalityFilter",
			GivenFilter: theory.KeyFilter{Cardinality: 1},
			ExpectedListQuery: `
				WITH numbers AS(
					SELECT
						zeitler_number
					FROM keys
					WHERE id = $1
				)
				SELECT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.zeitler_number IN (SELECT zeitler_number FROM numbers) AND
					s.cardinality = $2
				ORDER BY
					k.id;`,
			ExpectedListArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			ExpectedListQuery: `
				WITH numbers AS(
					SELECT
						zeitler_number
					FROM keys
					WHERE id = $1
				)
				SELECT
					k.id,
					k.name
				FROM keys k
					JOIN scales s ON k.scale_id = s.id
					JOIN pitches p ON k.tonic_id = p.id
				WHERE
					k.zeitler_number IN (SELECT zeitler_number FROM numbers)
				ORDER BY
					k.id;`,
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg()},
			ListKeyModesError: sql.ErrConnDone,
		},
	}

	listKeyModesColumns := []string{
		"id",
		"name",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger := mock.Logger()

			db, sqlMock, err := mockDatabase()
			require.NoError(t, err)

			if tc.ListKeyModesError != nil {
				sqlMock.ExpectQuery(tc.ExpectedListQuery).
					WithArgs(tc.ExpectedListArgs...).
					WillReturnError(tc.ListKeyModesError)
			} else {
				sqlMock.ExpectQuery(tc.ExpectedListQuery).
					WithArgs(tc.ExpectedListArgs...).
					WillReturnRows(sqlmock.NewRows(listKeyModesColumns).
						AddRow(1, "name"))
			}

			repository := theory.NewRepository(logger, db)
			keys, err := repository.ListKeyModes(context.Background(), 1, tc.GivenFilter)
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

func TestTheoryRepository_ListKeyChords(t *testing.T) {
	type testCase struct {
		Title               string
		GivenFilter         theory.ChordFilter
		ExpectedCountQuery  string
		ExpectedListQuery   string
		ExpectedCountArgs   []driver.Value
		ExpectedListArgs    []driver.Value
		CountKeyChordsError error
		ListKeyChordsError  error
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordsWhenSucceededWithoutFilter",
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT c.id)
				FROM key_pitch_chords kpc
					JOIN chords c ON kpc.chord_id = c.id
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
				WHERE kpc.key_id = $1;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					c.id,
					c.name
				FROM key_pitch_chords kpc 
					JOIN chords c ON kpc.chord_id = c.id
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
					JOIN pitches p ON c.root_id = p.id
				WHERE
					kpc.key_id = $1
				ORDER BY
					c.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordsWhenSucceededWithChordQualityIDFilter",
			GivenFilter: theory.ChordFilter{ChordQualityID: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT c.id)
				FROM key_pitch_chords kpc
					JOIN chords c ON kpc.chord_id = c.id
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
				WHERE
					kpc.key_id         = $1 AND
					c.chord_quality_id = $2;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					c.id,
					c.name
				FROM key_pitch_chords kpc 
					JOIN chords c ON kpc.chord_id = c.id
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
					JOIN pitches p ON c.root_id = p.id
				WHERE
					kpc.key_id         = $1 AND
					c.chord_quality_id = $2
				ORDER BY
					c.id
				OFFSET $3
				LIMIT  $4;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordsWhenSucceededWithRootIDFilter",
			GivenFilter: theory.ChordFilter{RootID: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT c.id)
				FROM key_pitch_chords kpc
					JOIN chords c ON kpc.chord_id = c.id
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
				WHERE
					kpc.key_id = $1 AND
					c.root_id  = $2;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					c.id,
					c.name
				FROM key_pitch_chords kpc 
					JOIN chords c ON kpc.chord_id = c.id
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
					JOIN pitches p ON c.root_id = p.id
				WHERE
					kpc.key_id = $1 AND
					c.root_id  = $2
				ORDER BY
					c.id
				OFFSET $3
				LIMIT  $4;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordsWhenSucceededWithZeitlerNumberFilter",
			GivenFilter: theory.ChordFilter{ZeitlerNumber: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT c.id)
				FROM key_pitch_chords kpc
					JOIN chords c ON kpc.chord_id = c.id
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
				WHERE
					kpc.key_id       = $1 AND
					c.zeitler_number = $2;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					c.id,
					c.name
				FROM key_pitch_chords kpc 
					JOIN chords c ON kpc.chord_id = c.id
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
					JOIN pitches p ON c.root_id = p.id
				WHERE
					kpc.key_id       = $1 AND
					c.zeitler_number = $2
				ORDER BY
					c.id
				OFFSET $3
				LIMIT  $4;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordsWhenSucceededWithRingNumberFilter",
			GivenFilter: theory.ChordFilter{RingNumber: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT c.id)
				FROM key_pitch_chords kpc
					JOIN chords c ON kpc.chord_id = c.id
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
				WHERE
					kpc.key_id    = $1 AND
					c.ring_number = $2;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					c.id,
					c.name
				FROM key_pitch_chords kpc 
					JOIN chords c ON kpc.chord_id = c.id
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
					JOIN pitches p ON c.root_id = p.id
				WHERE
					kpc.key_id    = $1 AND
					c.ring_number = $2
				ORDER BY
					c.id
				OFFSET $3
				LIMIT  $4;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title:       "ReturnsChordsWhenSucceededWithCardinalityFilter",
			GivenFilter: theory.ChordFilter{Cardinality: 1},
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT c.id)
				FROM key_pitch_chords kpc
					JOIN chords c ON kpc.chord_id = c.id
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
				WHERE
					kpc.key_id     = $1 AND
					cq.cardinality = $2;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					c.id,
					c.name
				FROM key_pitch_chords kpc 
					JOIN chords c ON kpc.chord_id = c.id
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
					JOIN pitches p ON c.root_id = p.id
				WHERE
					kpc.key_id     = $1 AND
					cq.cardinality = $2
				ORDER BY
					c.id
				OFFSET $3
				LIMIT  $4;`,
			ExpectedCountArgs: []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg()},
			ExpectedListArgs:  []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
		},
		{
			Title: "ReturnsErrorWhenCountKeyChordsFailed",
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT c.id)
				FROM key_pitch_chords kpc
					JOIN chords c ON kpc.chord_id = c.id
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
				WHERE kpc.key_id = $1;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					c.id,
					c.name
				FROM key_pitch_chords kpc 
					JOIN chords c ON kpc.chord_id = c.id
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
					JOIN pitches p ON c.root_id = p.id
				WHERE
					kpc.key_id = $1
				ORDER BY
					c.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs:   []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:    []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
			CountKeyChordsError: sql.ErrConnDone,
		},
		{
			Title: "ReturnsErrorWhenListKeyChordsFailed",
			ExpectedCountQuery: `
				SELECT
					COUNT(DISTINCT c.id)
				FROM key_pitch_chords kpc
					JOIN chords c ON kpc.chord_id = c.id
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
				WHERE kpc.key_id = $1;`,
			ExpectedListQuery: `
				SELECT DISTINCT
					c.id,
					c.name
				FROM key_pitch_chords kpc 
					JOIN chords c ON kpc.chord_id = c.id
					JOIN chord_qualities cq ON c.chord_quality_id = cq.id
					JOIN pitches p ON c.root_id = p.id
				WHERE
					kpc.key_id = $1
				ORDER BY
					c.id
				OFFSET $2
				LIMIT  $3;`,
			ExpectedCountArgs:  []driver.Value{sqlmock.AnyArg()},
			ExpectedListArgs:   []driver.Value{sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()},
			ListKeyChordsError: sql.ErrConnDone,
		},
	}

	countChordsColumns := []string{"count"}

	listChordsColumns := []string{
		"id",
		"name",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger := mock.Logger()

			db, sqlMock, err := mockDatabase()
			require.NoError(t, err)

			if tc.CountKeyChordsError != nil {
				sqlMock.ExpectQuery(tc.ExpectedCountQuery).
					WithArgs(tc.ExpectedCountArgs...).
					WillReturnError(tc.CountKeyChordsError)
			} else {
				sqlMock.ExpectQuery(tc.ExpectedCountQuery).
					WithArgs(tc.ExpectedCountArgs...).
					WillReturnRows(sqlmock.NewRows(countChordsColumns).
						AddRow(1))

				if tc.ListKeyChordsError != nil {
					sqlMock.ExpectQuery(tc.ExpectedListQuery).
						WithArgs(tc.ExpectedListArgs...).
						WillReturnError(tc.ListKeyChordsError)
				} else {
					sqlMock.ExpectQuery(tc.ExpectedListQuery).
						WithArgs(tc.ExpectedListArgs...).
						WillReturnRows(sqlmock.NewRows(listChordsColumns).
							AddRow(1, "name1"))
				}
			}

			var pagination api.Pagination
			repository := theory.NewRepository(logger, db)
			chords, _, err := repository.ListKeyChords(context.Background(), 1, tc.GivenFilter, pagination)
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
			p.name
		FROM key_pitches k
			JOIN pitches p ON k.pitch_id = p.id
		WHERE
			k.key_id = $1
		ORDER BY
			p.id;`

	listKeyPitchesColumns := []string{
		"id",
		"name",
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger := mock.Logger()

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
						AddRow(1, "name"))
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
