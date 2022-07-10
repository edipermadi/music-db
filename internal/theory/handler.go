package theory

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/platform/handler"
	"github.com/go-playground/form/v4"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Handler is theory handler
type Handler interface {
	InstallEndpoints(router *mux.Router)

	ListPitches(writer http.ResponseWriter, request *http.Request)

	ListChords(writer http.ResponseWriter, request *http.Request)
	ListChordPitches(writer http.ResponseWriter, request *http.Request)
	ListChordKeys(writer http.ResponseWriter, request *http.Request)
	GetChord(writer http.ResponseWriter, request *http.Request)
	GetChordQuality(writer http.ResponseWriter, request *http.Request)

	ListScales(writer http.ResponseWriter, request *http.Request)
	ListScaleKeys(writer http.ResponseWriter, request *http.Request)
	GetScale(writer http.ResponseWriter, request *http.Request)

	ListKeys(writer http.ResponseWriter, request *http.Request)
	ListKeyModes(writer http.ResponseWriter, request *http.Request)
	ListKeyChords(writer http.ResponseWriter, request *http.Request)
	ListKeyPitches(writer http.ResponseWriter, request *http.Request)
	GetKey(writer http.ResponseWriter, request *http.Request)
}

// NewHandler instantiates theory handler
func NewHandler(logger *zap.Logger, service Service) Handler {
	return theoryHandler{
		Base:    handler.NewBase(logger),
		service: service,
		decoder: form.NewDecoder(),
	}
}

type theoryHandler struct {
	handler.Base
	service Service
	decoder *form.Decoder
}

func (h theoryHandler) InstallEndpoints(router *mux.Router) {
	router.HandleFunc("/pitches", h.ListPitches).
		Methods(http.MethodGet).Name("LIST_PITCHES")

	router.HandleFunc("/chords", h.ListChords).
		Methods(http.MethodGet).Name("LIST_CHORDS")
	router.HandleFunc("/chords/{id:[0-9]+}", h.GetChord).
		Methods(http.MethodGet).Name("GET_CHORD")
	router.HandleFunc("/chords/{id:[0-9]+}/quality", h.GetChordQuality).
		Methods(http.MethodGet).Name("GET_CHORD_QUALITY")
	router.HandleFunc("/chords/{id:[0-9]+}/pitches", h.ListChordPitches).
		Methods(http.MethodGet).Name("GET_CHORD_PITCHES")
	router.HandleFunc("/chords/{id:[0-9]+}/keys", h.ListChordKeys).
		Methods(http.MethodGet).Name("GET_CHORD_KEYS")

	router.HandleFunc("/scales", h.ListScales).
		Methods(http.MethodGet).Name("LIST_SCALES")
	router.HandleFunc("/scales/{id:[0-9]+}", h.GetScale).
		Methods(http.MethodGet).Name("GET_SCALE")
	router.HandleFunc("/scales/{id:[0-9]+}/keys", h.ListScaleKeys).
		Methods(http.MethodGet).Name("LIST_KEYS_FROM_SCALE")

	router.HandleFunc("/keys", h.ListKeys).
		Methods(http.MethodGet).Name("LIST_KEYS")
	router.HandleFunc("/keys/{id:[0-9]+}/modes", h.ListKeyModes).
		Methods(http.MethodGet).Name("LIST_KEY_MODES")
	router.HandleFunc("/keys/{id:[0-9]+}/chords", h.ListKeyChords).
		Methods(http.MethodGet).Name("LIST_KEY_CHORDS")
	router.HandleFunc("/keys/{id:[0-9]+}/pitches", h.ListKeyPitches).
		Methods(http.MethodGet).Name("LIST_KEY_PITCHES")
	router.HandleFunc("/keys/{id:[0-9]+}", h.GetKey).
		Methods(http.MethodGet).Name("GET_KEY")
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

func (h theoryHandler) ListScales(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var pagination api.Pagination
	if err := h.decoder.Decode(&pagination, request.URL.Query()); err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list scales")
		h.ReplyJSON(writer, http.StatusBadRequest, api.ErrBadQueryParameter)
		return
	}

	pagination.Sanitize()
	scales, paginationOut, err := h.service.ListScales(ctx, pagination)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list scales")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.SetPagination(writer, paginationOut)
		h.ReplyJSON(writer, http.StatusOK, scales)
	}
}

func (h theoryHandler) ListScaleKeys(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	scaleID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	keys, err := h.service.ListScaleKeys(ctx, scaleID)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list keys from scale")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.ReplyJSON(writer, http.StatusOK, keys)
	}
}

func (h theoryHandler) GetScale(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	scaleID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	scale, err := h.service.GetScale(ctx, scaleID)
	switch {
	case errors.Is(err, ErrScaleNotFound):
		h.ReplyJSON(writer, http.StatusNotFound, api.ErrResourceNotFound)
	case err != nil:
		h.Logger().With(zap.Error(err)).Error("failed to get scale")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	default:
		h.ReplyJSON(writer, http.StatusOK, scale)
	}
}

func (h theoryHandler) ListKeys(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	type params struct {
		KeyFilter
		api.Pagination
	}

	var data params
	if err := h.decoder.Decode(&data, request.URL.Query()); err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list keys")
		h.ReplyJSON(writer, http.StatusBadRequest, api.ErrBadQueryParameter)
		return
	}

	keys, paginationOut, err := h.service.ListKeys(ctx, data.KeyFilter, data.Pagination)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list keys")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.SetPagination(writer, paginationOut)
		h.ReplyJSON(writer, http.StatusOK, keys)
	}
}

func (h theoryHandler) ListKeyModes(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var data KeyFilter
	if err := h.decoder.Decode(&data, request.URL.Query()); err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list modes")
		h.ReplyJSON(writer, http.StatusBadRequest, api.ErrBadQueryParameter)
		return
	}

	keyID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	keys, err := h.service.ListKeyModes(ctx, keyID, data)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list key modes")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.ReplyJSON(writer, http.StatusOK, keys)
	}
}

func (h theoryHandler) ListKeyChords(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	type params struct {
		ChordFilter
		api.Pagination
	}

	var data params
	if err := h.decoder.Decode(&data, request.URL.Query()); err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list keys")
		h.ReplyJSON(writer, http.StatusBadRequest, api.ErrBadQueryParameter)
		return
	}

	keyID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	keys, paginationOut, err := h.service.ListKeyChords(ctx, keyID, data.ChordFilter, data.Pagination)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list key chords")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.SetPagination(writer, paginationOut)
		h.ReplyJSON(writer, http.StatusOK, keys)
	}
}

func (h theoryHandler) ListKeyPitches(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	keyID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	pitches, err := h.service.ListKeyPitches(ctx, keyID)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list key pitches")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.ReplyJSON(writer, http.StatusOK, pitches)
	}
}

func (h theoryHandler) GetKey(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	keyID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	key, err := h.service.GetKey(ctx, keyID)
	switch {
	case errors.Is(err, ErrKeyNotFound):
		h.ReplyJSON(writer, http.StatusNotFound, api.ErrResourceNotFound)
	case err != nil:
		h.Logger().With(zap.Error(err)).Error("failed to get key")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	default:
		h.ReplyJSON(writer, http.StatusOK, key)
	}
}
