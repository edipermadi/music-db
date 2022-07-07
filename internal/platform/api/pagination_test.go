package api_test

import (
	"testing"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/stretchr/testify/assert"
)

func TestPagination_Offset(t *testing.T) {
	assert.Equal(t, 0, api.Pagination{Page: 1, PerPage: 10}.Offset())
	assert.Equal(t, 10, api.Pagination{Page: 2, PerPage: 10}.Offset())
}

func TestPagination_Sanitize(t *testing.T) {
	type testCase struct {
		Title    string
		Given    api.Pagination
		Expected api.Pagination
	}

	testCases := []testCase{
		{
			Title:    "SanitizesEmptyPage",
			Given:    api.Pagination{Page: 0, PerPage: 1},
			Expected: api.Pagination{Page: 1, PerPage: 1},
		},
		{
			Title:    "SanitizesEmptyPerPage",
			Given:    api.Pagination{Page: 1, PerPage: 0},
			Expected: api.Pagination{Page: 1, PerPage: 50},
		},
		{
			Title:    "SanitizesOutOfRangePerPage",
			Given:    api.Pagination{Page: 1, PerPage: 200},
			Expected: api.Pagination{Page: 1, PerPage: 100},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			tc.Given.Sanitize()
			assert.Equal(t, tc.Expected, tc.Given)
		})
	}
}
