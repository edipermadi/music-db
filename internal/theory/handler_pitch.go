package theory

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type pitchHandlers interface {
	GetPitch(writer http.ResponseWriter, request *http.Request)
	ListPitchChords(writer http.ResponseWriter, request *http.Request)
	ListPitchKeys(writer http.ResponseWriter, request *http.Request)
	ListPitchScales(writer http.ResponseWriter, request *http.Request)
	ListPitches(writer http.ResponseWriter, request *http.Request)
}

func (h theoryHandler) installPitchEndpoints(router *mux.Router) {
	router.HandleFunc("/pitches", h.ListPitches).Methods(http.MethodGet).Name("LIST_PITCHES")
	router.HandleFunc("/pitches/{id:[0-9]+}", h.GetPitch).Methods(http.MethodGet).Name("GET_PITCH")
	router.HandleFunc("/pitches/{id:[0-9]+}/chords", h.ListPitchChords).Methods(http.MethodGet).Name("LIST_PITCH_CHORDS")
	router.HandleFunc("/pitches/{id:[0-9]+}/scales", h.ListPitchScales).Methods(http.MethodGet).Name("LIST_PITCH_SCALES")
	router.HandleFunc("/pitches/{id:[0-9]+}/keys", h.ListPitchKeys).Methods(http.MethodGet).Name("LIST_PITCH_KEYS")
}

func (h theoryHandler) GetPitch(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	pitchID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	pitch, err := h.service.GetPitch(ctx, pitchID)
	switch {
	case errors.Is(err, ErrPitchNotFound):
		h.ReplyJSON(writer, http.StatusNotFound, api.ErrResourceNotFound)
	case err != nil:
		h.Logger().With(zap.Error(err)).Error("failed to get pitch")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	default:
		h.ReplyJSON(writer, http.StatusOK, pitch)
	}
}

func (h theoryHandler) ListPitchChords(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	type params struct {
		ChordFilter
		api.Pagination
	}

	var data params
	if err := h.decoder.Decode(&data, request.URL.Query()); err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list pitch chords")
		h.ReplyJSON(writer, http.StatusBadRequest, api.ErrBadQueryParameter)
		return
	}

	pitchID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	pitch, paginationOut, err := h.service.ListPitchChords(ctx, pitchID, data.ChordFilter, data.Pagination)
	switch {
	case err != nil:
		h.Logger().With(zap.Error(err)).Error("failed to list pitch chords")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	default:
		h.SetPagination(writer, paginationOut)
		h.ReplyJSON(writer, http.StatusOK, pitch)
	}
}

func (h theoryHandler) ListPitchKeys(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	type params struct {
		KeyFilter
		api.Pagination
	}

	var data params
	if err := h.decoder.Decode(&data, request.URL.Query()); err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list pitch keys")
		h.ReplyJSON(writer, http.StatusBadRequest, api.ErrBadQueryParameter)
		return
	}

	pitchID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	pitch, paginationOut, err := h.service.ListPitchKeys(ctx, pitchID, data.KeyFilter, data.Pagination)
	switch {
	case err != nil:
		h.Logger().With(zap.Error(err)).Error("failed to list pitch keys")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	default:
		h.SetPagination(writer, paginationOut)
		h.ReplyJSON(writer, http.StatusOK, pitch)
	}
}

func (h theoryHandler) ListPitchScales(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	type params struct {
		ScaleFilter
		api.Pagination
	}

	var data params
	if err := h.decoder.Decode(&data, request.URL.Query()); err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list pitch scales")
		h.ReplyJSON(writer, http.StatusBadRequest, api.ErrBadQueryParameter)
		return
	}

	pitchID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	pitch, paginationOut, err := h.service.ListPitchScales(ctx, pitchID, data.ScaleFilter, data.Pagination)
	switch {
	case err != nil:
		h.Logger().With(zap.Error(err)).Error("failed to list pitch scales")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	default:
		h.SetPagination(writer, paginationOut)
		h.ReplyJSON(writer, http.StatusOK, pitch)
	}
}

func (h theoryHandler) ListPitches(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	pitches, err := h.service.ListPitches(ctx)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list pitches")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.ReplyJSON(writer, http.StatusOK, pitches)
	}
}
