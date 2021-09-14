package apiserver

import "net/http"

type responseWriter struct {
	// it is an anonymous field.
	http.ResponseWriter
	code int
}

func (w *responseWriter) WriterHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
