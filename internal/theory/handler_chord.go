package theory

import (
	"bytes"
	"errors"
	"fmt"
	"image/png"
	"io"
	"math"
	"net/http"
	"strconv"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/pkg/illustations"
	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/gorilla/mux"
	"github.com/youpy/go-wav"
	"go.uber.org/zap"
)

type chordHandlers interface {
	GetChord(writer http.ResponseWriter, request *http.Request)
	GetChordQuality(writer http.ResponseWriter, request *http.Request)
	ListChordKeys(writer http.ResponseWriter, request *http.Request)
	ListChordPitches(writer http.ResponseWriter, request *http.Request)
	ListChordScales(writer http.ResponseWriter, request *http.Request)
	ListChords(writer http.ResponseWriter, request *http.Request)
}

func (h theoryHandler) installChordEndpoints(router *mux.Router) {
	router.HandleFunc("/chords", h.ListChords).Methods(http.MethodGet).Name("LIST_CHORDS")
	router.HandleFunc("/chords/{id:[0-9]+}", h.GetChord).Methods(http.MethodGet).Name("GET_CHORD")
	router.HandleFunc("/chords/{id:[0-9]+}/keys", h.ListChordKeys).Methods(http.MethodGet).Name("LIST_CHORD_KEYS")
	router.HandleFunc("/chords/{id:[0-9]+}/pitches", h.ListChordPitches).Methods(http.MethodGet).Name("GET_CHORD_PITCHES")
	router.HandleFunc("/chords/{id:[0-9]+}/quality", h.GetChordQuality).Methods(http.MethodGet).Name("GET_CHORD_QUALITY")
	router.HandleFunc("/chords/{id:[0-9]+}/scales", h.ListChordScales).Methods(http.MethodGet).Name("LIST_CHORD_SCALES")
	router.HandleFunc("/chords/{id:[0-9]+}/illustrations/keyboard", h.IllustrateChordWithKeyboard).Methods(http.MethodGet).Name("ILLUSTRATE_CHORD_WITH_KEYBOARD")
	router.HandleFunc("/chords/{id:[0-9]+}/illustrations/wav", h.IllustrateChordAsWavFile).Methods(http.MethodGet).Name("ILLUSTRATE_CHORD_AS_WAVE_FILE")
}

func (h theoryHandler) ListChords(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	type params struct {
		ChordFilter
		api.Pagination
	}

	var data params
	if err := h.decoder.Decode(&data, request.URL.Query()); err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list chords")
		h.ReplyJSON(writer, http.StatusBadRequest, api.ErrBadQueryParameter)
		return
	}

	chords, paginationOut, err := h.service.ListChords(ctx, data.ChordFilter, data.Pagination)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list chords")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.SetPagination(writer, paginationOut)
		h.ReplyJSON(writer, http.StatusOK, chords)
	}
}

func (h theoryHandler) ListChordPitches(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	chordID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	simplifiedPitches, err := h.service.ListChordPitches(ctx, chordID)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list chord pitches")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.ReplyJSON(writer, http.StatusOK, simplifiedPitches)
	}
}

func (h theoryHandler) ListChordKeys(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	type params struct {
		KeyFilter
		api.Pagination
	}

	var data params
	if err := h.decoder.Decode(&data, request.URL.Query()); err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list chord keys")
		h.ReplyJSON(writer, http.StatusBadRequest, api.ErrBadQueryParameter)
		return
	}

	chordID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	keys, paginationOut, err := h.service.ListChordKeys(ctx, chordID, data.KeyFilter, data.Pagination)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list chord keys")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.SetPagination(writer, paginationOut)
		h.ReplyJSON(writer, http.StatusOK, keys)
	}
}

func (h theoryHandler) ListChordScales(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	type params struct {
		ScaleFilter
		api.Pagination
	}

	var data params
	if err := h.decoder.Decode(&data, request.URL.Query()); err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list chord scales")
		h.ReplyJSON(writer, http.StatusBadRequest, api.ErrBadQueryParameter)
		return
	}

	chordID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	scales, paginationOut, err := h.service.ListChordScales(ctx, chordID, data.ScaleFilter, data.Pagination)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list chord scales")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.SetPagination(writer, paginationOut)
		h.ReplyJSON(writer, http.StatusOK, scales)
	}
}

func (h theoryHandler) GetChord(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	chordID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	chord, err := h.service.GetChord(ctx, chordID)
	switch {
	case errors.Is(err, ErrChordNotFound):
		h.ReplyJSON(writer, http.StatusNotFound, api.ErrResourceNotFound)
	case err != nil:
		h.Logger().With(zap.Error(err)).Error("failed to get chord")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	default:
		h.ReplyJSON(writer, http.StatusOK, chord)
	}
}

func (h theoryHandler) GetChordQuality(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	chordID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	quality, err := h.service.GetChordQuality(ctx, chordID)
	switch {
	case errors.Is(err, ErrChordQualityNotFound):
		h.ReplyJSON(writer, http.StatusNotFound, api.ErrResourceNotFound)
	case err != nil:
		h.Logger().With(zap.Error(err)).Error("failed to get chord quality")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	default:
		h.ReplyJSON(writer, http.StatusOK, quality)
	}
}

func (h theoryHandler) IllustrateChordWithKeyboard(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	chordID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)

	// get chord
	chord, err := h.service.GetChord(ctx, chordID)
	switch {
	case errors.Is(err, ErrChordNotFound):
		h.ReplyJSON(writer, http.StatusNotFound, api.ErrResourceNotFound)
		return
	case err != nil:
		h.Logger().With(zap.Error(err)).Error("failed to list chord pitches")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
		return
	}

	// list pitches
	simplifiedPitches, err := h.service.ListChordPitches(ctx, chordID)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list chord pitches")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
		return
	}

	pitches := make([]pitch.Type, 0)
	for _, v := range simplifiedPitches {
		pitches = append(pitches, pitch.FromInt(int(v.ID)))
	}

	// draw keyboard illustration
	img, err := illustations.Keyboard(pitches)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to draw keyboard illustration for chord")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "image/png")
	writer.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%q", fmt.Sprintf("%sKeyboard.png", chord.Name)))
	_ = png.Encode(writer, img)
}

func (h theoryHandler) IllustrateChordAsWavFile(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	chordID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)

	// get chord
	chord, err := h.service.GetChord(ctx, chordID)
	switch {
	case errors.Is(err, ErrChordNotFound):
		h.ReplyJSON(writer, http.StatusNotFound, api.ErrResourceNotFound)
		return
	case err != nil:
		h.Logger().With(zap.Error(err)).Error("failed to list chord pitches")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
		return
	}

	// list pitches
	simplifiedPitches, err := h.service.ListChordPitches(ctx, chordID)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list chord pitches")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
		return
	}

	// instantiate synthesizer
	sampleRate := 44100
	synthesizer, err := h.synthesizerFactory.Instantiate(int32(sampleRate))
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to instantiate synthesizer")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
		return
	}

	for _, v := range simplifiedPitches {
		p := pitch.FromInt(int(v.ID))
		if p >= pitch.CNatural && p < pitch.BNatural {
			synthesizer.NoteOn(0, int32(p)+59, 100)
		}
	}

	numSamples := 3 * sampleRate
	left := make([]float32, numSamples)
	right := make([]float32, numSamples)

	// render
	synthesizer.Render(left, right)

	// determine peak amplitude
	var maxValue float64
	for i := 0; i < numSamples; i++ {
		absLeft := math.Abs(float64(left[i]))
		absRight := math.Abs(float64(right[i]))
		if maxValue < absLeft {
			maxValue = absLeft
		}
		if maxValue < absRight {
			maxValue = absRight
		}
	}

	// convert to integer relative to amplitude
	a := 32768 * float32(0.99/maxValue)
	samples := make([]wav.Sample, numSamples)
	for i := 0; i < numSamples; i++ {
		samples[i].Values[0] = int(a * left[i])
		samples[i].Values[1] = int(a * right[i])
	}

	var buff bytes.Buffer
	encoder := wav.NewWriter(&buff, uint32(numSamples), 2, 44100, 16)
	if err := encoder.WriteSamples(samples); err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to convert PCM to wav file")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "audio/wav")
	writer.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%q", fmt.Sprintf("%sPiano.wav", chord.Name)))
	_, _ = io.Copy(writer, bytes.NewReader(buff.Bytes()))
}
