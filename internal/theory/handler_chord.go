package theory

import (
	"errors"
	"fmt"
	"image/png"
	"net/http"
	"strconv"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/pkg/illustations"
	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/gorilla/mux"
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
