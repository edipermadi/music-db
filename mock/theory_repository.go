package mock

import (
	"context"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/stretchr/testify/mock"
)

// TheoryRepositoryReturnValues stores return value of mocked theory.Repository
type TheoryRepositoryReturnValues struct {
	GetPitch        []interface{}
	ListPitchChords []interface{}
	ListPitchKeys   []interface{}
	ListPitchScales []interface{}
	ListPitches     []interface{}

	GetChord         []interface{}
	GetChordQuality  []interface{}
	ListChordKeys    []interface{}
	ListChordPitches []interface{}
	ListChordScales  []interface{}
	ListChords       []interface{}

	GetScale         []interface{}
	ListScaleChords  []interface{}
	ListScaleKeys    []interface{}
	ListScalePitches []interface{}
	ListScales       []interface{}

	GetKey         []interface{}
	ListKeyChords  []interface{}
	ListKeyModes   []interface{}
	ListKeyPitches []interface{}
	ListKeys       []interface{}
}

// TheoryRepository returns a mock implementation of theory.Repository
func TheoryRepository(values TheoryRepositoryReturnValues) theory.Repository {
	repository := &theoryRepository{}

	// setup mocked chord functions
	repository.On("GetChord", mock.Anything, mock.Anything).Return(values.GetChord...)
	repository.On("GetChordQuality", mock.Anything, mock.Anything).Return(values.GetChordQuality...)
	repository.On("ListChordKeys", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(values.ListChordKeys...)
	repository.On("ListChordPitches", mock.Anything, mock.Anything).Return(values.ListChordPitches...)
	repository.On("ListChordScales", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(values.ListChordScales...)
	repository.On("ListChords", mock.Anything, mock.Anything, mock.Anything).Return(values.ListChords...)

	// setup mocked key functions
	repository.On("GetKey", mock.Anything, mock.Anything).Return(values.GetKey...)
	repository.On("ListKeyChords", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(values.ListKeyChords...)
	repository.On("ListKeyModes", mock.Anything, mock.Anything, mock.Anything).Return(values.ListKeyModes...)
	repository.On("ListKeyPitches", mock.Anything, mock.Anything).Return(values.ListKeyPitches...)
	repository.On("ListKeys", mock.Anything, mock.Anything, mock.Anything).Return(values.ListKeys...)

	// setup mocked pitch functions
	repository.On("GetPitch", mock.Anything, mock.Anything).Return(values.GetPitch...)
	repository.On("ListPitchChords", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(values.ListPitchChords...)
	repository.On("ListPitchKeys", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(values.ListPitchKeys...)
	repository.On("ListPitchScales", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(values.ListPitchScales...)
	repository.On("ListPitches", mock.Anything).Return(values.ListPitches...)

	// setup mocked scale functions
	repository.On("GetScale", mock.Anything, mock.Anything).Return(values.GetScale...)
	repository.On("ListScaleChords", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(values.ListScaleChords...)
	repository.On("ListScaleKeys", mock.Anything, mock.Anything).Return(values.ListScaleKeys...)
	repository.On("ListScalePitches", mock.Anything, mock.Anything).Return(values.ListScalePitches...)
	repository.On("ListScales", mock.Anything, mock.Anything, mock.Anything).Return(values.ListScales...)

	return repository
}

// theoryRepository mock theory.Repository
type theoryRepository struct {
	mock.Mock
}

// ListPitches mock theory.Repository#ListPitches
func (m *theoryRepository) ListPitches(ctx context.Context) ([]theory.SimplifiedPitch, error) {
	args := m.Called(ctx)

	var entries []theory.SimplifiedPitch
	if v, ok := args.Get(0).([]theory.SimplifiedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// GetPitch mock theory.Repository#GetPitch
func (m *theoryRepository) GetPitch(ctx context.Context, pitchID int64) (*theory.DetailedPitch, error) {
	args := m.Called(ctx, pitchID)

	var entry *theory.DetailedPitch
	if v, ok := args.Get(0).(*theory.DetailedPitch); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// ListPitchChords mock theory.Repository#SimplifiedChord
func (m *theoryRepository) ListPitchChords(ctx context.Context, pitchID int64, filter theory.ChordFilter, pagination api.Pagination) ([]theory.SimplifiedChord, *api.Pagination, error) {
	args := m.Called(ctx, pitchID, filter, pagination)

	var entries []theory.SimplifiedChord
	if v, ok := args.Get(0).([]theory.SimplifiedChord); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(1).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// ListPitchKeys mock theory.Repository#ListPitchKeys
func (m *theoryRepository) ListPitchKeys(ctx context.Context, pitchID int64, filter theory.KeyFilter, pagination api.Pagination) ([]theory.SimplifiedKey, *api.Pagination, error) {
	args := m.Called(ctx, pitchID, filter, pagination)

	var entries []theory.SimplifiedKey
	if v, ok := args.Get(0).([]theory.SimplifiedKey); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(1).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// ListPitchScales mock theory.Repository#ListPitchScales
func (m *theoryRepository) ListPitchScales(ctx context.Context, pitchID int64, filter theory.ScaleFilter, pagination api.Pagination) ([]theory.SimplifiedScale, *api.Pagination, error) {
	args := m.Called(ctx, pitchID, filter, pagination)

	var entries []theory.SimplifiedScale
	if v, ok := args.Get(0).([]theory.SimplifiedScale); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(1).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// ListChords mock theory.Repository#ListChords
func (m *theoryRepository) ListChords(ctx context.Context, filter theory.ChordFilter, pagination api.Pagination) ([]theory.SimplifiedChord, *api.Pagination, error) {
	args := m.Called(ctx, filter, pagination)

	var entries []theory.SimplifiedChord
	if v, ok := args.Get(0).([]theory.SimplifiedChord); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(0).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// ListChordPitches mock theory.Repository#ListChordPitches
func (m *theoryRepository) ListChordPitches(ctx context.Context, chordID int64) ([]theory.SimplifiedPitch, error) {
	args := m.Called(ctx, chordID)

	var entries []theory.SimplifiedPitch
	if v, ok := args.Get(0).([]theory.SimplifiedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListChordKeys mock theory.Repository#ListChordKeys
func (m *theoryRepository) ListChordKeys(ctx context.Context, chordID int64, filter theory.KeyFilter, pagination api.Pagination) ([]theory.SimplifiedKey, *api.Pagination, error) {
	args := m.Called(ctx, chordID, filter, pagination)

	var entries []theory.SimplifiedKey
	if v, ok := args.Get(0).([]theory.SimplifiedKey); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(0).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// ListChordScales mock theory.Repository#ListChordScales
func (m *theoryRepository) ListChordScales(ctx context.Context, chordID int64, filter theory.ScaleFilter, pagination api.Pagination) ([]theory.SimplifiedScale, *api.Pagination, error) {
	args := m.Called(ctx, chordID, filter, pagination)

	var entries []theory.SimplifiedScale
	if v, ok := args.Get(0).([]theory.SimplifiedScale); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(0).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// GetChord mock theory.Repository#GetChord
func (m *theoryRepository) GetChord(ctx context.Context, chordID int64) (*theory.DetailedChord, error) {
	args := m.Called(ctx, chordID)

	var entry *theory.DetailedChord
	if v, ok := args.Get(0).(*theory.DetailedChord); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// GetChordQuality mock theory.Repository#GetChordQuality
func (m *theoryRepository) GetChordQuality(ctx context.Context, chordID int64) (*theory.DetailedChordQuality, error) {
	args := m.Called(ctx, chordID)

	var entry *theory.DetailedChordQuality
	if v, ok := args.Get(0).(*theory.DetailedChordQuality); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// ListScales mock theory.Repository#ListScales
func (m *theoryRepository) ListScales(ctx context.Context, filter theory.ScaleFilter, pagination api.Pagination) ([]theory.SimplifiedScale, *api.Pagination, error) {
	args := m.Called(ctx, filter, pagination)

	var entries []theory.SimplifiedScale
	if v, ok := args.Get(0).([]theory.SimplifiedScale); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(0).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// ListScaleKeys mock theory.Repository#ListScaleKeys
func (m *theoryRepository) ListScaleKeys(ctx context.Context, scaleID int64) ([]theory.SimplifiedKey, error) {
	args := m.Called(ctx, scaleID)

	var entries []theory.SimplifiedKey
	if v, ok := args.Get(0).([]theory.SimplifiedKey); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListScalePitches mock theory.Repository#ListScalePitches
func (m *theoryRepository) ListScalePitches(ctx context.Context, scaleID int64) ([]theory.SimplifiedPitch, error) {
	args := m.Called(ctx, scaleID)

	var entries []theory.SimplifiedPitch
	if v, ok := args.Get(0).([]theory.SimplifiedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListScaleChords mock theory.Repository#ListScaleChords
func (m *theoryRepository) ListScaleChords(ctx context.Context, scaleID int64, filter theory.ChordFilter, pagination api.Pagination) ([]theory.SimplifiedChord, *api.Pagination, error) {
	args := m.Called(ctx, scaleID, filter, pagination)

	var entries []theory.SimplifiedChord
	if v, ok := args.Get(0).([]theory.SimplifiedChord); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(1).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// GetScale mock theory.Repository#GetScale
func (m *theoryRepository) GetScale(ctx context.Context, scaleID int64) (*theory.DetailedScale, error) {
	args := m.Called(ctx, scaleID)

	var entry *theory.DetailedScale
	if v, ok := args.Get(0).(*theory.DetailedScale); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// ListKeys mock theory.Repository#ListKeys
func (m *theoryRepository) ListKeys(ctx context.Context, filter theory.KeyFilter, pagination api.Pagination) ([]theory.SimplifiedKey, *api.Pagination, error) {
	args := m.Called(ctx, filter, pagination)

	var entries []theory.SimplifiedKey
	if v, ok := args.Get(0).([]theory.SimplifiedKey); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(0).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// ListKeyModes mock theory.Repository#ListKeyModes
func (m *theoryRepository) ListKeyModes(ctx context.Context, keyID int64, filter theory.KeyFilter) ([]theory.SimplifiedKey, error) {
	args := m.Called(ctx, keyID, filter)

	var entries []theory.SimplifiedKey
	if v, ok := args.Get(0).([]theory.SimplifiedKey); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListKeyChords mock theory.Repository#ListKeyChords
func (m *theoryRepository) ListKeyChords(ctx context.Context, keyID int64, filter theory.ChordFilter, pagination api.Pagination) ([]theory.SimplifiedChord, *api.Pagination, error) {
	args := m.Called(ctx, keyID, filter, pagination)

	var entries []theory.SimplifiedChord
	if v, ok := args.Get(0).([]theory.SimplifiedChord); ok {
		entries = v
	}

	var paginationOut *api.Pagination
	if v, ok := args.Get(0).(*api.Pagination); ok {
		paginationOut = v
	}

	return entries, paginationOut, args.Error(2)
}

// ListKeyPitches mock theory.Repository#ListKeyPitches
func (m *theoryRepository) ListKeyPitches(ctx context.Context, keyID int64) ([]theory.SimplifiedPitch, error) {
	args := m.Called(ctx, keyID)

	var entries []theory.SimplifiedPitch
	if v, ok := args.Get(0).([]theory.SimplifiedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// GetKey mock theory.Repository#GetKey
func (m *theoryRepository) GetKey(ctx context.Context, keyID int64) (*theory.DetailedKey, error) {
	args := m.Called(ctx, keyID)

	var entry *theory.DetailedKey
	if v, ok := args.Get(0).(*theory.DetailedKey); ok {
		entry = v
	}

	return entry, args.Error(1)
}
