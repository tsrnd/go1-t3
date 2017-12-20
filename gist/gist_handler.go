package gist

import (
	"net/http"

	"github.com/monstar-lab/fr-circle-api/shared/handler"
)

type HTTPGistHandler struct {
	handler.HTTPBaseHandler
	GUsecase GistUsecase
}

// ListGists call usecase to get git urls.
func (h *HTTPGistHandler) ListGists(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	id := queryValues.Get("id")
	if id == "" {
		// default value
		id = "DaisukeHirata"
	}

	list, err := h.GUsecase.ListGists(id)

	if err != nil || len(list) == 0 {
		list = append(list, "404 not found")
		h.StatusNotFoundRequest(w, list)
		return
	}

	h.ResponseJSON(w, list)
}

func NewGistHTTPHandler(bh *handler.HTTPBaseHandler, gu GistUsecase) *HTTPGistHandler {
	return &HTTPGistHandler{
		HTTPBaseHandler: *bh,
		GUsecase:        gu,
	}
}
