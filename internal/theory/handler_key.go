package theory

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type keyHandlers interface {
	ListKeys(writer http.ResponseWriter, request *http.Request)
	ListKeyModes(writer http.ResponseWriter, request *http.Request)
	ListKeyChords(writer http.ResponseWriter, request *http.Request)
	ListKeyPitches(writer http.ResponseWriter, request *http.Request)
	GetKey(writer http.ResponseWriter, request *http.Request)
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
