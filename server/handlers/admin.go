package handlers

import "net/http"

func (h *Handlers) AdminAuth(w http.ResponseWriter, r *http.Request) {
	h.infoLog.Println("hit the welcome handler")
	var payload jsonResponse
	payload.Error = false
	payload.Message = "Welcome to Web admin user!"
	h.writeJSON(w, http.StatusOK, payload)
}
