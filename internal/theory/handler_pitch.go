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
	ListPitches(writer http.ResponseWriter, request *http.Request)
	GetPitch(writer http.ResponseWriter, request *http.Request)
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
