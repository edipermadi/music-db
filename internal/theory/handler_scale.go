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

type scaleHandlers interface {
	ListScales(writer http.ResponseWriter, request *http.Request)
	ListScaleKeys(writer http.ResponseWriter, request *http.Request)
	ListScaleChords(writer http.ResponseWriter, request *http.Request)
	ListScalePitches(writer http.ResponseWriter, request *http.Request)
	GetScale(writer http.ResponseWriter, request *http.Request)
}

func (h theoryHandler) installScaleEndpoints(router *mux.Router) {
	router.HandleFunc("/scales", h.ListScales).Methods(http.MethodGet).Name("LIST_SCALES")
	router.HandleFunc("/scales/{id:[0-9]+}", h.GetScale).Methods(http.MethodGet).Name("GET_SCALE")
	router.HandleFunc("/scales/{id:[0-9]+}/keys", h.ListScaleKeys).Methods(http.MethodGet).Name("LIST_SCALE_KEYS")
	router.HandleFunc("/scales/{id:[0-9]+}/chords", h.ListScaleChords).Methods(http.MethodGet).Name("LIST_SCALE_CHORDS")
	router.HandleFunc("/scales/{id:[0-9]+}/pitches", h.ListScalePitches).Methods(http.MethodGet).Name("LIST_SCALE_PITCHES")
	router.HandleFunc("/scales/{id:[0-9]+}/illustrations/pitch_class_bracelet", h.IllustrateScaleAsPitchClassBraceletDiagram).Methods(http.MethodGet).Name("ILLUSTRATE_SCALE_AS_PITCH_CLASS_BRACELET")
	router.HandleFunc("/scales/{id:[0-9]+}/illustrations/circle_of_fifth_bracelet", h.IllustrateScaleAsCircleOfFifthBraceletDiagram).Methods(http.MethodGet).Name("ILLUSTRATE_SCALE_AS_CIRCLE_OF_FIFTH_BRACELET")
}

func (h theoryHandler) ListScales(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	type params struct {
		ScaleFilter
		api.Pagination
	}

	var data params
	if err := h.decoder.Decode(&data, request.URL.Query()); err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list scales")
		h.ReplyJSON(writer, http.StatusBadRequest, api.ErrBadQueryParameter)
		return
	}

	scales, paginationOut, err := h.service.ListScales(ctx, data.ScaleFilter, data.Pagination)
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
		h.Logger().With(zap.Error(err)).Error("failed to list scale keys")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.ReplyJSON(writer, http.StatusOK, keys)
	}
}

func (h theoryHandler) ListScalePitches(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	scaleID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	pitches, err := h.service.ListScalePitches(ctx, scaleID)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list scale pitches")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.ReplyJSON(writer, http.StatusOK, pitches)
	}
}

func (h theoryHandler) ListScaleChords(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	type params struct {
		ChordFilter
		api.Pagination
	}

	var data params
	if err := h.decoder.Decode(&data, request.URL.Query()); err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list scale chords")
		h.ReplyJSON(writer, http.StatusBadRequest, api.ErrBadQueryParameter)
		return
	}

	scaleID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	keys, pagination, err := h.service.ListScaleChords(ctx, scaleID, data.ChordFilter, data.Pagination)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list scale chords")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	} else {
		h.SetPagination(writer, pagination)
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

func (h theoryHandler) IllustrateScaleAsPitchClassBraceletDiagram(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	scaleID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)

	scale, err := h.service.GetScale(ctx, scaleID)
	switch {
	case errors.Is(err, ErrScaleNotFound):
		h.ReplyJSON(writer, http.StatusNotFound, api.ErrResourceNotFound)
		return
	case err != nil:
		h.Logger().With(zap.Error(err)).Error("failed to get scale")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
		return
	}

	pitches := make([]pitch.Type, 0)
	for _, v := range scale.PitchClass {
		pitches = append(pitches, pitch.FromInt(v+1))
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "image/png")
	writer.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%q", fmt.Sprintf("%sCircleOfFifthBracelet.png", scale.Name)))
	_ = png.Encode(writer, illustations.PitchClassBracelet(pitches))
}

func (h theoryHandler) IllustrateScaleAsCircleOfFifthBraceletDiagram(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	scaleID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)

	scale, err := h.service.GetScale(ctx, scaleID)
	switch {
	case errors.Is(err, ErrScaleNotFound):
		h.ReplyJSON(writer, http.StatusNotFound, api.ErrResourceNotFound)
		return
	case err != nil:
		h.Logger().With(zap.Error(err)).Error("failed to get scale")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
		return
	}

	pitches := make([]pitch.Type, 0)
	for _, v := range scale.PitchClass {
		pitches = append(pitches, pitch.FromInt(v+1))
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "image/png")
	writer.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%q", fmt.Sprintf("%sCircleOfFifthBracelet.png", scale.Name)))
	_ = png.Encode(writer, illustations.CircleOfFifthBracelet(pitches))
}
