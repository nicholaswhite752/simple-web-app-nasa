package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	h.infoLog.Println("hit the mock login handler")
	type credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var creds credentials
	var payload jsonResponse

	err := h.readJSON(w, r, &creds)
	if err != nil || creds.Username == "" || creds.Password == "" {
		h.errorLog.Println("Error:", err)
		payload.Error = true
		payload.Message = "json is invalid or missing"
		h.writeJSON(w, http.StatusBadRequest, payload)
		return
	}

	payload.Error = false
	payload.Message = fmt.Sprintf("Username: %s, Password: %s", creds.Username, creds.Password)
	h.writeJSON(w, http.StatusOK, payload)

}

func (h *Handlers) UserAuth(w http.ResponseWriter, r *http.Request) {
	h.infoLog.Println("hit the welcome handler")
	var payload jsonResponse
	payload.Error = false
	payload.Message = "Welcome to web app authorized user!"
	h.writeJSON(w, http.StatusOK, payload)
}
