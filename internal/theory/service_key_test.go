package theory_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/edipermadi/music-db/mock"
	mock2 "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestTheoryService_ListKeys(t *testing.T) {
	type testCase struct {
		Title                  string
		RepositoryReturnValues mock.RepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsKeysWhenSucceeded",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name", Number: 2}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListKeys: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &mock.TheoryRepository{}
			repository.On("ListKeys", mock2.Anything, mock2.Anything, mock2.Anything).
				Return(tc.RepositoryReturnValues.ListKeys...)

			service := theory.NewService(logger, repository)
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
	type testCase struct {
		Title                  string
		RepositoryReturnValues mock.RepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsKeysWhenSucceeded",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListKeyModes: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name", Number: 2}}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListKeyModes: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &mock.TheoryRepository{}
			repository.On("ListKeyModes", mock2.Anything, mock2.Anything, mock2.Anything).
				Return(tc.RepositoryReturnValues.ListKeyModes...)

			service := theory.NewService(logger, repository)
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
	type testCase struct {
		Title                  string
		RepositoryReturnValues mock.RepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListKeyChords: []interface{}{[]theory.DetailedChord{{ID: 1, Name: "name", Number: 2}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListKeyChords: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &mock.TheoryRepository{}
			repository.On("ListKeyChords", mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything).
				Return(tc.RepositoryReturnValues.ListKeyChords...)

			service := theory.NewService(logger, repository)
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
	type testCase struct {
		Title                  string
		RepositoryReturnValues mock.RepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsPitchesWhenSucceeded",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListKeyPitches: []interface{}{[]theory.DetailedPitch{{ID: 1, Name: "name", Number: 2}}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListKeyPitches: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &mock.TheoryRepository{}
			repository.On("ListKeyPitches", mock2.Anything, mock2.Anything).
				Return(tc.RepositoryReturnValues.ListKeyPitches...)

			service := theory.NewService(logger, repository)
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
	type testCase struct {
		Title                  string
		RepositoryReturnValues mock.RepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsKeyWhenSucceeded",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				GetKey: []interface{}{&theory.DetailedKey{ID: 1, Name: "name", Number: 2}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				GetKey: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &mock.TheoryRepository{}
			repository.On("GetKey", mock2.Anything, mock2.Anything).
				Return(tc.RepositoryReturnValues.GetKey...)

			service := theory.NewService(logger, repository)
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
