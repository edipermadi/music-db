package mock

import (
	"context"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/stretchr/testify/mock"
)

// RepositoryReturnValues stores return value of mocked theory.Repository
type RepositoryReturnValues struct {
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

// TheoryRepository mock theory.Repository
type TheoryRepository struct {
	mock.Mock
}

// ListPitches mock theory.Repository#ListPitches
func (m *TheoryRepository) ListPitches(ctx context.Context) ([]theory.DetailedPitch, error) {
	args := m.Called(ctx)

	var entries []theory.DetailedPitch
	if v, ok := args.Get(0).([]theory.DetailedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// GetPitch mock theory.Repository#GetPitch
func (m *TheoryRepository) GetPitch(ctx context.Context, pitchID int64) (*theory.DetailedPitch, error) {
	args := m.Called(ctx, pitchID)

	var entry *theory.DetailedPitch
	if v, ok := args.Get(0).(*theory.DetailedPitch); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// ListPitchChords mock theory.Repository#SimplifiedChord
func (m *TheoryRepository) ListPitchChords(ctx context.Context, pitchID int64, filter theory.ChordFilter, pagination api.Pagination) ([]theory.SimplifiedChord, *api.Pagination, error) {
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
func (m *TheoryRepository) ListPitchKeys(ctx context.Context, pitchID int64, filter theory.KeyFilter, pagination api.Pagination) ([]theory.SimplifiedKey, *api.Pagination, error) {
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

// ListChords mock theory.Repository#ListChords
func (m *TheoryRepository) ListChords(ctx context.Context, filter theory.ChordFilter, pagination api.Pagination) ([]theory.DetailedChord, *api.Pagination, error) {
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

// ListChordPitches mock theory.Repository#ListChordPitches
func (m *TheoryRepository) ListChordPitches(ctx context.Context, chordID int64) ([]theory.DetailedPitch, error) {
	args := m.Called(ctx, chordID)

	var entries []theory.DetailedPitch
	if v, ok := args.Get(0).([]theory.DetailedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListChordKeys mock theory.Repository#ListChordKeys
func (m *TheoryRepository) ListChordKeys(ctx context.Context, chordID int64, filter theory.KeyFilter, pagination api.Pagination) ([]theory.DetailedKey, *api.Pagination, error) {
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

// GetChord mock theory.Repository#GetChord
func (m *TheoryRepository) GetChord(ctx context.Context, chordID int64) (*theory.DetailedChord, error) {
	args := m.Called(ctx, chordID)

	var entry *theory.DetailedChord
	if v, ok := args.Get(0).(*theory.DetailedChord); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// GetChordQuality mock theory.Repository#GetChordQuality
func (m *TheoryRepository) GetChordQuality(ctx context.Context, chordID int64) (*theory.DetailedChordQuality, error) {
	args := m.Called(ctx, chordID)

	var entry *theory.DetailedChordQuality
	if v, ok := args.Get(0).(*theory.DetailedChordQuality); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// ListScales mock theory.Repository#ListScales
func (m *TheoryRepository) ListScales(ctx context.Context, pagination api.Pagination) ([]theory.DetailedScale, *api.Pagination, error) {
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

// ListScaleKeys mock theory.Repository#ListScaleKeys
func (m *TheoryRepository) ListScaleKeys(ctx context.Context, scaleID int64) ([]theory.DetailedKey, error) {
	args := m.Called(ctx, scaleID)

	var entries []theory.DetailedKey
	if v, ok := args.Get(0).([]theory.DetailedKey); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// GetScale mock theory.Repository#GetScale
func (m *TheoryRepository) GetScale(ctx context.Context, scaleID int64) (*theory.DetailedScale, error) {
	args := m.Called(ctx, scaleID)

	var entry *theory.DetailedScale
	if v, ok := args.Get(0).(*theory.DetailedScale); ok {
		entry = v
	}

	return entry, args.Error(1)
}

// ListKeys mock theory.Repository#ListKeys
func (m *TheoryRepository) ListKeys(ctx context.Context, filter theory.KeyFilter, pagination api.Pagination) ([]theory.DetailedKey, *api.Pagination, error) {
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

// ListKeyModes mock theory.Repository#ListKeyModes
func (m *TheoryRepository) ListKeyModes(ctx context.Context, keyID int64, filter theory.KeyFilter) ([]theory.DetailedKey, error) {
	args := m.Called(ctx, keyID, filter)

	var entries []theory.DetailedKey
	if v, ok := args.Get(0).([]theory.DetailedKey); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// ListKeyChords mock theory.Repository#ListKeyChords
func (m *TheoryRepository) ListKeyChords(ctx context.Context, keyID int64, filter theory.ChordFilter, pagination api.Pagination) ([]theory.DetailedChord, *api.Pagination, error) {
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

// ListKeyPitches mock theory.Repository#ListKeyPitches
func (m *TheoryRepository) ListKeyPitches(ctx context.Context, keyID int64) ([]theory.DetailedPitch, error) {
	args := m.Called(ctx, keyID)

	var entries []theory.DetailedPitch
	if v, ok := args.Get(0).([]theory.DetailedPitch); ok {
		entries = v
	}

	return entries, args.Error(1)
}

// GetKey mock theory.Repository#GetKey
func (m *TheoryRepository) GetKey(ctx context.Context, keyID int64) (*theory.DetailedKey, error) {
	args := m.Called(ctx, keyID)

	var entry *theory.DetailedKey
	if v, ok := args.Get(0).(*theory.DetailedKey); ok {
		entry = v
	}

	return entry, args.Error(1)
}
