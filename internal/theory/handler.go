package theory

import (
	"net/http"

	"github.com/edipermadi/music-db/internal/platform/handler"
	"github.com/go-playground/form/v4"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Handler is theory handler
type Handler interface {
	InstallEndpoints(router *mux.Router)

	chordHandlers
	keyHandlers
	pitchHandlers
	scaleHandlers
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
	router.HandleFunc("/pitches/{id:[0-9]+}", h.GetPitch).
		Methods(http.MethodGet).Name("GET_PITCH")
	router.HandleFunc("/pitches/{id:[0-9]+}/chords", h.ListPitchChords).
		Methods(http.MethodGet).Name("LIST_PITCH_CHORDS")
	router.HandleFunc("/pitches/{id:[0-9]+}/keys", h.ListPitchKeys).
		Methods(http.MethodGet).Name("LIST_PITCH_KEYS")

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
