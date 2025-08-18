package handler

import "net/http"

type Handler struct {
	ctrl *metatada.Controller
}

func (h *Handler) GetMetadata(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	ctx := req.Context()
	m, err = h.ctrl.Get(ctx, id)
	return m
}
