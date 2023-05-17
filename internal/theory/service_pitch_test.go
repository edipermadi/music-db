package theory_test

import (
	"context"
	"errors"
	"github.com/edipermadi/music-db/internal/platform/api"
	"strings"
	"testing"

	"github.com/edipermadi/music-db/internal/theory"
	"github.com/edipermadi/music-db/mock"
	mock2 "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestTheoryService_ListPitches(t *testing.T) {
	type testCase struct {
		Title                  string
		RepositoryReturnValues mock.RepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsPitchesWhenSucceeded",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListPitches: []interface{}{[]theory.DetailedPitch{{ID: 1, Name: "name", Number: 2}}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListPitches: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &mock.TheoryRepository{}
			repository.On("ListPitches", mock2.Anything).
				Return(tc.RepositoryReturnValues.ListPitches...)

			service := theory.NewService(logger, repository)
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
	type testCase struct {
		Title                  string
		RepositoryReturnValues mock.RepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsPitchWhenSucceeded",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				GetPitch: []interface{}{&theory.DetailedPitch{ID: 1, Name: "name", Number: 2}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				GetPitch: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &mock.TheoryRepository{}
			repository.On("GetPitch", mock2.Anything, mock2.Anything).
				Return(tc.RepositoryReturnValues.GetPitch...)

			service := theory.NewService(logger, repository)
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
	type testCase struct {
		Title                  string
		RepositoryReturnValues mock.RepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListPitchChords: []interface{}{[]theory.SimplifiedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListPitchChords: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &mock.TheoryRepository{}
			repository.On("ListPitchChords", mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything).
				Return(tc.RepositoryReturnValues.ListPitchChords...)

			service := theory.NewService(logger, repository)
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
	type testCase struct {
		Title                  string
		RepositoryReturnValues mock.RepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListPitchKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListPitchKeys: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &mock.TheoryRepository{}
			repository.On("ListPitchKeys", mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything).
				Return(tc.RepositoryReturnValues.ListPitchKeys...)

			service := theory.NewService(logger, repository)
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
