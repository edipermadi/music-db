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

func TestTheoryService_ListChords(t *testing.T) {
	testCases := []serviceTestCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListChords: []interface{}{[]theory.SimplifiedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListChords: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			service := tc.mockedService()
			entries, _, err := service.ListChords(context.Background(), theory.ChordFilter{}, api.Pagination{})
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

func TestTheoryService_ListChordKeys(t *testing.T) {
	testCases := []serviceTestCase{
		{
			Title: "ReturnsKeysWhenSucceeded",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListChordKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListChordKeys: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			service := tc.mockedService()
			entries, _, err := service.ListChordKeys(context.Background(), 1, theory.KeyFilter{}, api.Pagination{})
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

func TestTheoryService_ListChordScales(t *testing.T) {
	testCases := []serviceTestCase{
		{
			Title: "ReturnsKeysWhenSucceeded",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListChordScales: []interface{}{[]theory.SimplifiedScale{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListChordScales: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			service := tc.mockedService()
			entries, _, err := service.ListChordScales(context.Background(), 1, theory.ScaleFilter{}, api.Pagination{})
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

func TestTheoryService_ListChordPitches(t *testing.T) {
	testCases := []serviceTestCase{
		{
			Title: "ReturnsPitchesWhenSucceeded",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListChordPitches: []interface{}{[]theory.SimplifiedPitch{{ID: 1, Name: "name"}}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListChordPitches: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			service := tc.mockedService()
			entries, err := service.ListChordPitches(context.Background(), 1)
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

func TestTheoryService_GetChord(t *testing.T) {
	testCases := []serviceTestCase{
		{
			Title: "ReturnsChordWhenSucceeded",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				GetChord: []interface{}{&theory.DetailedChord{ID: 1, Name: "name", ZeitlerNumber: 2}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				GetChord: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			service := tc.mockedService()
			entry, err := service.GetChord(context.Background(), 1)
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

func TestTheoryService_GetChordQuality(t *testing.T) {
	testCases := []serviceTestCase{
		{
			Title: "ReturnsChordQualityWhenSucceeded",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				GetChordQuality: []interface{}{&theory.DetailedChordQuality{ID: 1, Name: "name", ZeitlerNumber: 2}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				GetChordQuality: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			service := tc.mockedService()
			entry, err := service.GetChordQuality(context.Background(), 1)
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
