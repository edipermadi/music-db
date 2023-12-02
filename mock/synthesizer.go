package mock

import (
	"github.com/edipermadi/music-db/pkg/midi"
	"github.com/stretchr/testify/mock"
)

type SynthesizerFactory struct {
	mock.Mock
}

func (m *SynthesizerFactory) Instantiate(sampleRate int32) (midi.Synthesizer, error) {
	args := m.Called(sampleRate)
	var s midi.Synthesizer
	if v, ok := args.Get(0).(midi.Synthesizer); ok {
		s = v
	}

	return s, args.Error(1)
}
