package slogKit

import (
	"io"
	"log/slog"
)

func NewTextLogger(w io.Writer, opts *slog.HandlerOptions) *slog.Logger {
	handler := slog.NewTextHandler(w, opts)
	return slog.New(handler)
}

func NewJsonLogger(w io.Writer, opts *slog.HandlerOptions) *slog.Logger {
	handler := slog.NewJSONHandler(w, opts)
	return slog.New(handler)
}
