package theory_test

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"strings"
	"testing"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/edipermadi/music-db/mock"
	mock2 "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestTheoryService_ListChords(t *testing.T) {
	type testCase struct {
		Title                  string
		RepositoryReturnValues mock.RepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListChords: []interface{}{[]theory.DetailedChord{{ID: 1, Name: "name", Number: 2}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListChords: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &mock.TheoryRepository{}
			repository.On("ListChords", mock2.Anything, mock2.Anything, mock2.Anything).
				Return(tc.RepositoryReturnValues.ListChords...)

			service := theory.NewService(logger, repository)
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
	type testCase struct {
		Title                  string
		RepositoryReturnValues mock.RepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsKeysWhenSucceeded",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListChordKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name", Number: 2}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListChordKeys: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &mock.TheoryRepository{}
			repository.On("ListChordKeys", mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything).
				Return(tc.RepositoryReturnValues.ListChordKeys...)

			service := theory.NewService(logger, repository)
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

func TestTheoryService_ListChordPitches(t *testing.T) {
	type testCase struct {
		Title                  string
		RepositoryReturnValues mock.RepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsPitchesWhenSucceeded",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListChordPitches: []interface{}{[]theory.DetailedPitch{{ID: 1, Name: "name", Number: 2}}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListChordPitches: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &mock.TheoryRepository{}
			repository.On("ListChordPitches", mock2.Anything, mock2.Anything).
				Return(tc.RepositoryReturnValues.ListChordPitches...)

			service := theory.NewService(logger, repository)
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
	type testCase struct {
		Title                  string
		RepositoryReturnValues mock.RepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordWhenSucceeded",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				GetChord: []interface{}{&theory.DetailedChord{ID: 1, Name: "name", Number: 2}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				GetChord: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &mock.TheoryRepository{}
			repository.On("GetChord", mock2.Anything, mock2.Anything).
				Return(tc.RepositoryReturnValues.GetChord...)

			service := theory.NewService(logger, repository)
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
	type testCase struct {
		Title                  string
		RepositoryReturnValues mock.RepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordQualityWhenSucceeded",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				GetChordQuality: []interface{}{&theory.DetailedChordQuality{ID: 1, Name: "name", Number: 2}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				GetChordQuality: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &mock.TheoryRepository{}
			repository.On("GetChordQuality", mock2.Anything, mock2.Anything).
				Return(tc.RepositoryReturnValues.GetChordQuality...)

			service := theory.NewService(logger, repository)
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
