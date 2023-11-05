package theory_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/edipermadi/music-db/mock"
	"github.com/stretchr/testify/require"
)

func TestTheoryService_ListPitches(t *testing.T) {
	testCases := []serviceTestCase{
		{
			Title: "ReturnsPitchesWhenSucceeded",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListPitches: []interface{}{[]theory.SimplifiedPitch{{ID: 1, Name: "name"}}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListPitches: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			service := tc.mockedService()
			entries, err := service.ListPitches(context.Background())
			if strings.HasPrefix(tc.Title, "ReturnsError") {
				require.Error(t, err)
				require.Empty(t, entries)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, entries)
			}
		})
	}
}

func TestTheoryService_GetPitch(t *testing.T) {
	testCases := []serviceTestCase{
		{
			Title: "ReturnsPitchWhenSucceeded",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				GetPitch: []interface{}{&theory.DetailedPitch{ID: 1, Name: "name", ZeitlerNumber: 2}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				GetPitch: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			service := tc.mockedService()
			entry, err := service.GetPitch(context.Background(), 1)
			if strings.HasPrefix(tc.Title, "ReturnsError") {
				require.Error(t, err)
				require.Empty(t, entry)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, entry)
			}
		})
	}
}

func TestTheoryService_ListPitchChords(t *testing.T) {
	testCases := []serviceTestCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListPitchChords: []interface{}{[]theory.SimplifiedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListPitchChords: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			service := tc.mockedService()
			entry, _, err := service.ListPitchChords(context.Background(), 1, theory.ChordFilter{}, api.Pagination{})
			if strings.HasPrefix(tc.Title, "ReturnsError") {
				require.Error(t, err)
				require.Empty(t, entry)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, entry)
			}
		})
	}
}

func TestTheoryService_ListPitchKeys(t *testing.T) {
	testCases := []serviceTestCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListPitchKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListPitchKeys: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			service := tc.mockedService()
			entry, _, err := service.ListPitchKeys(context.Background(), 1, theory.KeyFilter{}, api.Pagination{})
			if strings.HasPrefix(tc.Title, "ReturnsError") {
				require.Error(t, err)
				require.Empty(t, entry)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, entry)
			}
		})
	}
}

func TestTheoryService_ListPitchScales(t *testing.T) {
	testCases := []serviceTestCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListPitchScales: []interface{}{[]theory.SimplifiedScale{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListPitchScales: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			service := tc.mockedService()
			entry, _, err := service.ListPitchScales(context.Background(), 1, theory.ScaleFilter{}, api.Pagination{})
			if strings.HasPrefix(tc.Title, "ReturnsError") {
				require.Error(t, err)
				require.Empty(t, entry)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, entry)
			}
		})
	}
}
