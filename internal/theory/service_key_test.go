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

func TestTheoryService_ListKeys(t *testing.T) {
	testCases := []serviceTestCase{
		{
			Title: "ReturnsKeysWhenSucceeded",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListKeys: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			service := tc.mockedService()
			entries, _, err := service.ListKeys(context.Background(), theory.KeyFilter{}, api.Pagination{})
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

func TestTheoryService_ListKeyModes(t *testing.T) {
	testCases := []serviceTestCase{
		{
			Title: "ReturnsKeysWhenSucceeded",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListKeyModes: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListKeyModes: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			service := tc.mockedService()
			entries, err := service.ListKeyModes(context.Background(), 1, theory.KeyFilter{})
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

func TestTheoryService_ListKeyChords(t *testing.T) {
	testCases := []serviceTestCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListKeyChords: []interface{}{[]theory.SimplifiedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListKeyChords: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			service := tc.mockedService()
			entries, _, err := service.ListKeyChords(context.Background(), 1, theory.ChordFilter{}, api.Pagination{})
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

func TestTheoryService_ListKeyPitches(t *testing.T) {
	testCases := []serviceTestCase{
		{
			Title: "ReturnsPitchesWhenSucceeded",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListKeyPitches: []interface{}{[]theory.SimplifiedPitch{{ID: 1, Name: "name"}}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				ListKeyPitches: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			service := tc.mockedService()
			entries, err := service.ListKeyPitches(context.Background(), 1)
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

func TestTheoryService_GetKey(t *testing.T) {
	testCases := []serviceTestCase{
		{
			Title: "ReturnsKeyWhenSucceeded",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				GetKey: []interface{}{&theory.DetailedKey{ID: 1, Name: "name", ZeitlerNumber: 2}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.TheoryRepositoryReturnValues{
				GetKey: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			service := tc.mockedService()
			entry, err := service.GetKey(context.Background(), 1)
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
