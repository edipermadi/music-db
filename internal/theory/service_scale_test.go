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

func TestTheoryService_ListScales(t *testing.T) {
	type testCase struct {
		Title                  string
		RepositoryReturnValues mock.RepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListScales: []interface{}{[]theory.DetailedScale{{ID: 1, Name: "name", Number: 2}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListScales: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &mock.TheoryRepository{}
			repository.On("ListScales", mock2.Anything, mock2.Anything).
				Return(tc.RepositoryReturnValues.ListScales...)

			service := theory.NewService(logger, repository)
			entries, _, err := service.ListScales(context.Background(), api.Pagination{})
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

func TestTheoryService_ListScaleKeys(t *testing.T) {
	type testCase struct {
		Title                  string
		RepositoryReturnValues mock.RepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsKeysWhenSucceeded",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListScaleKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name", Number: 2}}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				ListScaleKeys: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &mock.TheoryRepository{}
			repository.On("ListScaleKeys", mock2.Anything, mock2.Anything).
				Return(tc.RepositoryReturnValues.ListScaleKeys...)

			service := theory.NewService(logger, repository)
			entries, err := service.ListScaleKeys(context.Background(), 1)
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

func TestTheoryService_GetScale(t *testing.T) {
	type testCase struct {
		Title                  string
		RepositoryReturnValues mock.RepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsScaleWhenSucceeded",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				GetScale: []interface{}{&theory.DetailedScale{ID: 1, Name: "name", Number: 2}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			RepositoryReturnValues: mock.RepositoryReturnValues{
				GetScale: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &mock.TheoryRepository{}
			repository.On("GetScale", mock2.Anything, mock2.Anything).
				Return(tc.RepositoryReturnValues.GetScale...)

			service := theory.NewService(logger, repository)
			entry, err := service.GetScale(context.Background(), 1)
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
