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
	ListPitches(writer http.ResponseWriter, request *http.Request)

	ListChords(writer http.ResponseWriter, request *http.Request)

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

	var pagination api.Pagination
	if err := h.decoder.Decode(&pagination, request.URL.Query()); err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list chords")
		h.ReplyJSON(writer, http.StatusBadRequest, api.ErrBadQueryParameter)
		return
	}

	pagination.Sanitize()
	chords, paginationOut, err := h.service.ListChords(ctx, pagination)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list chords")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.SetPagination(writer, paginationOut)
		h.ReplyJSON(writer, http.StatusOK, chords)
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

	var pagination api.Pagination
	if err := h.decoder.Decode(&pagination, request.URL.Query()); err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list keys")
		h.ReplyJSON(writer, http.StatusBadRequest, api.ErrBadQueryParameter)
		return
	}

	pagination.Sanitize()
	keys, paginationOut, err := h.service.ListKeys(ctx, pagination)
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

	keyID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	keys, err := h.service.ListKeyModes(ctx, keyID)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list key modes")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.ReplyJSON(writer, http.StatusOK, keys)
	}
}

func (h theoryHandler) ListKeyChords(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var pagination api.Pagination
	if err := h.decoder.Decode(&pagination, request.URL.Query()); err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list keys")
		h.ReplyJSON(writer, http.StatusBadRequest, api.ErrBadQueryParameter)
		return
	}

	pagination.Sanitize()
	keyID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	keys, paginationOut, err := h.service.ListKeyChords(ctx, keyID, pagination)
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
	case errors.Is(err, ErrScaleNotFound):
		h.ReplyJSON(writer, http.StatusNotFound, api.ErrResourceNotFound)
	case err != nil:
		h.Logger().With(zap.Error(err)).Error("failed to get key")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	default:
		h.ReplyJSON(writer, http.StatusOK, key)
	}
}
