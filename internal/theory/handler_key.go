package theory

import (
	"errors"
	"fmt"
	"github.com/edipermadi/music-db/pkg/illustations"
	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"image/png"
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

func (h theoryHandler) installKeyEndpoints(router *mux.Router) {
	router.HandleFunc("/keys", h.ListKeys).Methods(http.MethodGet).Name("LIST_KEYS")
	router.HandleFunc("/keys/{id:[0-9]+}", h.GetKey).Methods(http.MethodGet).Name("GET_KEY")
	router.HandleFunc("/keys/{id:[0-9]+}/chords", h.ListKeyChords).Methods(http.MethodGet).Name("LIST_KEY_CHORDS")
	router.HandleFunc("/keys/{id:[0-9]+}/modes", h.ListKeyModes).Methods(http.MethodGet).Name("LIST_KEY_MODES")
	router.HandleFunc("/keys/{id:[0-9]+}/pitches", h.ListKeyPitches).Methods(http.MethodGet).Name("LIST_KEY_PITCHES")
	router.HandleFunc("/keys/{id:[0-9]+}/illustrations/pitch_class_bracelet", h.IllustrateKeyAsPitchClassBraceletDiagram).Methods(http.MethodGet).Name("ILLUSTRATE_KEY_AS_PITCH_CLASSES_BRACELET")
	router.HandleFunc("/keys/{id:[0-9]+}/illustrations/circle_of_fifth_bracelet", h.IllustrateKeyAsCircleOfFifthBraceletDiagram).Methods(http.MethodGet).Name("ILLUSTRATE_KEY_AS_CIRCLE_OF_FIFTH_BRACELET")
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
		h.Logger().With(zap.Error(err)).Error("failed to list key modes")
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
		h.Logger().With(zap.Error(err)).Error("failed to list key chords")
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

func (h theoryHandler) IllustrateKeyAsPitchClassBraceletDiagram(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	// get key
	keyID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	key, err := h.service.GetKey(ctx, keyID)
	switch {
	case errors.Is(err, ErrKeyNotFound):
		h.ReplyJSON(writer, http.StatusNotFound, api.ErrResourceNotFound)
	case err != nil:
		h.Logger().With(zap.Error(err)).Error("failed to get key")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	}

	// list pitches
	simplifiedPitches, err := h.service.ListKeyPitches(ctx, keyID)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list key pitches")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	}

	pitches := make([]pitch.Type, 0)
	for _, v := range simplifiedPitches {
		pitches = append(pitches, pitch.FromInt(int(v.ID)))
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "image/png")
	writer.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%q", fmt.Sprintf("%sCircleOfFifthBracelet.png", key.Name)))
	_ = png.Encode(writer, illustations.PitchClassBracelet(pitches))
}

func (h theoryHandler) IllustrateKeyAsCircleOfFifthBraceletDiagram(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	// get key
	keyID, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	key, err := h.service.GetKey(ctx, keyID)
	switch {
	case errors.Is(err, ErrKeyNotFound):
		h.ReplyJSON(writer, http.StatusNotFound, api.ErrResourceNotFound)
	case err != nil:
		h.Logger().With(zap.Error(err)).Error("failed to get key")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	}

	// list pitches
	simplifiedPitches, err := h.service.ListKeyPitches(ctx, keyID)
	if err != nil {
		h.Logger().With(zap.Error(err)).Error("failed to list key pitches")
		h.ReplyJSON(writer, http.StatusInternalServerError, api.ErrInternalServer)
	}

	pitches := make([]pitch.Type, 0)
	for _, v := range simplifiedPitches {
		pitches = append(pitches, pitch.FromInt(int(v.ID)))
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "image/png")
	writer.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%q", fmt.Sprintf("%sCircleOfFifthBracelet.png", key.Name)))
	_ = png.Encode(writer, illustations.CircleOfFifthBracelet(pitches))
}
