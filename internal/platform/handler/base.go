package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/edipermadi/music-db/internal/platform/api"
	"go.uber.org/zap"
)

// Base is base handler
type Base struct {
	logger *zap.Logger
}

// NewBase returns base handler
func NewBase(logger *zap.Logger) Base {
	return Base{logger: logger}
}

// Logger returns logger
func (b Base) Logger() *zap.Logger {
	return b.logger
}

// ReplyJSON reply request with JSON
func (b Base) ReplyJSON(writer http.ResponseWriter, status int, data interface{}) {
	writer.WriteHeader(status)
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		b.logger.With(zap.Error(err)).Error("failed to reply json")
	}
}

// SetPagination set pagination
func (b Base) SetPagination(writer http.ResponseWriter, pagination *api.Pagination) {
	if pagination != nil && pagination.NextPage > 0 {
		writer.Header().Set("X-Next-Page", strconv.Itoa(pagination.NextPage))
		writer.Header().Set("X-Total-Pages", strconv.Itoa(pagination.TotalPages))
		writer.Header().Set("X-Total-Items", strconv.Itoa(pagination.TotalItems))
	}
}
