package theory_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestNewService(t *testing.T) {
	logger, err := zap.NewProduction()
	require.NoError(t, err)

	repository := &MockRepository{}
	instance := theory.NewService(logger, repository)
	require.Implements(t, (*theory.Service)(nil), instance)
}

func TestTheoryService_GetChord(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordWhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetChord: []interface{}{&theory.DetailedChord{ID: 1, Name: "name", Number: 2}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetChord: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &MockRepository{}
			repository.On("GetChord", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.GetChord...)

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
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordQualityWhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetChordQuality: []interface{}{&theory.DetailedChordQuality{ID: 1, Name: "name", Number: 2}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetChordQuality: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &MockRepository{}
			repository.On("GetChordQuality", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.GetChordQuality...)

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

func TestTheoryService_GetKey(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsKeyWhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetKey: []interface{}{&theory.DetailedKey{ID: 1, Name: "name", Number: 2}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetKey: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &MockRepository{}
			repository.On("GetKey", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.GetKey...)

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

func TestTheoryService_GetScale(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsScaleWhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetScale: []interface{}{&theory.DetailedScale{ID: 1, Name: "name", Number: 2}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetScale: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &MockRepository{}
			repository.On("GetScale", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.GetScale...)

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

func TestTheoryService_ListChordKeys(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsKeysWhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListChordKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name", Number: 2}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListChordKeys: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &MockRepository{}
			repository.On("ListChordKeys", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListChordKeys...)

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
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsPitchesWhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListChordPitches: []interface{}{[]theory.DetailedPitch{{ID: 1, Name: "name", Number: 2}}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListChordPitches: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &MockRepository{}
			repository.On("ListChordPitches", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListChordPitches...)

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

func TestTheoryService_ListChords(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListChords: []interface{}{[]theory.DetailedChord{{ID: 1, Name: "name", Number: 2}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListChords: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &MockRepository{}
			repository.On("ListChords", mock.Anything, mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListChords...)

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

func TestTheoryService_ListKeyChords(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListKeyChords: []interface{}{[]theory.DetailedChord{{ID: 1, Name: "name", Number: 2}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListKeyChords: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &MockRepository{}
			repository.On("ListKeyChords", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListKeyChords...)

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

func TestTheoryService_ListKeyModes(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsKeysWhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListKeyModes: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name", Number: 2}}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListKeyModes: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &MockRepository{}
			repository.On("ListKeyModes", mock.Anything, mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListKeyModes...)

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

func TestTheoryService_ListKeyPitches(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsPitchesWhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListKeyPitches: []interface{}{[]theory.DetailedPitch{{ID: 1, Name: "name", Number: 2}}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListKeyPitches: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &MockRepository{}
			repository.On("ListKeyPitches", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListKeyPitches...)

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

func TestTheoryService_ListKeys(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsKeysWhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name", Number: 2}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListKeys: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &MockRepository{}
			repository.On("ListKeys", mock.Anything, mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListKeys...)

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

func TestTheoryService_ListPitches(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsPitchesWhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListPitches: []interface{}{[]theory.DetailedPitch{{ID: 1, Name: "name", Number: 2}}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListPitches: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &MockRepository{}
			repository.On("ListPitches", mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListPitches...)

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

func TestTheoryService_ListScaleKeys(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsKeysWhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListScaleKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name", Number: 2}}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListScaleKeys: []interface{}{nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &MockRepository{}
			repository.On("ListScaleKeys", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListScaleKeys...)

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

func TestTheoryService_ListScales(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
	}

	testCases := []testCase{
		{
			Title: "ReturnsChordsWhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListScales: []interface{}{[]theory.DetailedScale{{ID: 1, Name: "name", Number: 2}}, &api.Pagination{}, nil},
			},
		},
		{
			Title: "ReturnsErrorWhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListScales: []interface{}{nil, nil, errors.New("error")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			repository := &MockRepository{}
			repository.On("ListScales", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListScales...)

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

// MockServiceReturnValues stores return value of mocked theory.Service
type MockServiceReturnValues struct {
	ListPitches []interface{}

	ListChords       []interface{}
	ListChordPitches []interface{}
	ListChordKeys    []interface{}
	GetChord         []interface{}
	GetChordQuality  []interface{}

	ListScales    []interface{}
	ListScaleKeys []interface{}
	GetScale      []interface{}

	ListKeys       []interface{}
	ListKeyModes   []interface{}
	ListKeyChords  []interface{}
	ListKeyPitches []interface{}
	GetKey         []interface{}
}

// MockService mocks theory.Service
type MockService struct {
	mock.Mock
}

// ListPitches mocks theory.Service#ListPitches
func (m *MockService) ListPitches(ctx context.Context) ([]theory.DetailedPitch, error) {
	args := m.Called(ctx)

	var entries []theory.DetailedPitch
	if v, ok := args.Get(0).([]theory.DetailedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListChords mocks theory.Service#ListChords
func (m *MockService) ListChords(ctx context.Context, filter theory.ChordFilter, pagination api.Pagination) ([]theory.DetailedChord, *api.Pagination, error) {
	args := m.Called(ctx, filter, pagination)

	var entries []theory.DetailedChord
	if v, ok := args.Get(0).([]theory.DetailedChord); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(0).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// ListChordPitches mocks theory.Service#ListChordPitches
func (m *MockService) ListChordPitches(ctx context.Context, chordID int64) ([]theory.DetailedPitch, error) {
	args := m.Called(ctx, chordID)

	var entries []theory.DetailedPitch
	if v, ok := args.Get(0).([]theory.DetailedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListChordKeys mocks theory.Service#ListChordKeys
func (m *MockService) ListChordKeys(ctx context.Context, chordID int64, filter theory.KeyFilter, pagination api.Pagination) ([]theory.DetailedKey, *api.Pagination, error) {
	args := m.Called(ctx, chordID, filter, pagination)

	var entries []theory.DetailedKey
	if v, ok := args.Get(0).([]theory.DetailedKey); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(0).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// GetChord mocks theory.Service#GetChord
func (m *MockService) GetChord(ctx context.Context, chordID int64) (*theory.DetailedChord, error) {
	args := m.Called(ctx, chordID)

	var entry *theory.DetailedChord
	if v, ok := args.Get(0).(*theory.DetailedChord); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// GetChordQuality mocks theory.Service#GetChordQuality
func (m *MockService) GetChordQuality(ctx context.Context, chordID int64) (*theory.DetailedChordQuality, error) {
	args := m.Called(ctx, chordID)

	var entry *theory.DetailedChordQuality
	if v, ok := args.Get(0).(*theory.DetailedChordQuality); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// ListScales mocks theory.Service#ListScales
func (m *MockService) ListScales(ctx context.Context, pagination api.Pagination) ([]theory.DetailedScale, *api.Pagination, error) {
	args := m.Called(ctx, pagination)

	var entries []theory.DetailedScale
	if v, ok := args.Get(0).([]theory.DetailedScale); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(0).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// ListScaleKeys mocks theory.Service#ListScaleKeys
func (m *MockService) ListScaleKeys(ctx context.Context, scaleID int64) ([]theory.DetailedKey, error) {
	args := m.Called(ctx, scaleID)

	var entries []theory.DetailedKey
	if v, ok := args.Get(0).([]theory.DetailedKey); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// GetScale mocks theory.Service#GetScale
func (m *MockService) GetScale(ctx context.Context, scaleID int64) (*theory.DetailedScale, error) {
	args := m.Called(ctx, scaleID)

	var entry *theory.DetailedScale
	if v, ok := args.Get(0).(*theory.DetailedScale); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// ListKeys mocks theory.Service#ListKeys
func (m *MockService) ListKeys(ctx context.Context, filter theory.KeyFilter, pagination api.Pagination) ([]theory.DetailedKey, *api.Pagination, error) {
	args := m.Called(ctx, filter, pagination)

	var entries []theory.DetailedKey
	if v, ok := args.Get(0).([]theory.DetailedKey); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(0).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// ListKeyModes mocks theory.Service#ListKeyModes
func (m *MockService) ListKeyModes(ctx context.Context, keyID int64, filter theory.KeyFilter) ([]theory.DetailedKey, error) {
	args := m.Called(ctx, keyID, filter)

	var entries []theory.DetailedKey
	if v, ok := args.Get(0).([]theory.DetailedKey); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListKeyChords mocks theory.Service#ListKeyChords
func (m *MockService) ListKeyChords(ctx context.Context, keyID int64, filter theory.ChordFilter, pagination api.Pagination) ([]theory.DetailedChord, *api.Pagination, error) {
	args := m.Called(ctx, keyID, filter, pagination)

	var entries []theory.DetailedChord
	if v, ok := args.Get(0).([]theory.DetailedChord); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(0).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// ListKeyPitches mocks theory.Service#ListKeyPitches
func (m *MockService) ListKeyPitches(ctx context.Context, keyID int64) ([]theory.DetailedPitch, error) {
	args := m.Called(ctx, keyID)

	var entries []theory.DetailedPitch
	if v, ok := args.Get(0).([]theory.DetailedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// GetKey mocks theory.Service#GetKey
func (m *MockService) GetKey(ctx context.Context, keyID int64) (*theory.DetailedKey, error) {
	args := m.Called(ctx, keyID)

	var entry *theory.DetailedKey
	if v, ok := args.Get(0).(*theory.DetailedKey); ok {
		entry = v
	}

	return entry, args.Error(1)
}
