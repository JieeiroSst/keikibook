package database

import (
	"database/sql"
	"log"

	"github.com/go-kit/kit/metrics"
)

type Handler struct {
	db     *sql.DB
	dur    *metrics.Histogram
	logger *log.Logger
}

func NewHandler(db *sql.DB, repuestDuration *metrics.Histogram, logger *log.Logger) *Handler {
	return &Handler{
		db:     db,
		dur:    repuestDuration,
		logger: logger,
	}
}
