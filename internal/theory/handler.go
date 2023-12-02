package theory

import (
	"github.com/edipermadi/music-db/internal/platform/handler"
	"github.com/edipermadi/music-db/pkg/midi"
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
func NewHandler(logger *zap.Logger, service Service, synthesizerFactory midi.SynthesizerFactory) Handler {
	return theoryHandler{
		Base:               handler.NewBase(logger),
		service:            service,
		decoder:            form.NewDecoder(),
		synthesizerFactory: synthesizerFactory,
	}
}

type theoryHandler struct {
	handler.Base
	service            Service
	decoder            *form.Decoder
	synthesizerFactory midi.SynthesizerFactory
}

func (h theoryHandler) InstallEndpoints(router *mux.Router) {
	h.installChordEndpoints(router)
	h.installKeyEndpoints(router)
	h.installPitchEndpoints(router)
	h.installScaleEndpoints(router)
}
