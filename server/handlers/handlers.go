package handlers

import (
	"log"
	"net/http"
)

type Handlers struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	environment string
}

type jsonResponse struct {
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type envelope map[string]any

func NewHandlers(infoLog, errorLog *log.Logger, environment string) *Handlers {
	return &Handlers{
		infoLog:     infoLog,
		errorLog:    errorLog,
		environment: environment,
	}
}

// Welcome godoc
//	@Summary		Landing page for web app
//	@Description	An page that handles GET without authorization
//	@ID				welcome-landing-page
//	@Produce		json
//	@Router			/ [get]
//	@tags			mock
func (h *Handlers) Welcome(w http.ResponseWriter, r *http.Request) {
	h.infoLog.Println("hit the welcome handler")
	var payload jsonResponse
	payload.Error = false
	payload.Message = "Welcome to the Web Server!"
	h.writeJSON(w, http.StatusOK, payload)
}
