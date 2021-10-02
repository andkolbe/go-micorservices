package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GzipHandler struct {
}

// gzip will reduce the size of the data by 1/3

func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// create a gziped response
			wrw := NewWrapperResponseWriter(rw)
			wrw.Header().Set("Content-Encoding", "gzip")

			next.ServeHTTP(wrw, r)
			defer wrw.Flush()
			return
		}

		next.ServeHTTP(rw, r)
	})
}

type WrappedResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

func NewWrapperResponseWriter(rw http.ResponseWriter) *WrappedResponseWriter {
	gw := gzip.NewWriter(rw)
	return &WrappedResponseWriter{rw: rw, gw: gw}
}

func (wr *WrappedResponseWriter) Header() http.Header {
	return wr.rw.Header()
}

func (wr *WrappedResponseWriter) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}

func (wr *WrappedResponseWriter) WriteHeader(statuscode int) {
	wr.rw.WriteHeader(statuscode)
}

func (wr *WrappedResponseWriter) Flush() {
	wr.gw.Flush()
	wr.gw.Close()
}
