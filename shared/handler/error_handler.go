package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// HTTPBaseHandler base handler struct.
// HTTPArticleHandler struct.
type HTTPErrorHandler struct {
	HTTPBaseHandler
	logger *logrus.Logger
}

// StatusNotFound responses status code 404.
func (h *HTTPErrorHandler) StatusNotFound(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("not found error.")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	// status code 404
	w.WriteHeader(http.StatusNotFound)
}

// StatusNotFound responses status code 405.
func (h *HTTPErrorHandler) StatusMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("method not allowed.")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	// status code 405
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// NewHTTPErrorHandler responses new HTTPArticleHandler instance.
func NewHTTPErrorHandler(logger *logrus.Logger) *HTTPErrorHandler {
	return &HTTPErrorHandler{logger: logger}
}
