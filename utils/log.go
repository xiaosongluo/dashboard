package utils

import (
	"log"
	"net/http"
	"time"
)

var (
	version string
)

type LogResponseWriter struct {
	Status int
	Size   int
	http.ResponseWriter
}

// NewLogResponseWriter instanciates a new LogResponseWriter
func NewLogResponseWriter(res http.ResponseWriter) *LogResponseWriter {
	// Default the status code to 200
	return &LogResponseWriter{200, 0, res}
}

// Header returns & satisfies the http.ResponseWriter interface
func (w *LogResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

// Write satisfies the http.ResponseWriter interface and
// captures data written, in bytes
func (w *LogResponseWriter) Write(data []byte) (int, error) {

	written, err := w.ResponseWriter.Write(data)
	w.Size += written

	return written, err
}

// WriteHeader satisfies the http.ResponseWriter interface and
// allows us to cach the status code
func (w *LogResponseWriter) WriteHeader(statusCode int) {

	w.Status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func LogHTTPRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, r *http.Request) {
		start := time.Now().UnixNano()
		w := NewLogResponseWriter(res)

		w.Header().Set("X-Application-Version", version)

		h.ServeHTTP(w, r)

		d := (time.Now().UnixNano() - start) / 1000
		log.Printf("%s %s %d %dµs %dB\n", r.Method, r.URL.Path, w.Status, d, w.Size)
	})
}
