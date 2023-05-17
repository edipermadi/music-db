package theory

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type chordHandlers interface {
	ListChords(writer http.ResponseWriter, request *http.Request)
	ListChordPitches(writer http.ResponseWriter, request *http.Request)
	ListChordKeys(writer http.ResponseWriter, request *http.Request)
	GetChord(writer http.ResponseWriter, request *http.Request)
	GetChordQuality(writer http.ResponseWriter, request *http.Request)
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
	chords, err := h.service.ListChordPitches(ctx, chordID)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list chord pitches")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.ReplyJSON(writer, http.StatusOK, chords)
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
		h.Logger().With(zap.Error(err)).Error("failed to list chords")
		h.ReplyJSON(writer, http.StatusBadRequest, api.ErrBadQueryParameter)
		return
	}

	chordID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	chords, paginationOut, err := h.service.ListChordKeys(ctx, chordID, data.KeyFilter, data.Pagination)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list chord keys")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.SetPagination(writer, paginationOut)
		h.ReplyJSON(writer, http.StatusOK, chords)
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
	chord, err := h.service.GetChordQuality(ctx, chordID)
	switch {
	case errors.Is(err, ErrChordQualityNotFound):
		h.ReplyJSON(writer, http.StatusNotFound, api.ErrResourceNotFound)
	case err != nil:
		h.Logger().With(zap.Error(err)).Error("failed to get chord quality")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	default:
		h.ReplyJSON(writer, http.StatusOK, chord)
	}
}
