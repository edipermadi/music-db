package theory

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type scaleHandlers interface {
	ListScales(writer http.ResponseWriter, request *http.Request)
	ListScaleKeys(writer http.ResponseWriter, request *http.Request)
	GetScale(writer http.ResponseWriter, request *http.Request)
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
