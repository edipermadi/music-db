package midi

type Synthesizer interface {
	ProcessMidiMessage(channel int32, command int32, data1 int32, data2 int32)
	NoteOff(channel int32, key int32)
	NoteOn(channel int32, key int32, velocity int32)
	NoteOffAll(immediate bool)
	NoteOffAllChannel(channel int32, immediate bool)
	ResetAllControllers()
	ResetAllControllersChannel(channel int32)
	Reset()
	Render(left []float32, right []float32)
}

type SynthesizerFactory interface {
	Instantiate(sampleRate int32) (Synthesizer, error)
}
