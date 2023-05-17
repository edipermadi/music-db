package mock

import (
	"context"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/stretchr/testify/mock"
)

// ServiceReturnValues stores return value of mocked theory.Service
type ServiceReturnValues struct {
	GetPitch        []interface{}
	ListPitchChords []interface{}
	ListPitchKeys   []interface{}
	ListPitches     []interface{}

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

// TheoryService mock theory.Service
type TheoryService struct {
	mock.Mock
}

// ListPitches mock theory.Service#ListPitches
func (m *TheoryService) ListPitches(ctx context.Context) ([]theory.DetailedPitch, error) {
	args := m.Called(ctx)

	var entries []theory.DetailedPitch
	if v, ok := args.Get(0).([]theory.DetailedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// GetPitch mock theory.Service#GetPitch
func (m *TheoryService) GetPitch(ctx context.Context, pitchID int64) (*theory.DetailedPitch, error) {
	args := m.Called(ctx, pitchID)

	var entry *theory.DetailedPitch
	if v, ok := args.Get(0).(*theory.DetailedPitch); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// ListPitchChords mock theory.Service#ListPitchChords
func (m *TheoryService) ListPitchChords(ctx context.Context, pitchID int64, filter theory.ChordFilter, pagination api.Pagination) ([]theory.SimplifiedChord, *api.Pagination, error) {
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
func (m *TheoryService) ListPitchKeys(ctx context.Context, pitchID int64, filter theory.KeyFilter, pagination api.Pagination) ([]theory.SimplifiedKey, *api.Pagination, error) {
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

// ListChords mock theory.Service#ListChords
func (m *TheoryService) ListChords(ctx context.Context, filter theory.ChordFilter, pagination api.Pagination) ([]theory.DetailedChord, *api.Pagination, error) {
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

// ListChordPitches mock theory.Service#ListChordPitches
func (m *TheoryService) ListChordPitches(ctx context.Context, chordID int64) ([]theory.DetailedPitch, error) {
	args := m.Called(ctx, chordID)

	var entries []theory.DetailedPitch
	if v, ok := args.Get(0).([]theory.DetailedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListChordKeys mock theory.Service#ListChordKeys
func (m *TheoryService) ListChordKeys(ctx context.Context, chordID int64, filter theory.KeyFilter, pagination api.Pagination) ([]theory.DetailedKey, *api.Pagination, error) {
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

// GetChord mock theory.Service#GetChord
func (m *TheoryService) GetChord(ctx context.Context, chordID int64) (*theory.DetailedChord, error) {
	args := m.Called(ctx, chordID)

	var entry *theory.DetailedChord
	if v, ok := args.Get(0).(*theory.DetailedChord); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// GetChordQuality mock theory.Service#GetChordQuality
func (m *TheoryService) GetChordQuality(ctx context.Context, chordID int64) (*theory.DetailedChordQuality, error) {
	args := m.Called(ctx, chordID)

	var entry *theory.DetailedChordQuality
	if v, ok := args.Get(0).(*theory.DetailedChordQuality); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// ListScales mock theory.Service#ListScales
func (m *TheoryService) ListScales(ctx context.Context, pagination api.Pagination) ([]theory.DetailedScale, *api.Pagination, error) {
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

// ListScaleKeys mock theory.Service#ListScaleKeys
func (m *TheoryService) ListScaleKeys(ctx context.Context, scaleID int64) ([]theory.DetailedKey, error) {
	args := m.Called(ctx, scaleID)

	var entries []theory.DetailedKey
	if v, ok := args.Get(0).([]theory.DetailedKey); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// GetScale mock theory.Service#GetScale
func (m *TheoryService) GetScale(ctx context.Context, scaleID int64) (*theory.DetailedScale, error) {
	args := m.Called(ctx, scaleID)

	var entry *theory.DetailedScale
	if v, ok := args.Get(0).(*theory.DetailedScale); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// ListKeys mock theory.Service#ListKeys
func (m *TheoryService) ListKeys(ctx context.Context, filter theory.KeyFilter, pagination api.Pagination) ([]theory.DetailedKey, *api.Pagination, error) {
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

// ListKeyModes mock theory.Service#ListKeyModes
func (m *TheoryService) ListKeyModes(ctx context.Context, keyID int64, filter theory.KeyFilter) ([]theory.DetailedKey, error) {
	args := m.Called(ctx, keyID, filter)

	var entries []theory.DetailedKey
	if v, ok := args.Get(0).([]theory.DetailedKey); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListKeyChords mock theory.Service#ListKeyChords
func (m *TheoryService) ListKeyChords(ctx context.Context, keyID int64, filter theory.ChordFilter, pagination api.Pagination) ([]theory.DetailedChord, *api.Pagination, error) {
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

// ListKeyPitches mock theory.Service#ListKeyPitches
func (m *TheoryService) ListKeyPitches(ctx context.Context, keyID int64) ([]theory.DetailedPitch, error) {
	args := m.Called(ctx, keyID)

	var entries []theory.DetailedPitch
	if v, ok := args.Get(0).([]theory.DetailedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// GetKey mock theory.Service#GetKey
func (m *TheoryService) GetKey(ctx context.Context, keyID int64) (*theory.DetailedKey, error) {
	args := m.Called(ctx, keyID)

	var entry *theory.DetailedKey
	if v, ok := args.Get(0).(*theory.DetailedKey); ok {
		entry = v
	}

	return entry, args.Error(1)
}
