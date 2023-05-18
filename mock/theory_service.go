package mock

import (
	"context"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/stretchr/testify/mock"
)

// TheoryServiceReturnValues stores return value of mocked theory.Service
type TheoryServiceReturnValues struct {
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

// TheoryService return a mocked implementation of theory.Serice
func TheoryService(values TheoryServiceReturnValues) theory.Service {
	service := &theoryService{}

	// setup mocked chord functions
	service.On("GetChord", mock.Anything, mock.Anything).Return(values.GetChord...)
	service.On("GetChordQuality", mock.Anything, mock.Anything).Return(values.GetChordQuality...)
	service.On("ListChordKeys", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(values.ListChordKeys...)
	service.On("ListChordScales", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(values.ListChordScales...)
	service.On("ListChordPitches", mock.Anything, mock.Anything).Return(values.ListChordPitches...)
	service.On("ListChords", mock.Anything, mock.Anything, mock.Anything).Return(values.ListChords...)

	// setup mocked key functions
	service.On("GetKey", mock.Anything, mock.Anything).Return(values.GetKey...)
	service.On("ListKeyChords", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(values.ListKeyChords...)
	service.On("ListKeyModes", mock.Anything, mock.Anything, mock.Anything).Return(values.ListKeyModes...)
	service.On("ListKeyPitches", mock.Anything, mock.Anything).Return(values.ListKeyPitches...)
	service.On("ListKeys", mock.Anything, mock.Anything, mock.Anything).Return(values.ListKeys...)

	// setup mocked pitch functions
	service.On("GetPitch", mock.Anything, mock.Anything).Return(values.GetPitch...)
	service.On("ListPitchChords", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(values.ListPitchChords...)
	service.On("ListPitchKeys", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(values.ListPitchKeys...)
	service.On("ListPitchScales", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(values.ListPitchScales...)
	service.On("ListPitches", mock.Anything, mock.Anything).Return(values.ListPitches...)

	// setup mocked scale functions
	service.On("GetScale", mock.Anything, mock.Anything).Return(values.GetScale...)
	service.On("ListScaleKeys", mock.Anything, mock.Anything).Return(values.ListScaleKeys...)
	service.On("ListScaleChords", mock.Anything, mock.Anything).Return(values.ListScaleChords...)
	service.On("ListScalePitches", mock.Anything, mock.Anything).Return(values.ListScalePitches...)
	service.On("ListScales", mock.Anything, mock.Anything, mock.Anything).Return(values.ListScales...)

	return service
}

// theoryService mock theory.Service
type theoryService struct {
	mock.Mock
}

// ListPitches mock theory.Service#ListPitches
func (m *theoryService) ListPitches(ctx context.Context) ([]theory.SimplifiedPitch, error) {
	args := m.Called(ctx)

	var entries []theory.SimplifiedPitch
	if v, ok := args.Get(0).([]theory.SimplifiedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// GetPitch mock theory.Service#GetPitch
func (m *theoryService) GetPitch(ctx context.Context, pitchID int64) (*theory.DetailedPitch, error) {
	args := m.Called(ctx, pitchID)

	var entry *theory.DetailedPitch
	if v, ok := args.Get(0).(*theory.DetailedPitch); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// ListPitchChords mock theory.Service#ListPitchChords
func (m *theoryService) ListPitchChords(ctx context.Context, pitchID int64, filter theory.ChordFilter, pagination api.Pagination) ([]theory.SimplifiedChord, *api.Pagination, error) {
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

// ListPitchKeys mock theory.Service#ListPitchKeys
func (m *theoryService) ListPitchKeys(ctx context.Context, pitchID int64, filter theory.KeyFilter, pagination api.Pagination) ([]theory.SimplifiedKey, *api.Pagination, error) {
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

// ListPitchScales mock theory.Service#ListPitchScales
func (m *theoryService) ListPitchScales(ctx context.Context, pitchID int64, filter theory.ScaleFilter, pagination api.Pagination) ([]theory.SimplifiedScale, *api.Pagination, error) {
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

// ListChords mock theory.Service#ListChords
func (m *theoryService) ListChords(ctx context.Context, filter theory.ChordFilter, pagination api.Pagination) ([]theory.SimplifiedChord, *api.Pagination, error) {
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

// ListChordPitches mock theory.Service#ListChordPitches
func (m *theoryService) ListChordPitches(ctx context.Context, chordID int64) ([]theory.SimplifiedPitch, error) {
	args := m.Called(ctx, chordID)

	var entries []theory.SimplifiedPitch
	if v, ok := args.Get(0).([]theory.SimplifiedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListChordKeys mock theory.Service#ListChordKeys
func (m *theoryService) ListChordKeys(ctx context.Context, chordID int64, filter theory.KeyFilter, pagination api.Pagination) ([]theory.SimplifiedKey, *api.Pagination, error) {
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

// ListChordScales mock theory.Service#ListChordScales
func (m *theoryService) ListChordScales(ctx context.Context, chordID int64, filter theory.ScaleFilter, pagination api.Pagination) ([]theory.SimplifiedScale, *api.Pagination, error) {
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

// GetChord mock theory.Service#GetChord
func (m *theoryService) GetChord(ctx context.Context, chordID int64) (*theory.DetailedChord, error) {
	args := m.Called(ctx, chordID)

	var entry *theory.DetailedChord
	if v, ok := args.Get(0).(*theory.DetailedChord); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// GetChordQuality mock theory.Service#GetChordQuality
func (m *theoryService) GetChordQuality(ctx context.Context, chordID int64) (*theory.DetailedChordQuality, error) {
	args := m.Called(ctx, chordID)

	var entry *theory.DetailedChordQuality
	if v, ok := args.Get(0).(*theory.DetailedChordQuality); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// ListScales mock theory.Service#ListScales
func (m *theoryService) ListScales(ctx context.Context, filter theory.ScaleFilter, pagination api.Pagination) ([]theory.SimplifiedScale, *api.Pagination, error) {
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

// ListScaleKeys mock theory.Service#ListScaleKeys
func (m *theoryService) ListScaleKeys(ctx context.Context, scaleID int64) ([]theory.SimplifiedKey, error) {
	args := m.Called(ctx, scaleID)

	var entries []theory.SimplifiedKey
	if v, ok := args.Get(0).([]theory.SimplifiedKey); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListScalePitches mock theory.Service#ListScalePitches
func (m *theoryService) ListScalePitches(ctx context.Context, scaleID int64) ([]theory.SimplifiedPitch, error) {
	args := m.Called(ctx, scaleID)

	var entries []theory.SimplifiedPitch
	if v, ok := args.Get(0).([]theory.SimplifiedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListScaleChords mock theory.Service#ListScaleChords
func (m *theoryService) ListScaleChords(ctx context.Context, scaleID int64, filter theory.ChordFilter, pagination api.Pagination) ([]theory.SimplifiedChord, *api.Pagination, error) {
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

// GetScale mock theory.Service#GetScale
func (m *theoryService) GetScale(ctx context.Context, scaleID int64) (*theory.DetailedScale, error) {
	args := m.Called(ctx, scaleID)

	var entry *theory.DetailedScale
	if v, ok := args.Get(0).(*theory.DetailedScale); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// ListKeys mock theory.Service#ListKeys
func (m *theoryService) ListKeys(ctx context.Context, filter theory.KeyFilter, pagination api.Pagination) ([]theory.SimplifiedKey, *api.Pagination, error) {
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

// ListKeyModes mock theory.Service#ListKeyModes
func (m *theoryService) ListKeyModes(ctx context.Context, keyID int64, filter theory.KeyFilter) ([]theory.SimplifiedKey, error) {
	args := m.Called(ctx, keyID, filter)

	var entries []theory.SimplifiedKey
	if v, ok := args.Get(0).([]theory.SimplifiedKey); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListKeyChords mock theory.Service#ListKeyChords
func (m *theoryService) ListKeyChords(ctx context.Context, keyID int64, filter theory.ChordFilter, pagination api.Pagination) ([]theory.SimplifiedChord, *api.Pagination, error) {
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

// ListKeyPitches mock theory.Service#ListKeyPitches
func (m *theoryService) ListKeyPitches(ctx context.Context, keyID int64) ([]theory.SimplifiedPitch, error) {
	args := m.Called(ctx, keyID)

	var entries []theory.SimplifiedPitch
	if v, ok := args.Get(0).([]theory.SimplifiedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// GetKey mock theory.Service#GetKey
func (m *theoryService) GetKey(ctx context.Context, keyID int64) (*theory.DetailedKey, error) {
	args := m.Called(ctx, keyID)

	var entry *theory.DetailedKey
	if v, ok := args.Get(0).(*theory.DetailedKey); ok {
		entry = v
	}

	return entry, args.Error(1)
}
