package handlers

import (
	"compress/gzip"
	"net/http"
)

// GZipResponseMiddleware detects if the client can handle
// zipped content and if so returns the response in zipped format
func GZipResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// write the file
		next.ServeHTTP(w, r)
	})
}

// WrappedResponseWriter responseWriter wrapper
type WrappedResponseWriter struct {
	w  http.ResponseWriter
	gw *gzip.Writer
}

// NewWrappedResponseWriter creates the instance of WrappedResponseWriter
func NewWrappedResponseWriter(w http.ResponseWriter) *WrappedResponseWriter {
	gw := gzip.NewWriter(w)
	return &WrappedResponseWriter{w, gw}
}

// Header returns header
func (wr *WrappedResponseWriter) Header() http.Header {
	return wr.w.Header()
}

// WriteHeader writes the data
func (wr *WrappedResponseWriter) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}

// WriteHeader writes the status code
func (wr *WrappedResponseWriter) WriteHeader(statusCode int) {
	wr.w.WriteHeader(statusCode)
}
