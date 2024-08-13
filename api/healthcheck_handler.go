package api

import (
	"io"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "Text/Plain")
	_, _ = io.WriteString(w, "OK")
}
