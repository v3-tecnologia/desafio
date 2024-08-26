package api

import (
	"bytes"
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
	"time"
)

// key type for passing value through context
type key string

const (
	ctxKey key = "key"
)

const (
	noErrorContent = 0
)

type LogResponseWriter struct {
	http.ResponseWriter
	statusCode int
	buf        bytes.Buffer
}

func NewLogResponseWriter(w http.ResponseWriter) *LogResponseWriter {
	return &LogResponseWriter{ResponseWriter: w}
}

func (w *LogResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *LogResponseWriter) Write(body []byte) (int, error) {
	w.buf.Write(body)
	return w.ResponseWriter.Write(body)
}

//-------------------------------------------------------------------------------------

// Middleware to set a requestID.
func TraceIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create uuid for request
		uuid, err := uuid.NewUUID()
		if err != nil {
			logrus.Errorf("traceID middleware error: %v", err)
			http.Error(w, "", http.StatusInternalServerError)
		}

		ctx := context.WithValue(r.Context(), ctxKey, uuid.String())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//-------------------------------------------------------------------------------------

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Context().Value(ctxKey)
		start := time.Now()
		logrus.Infof("request started [%s] '%s' | TraceId: %v",
			r.Method,
			r.URL.Path,
			traceID,
		)

		lw := LogResponseWriter{ResponseWriter: w}
		next.ServeHTTP(&lw, r)

		if lw.statusCode < http.StatusOK || lw.statusCode > http.StatusAccepted {
			b := lw.buf.Bytes()
			if len(b) > noErrorContent {
				logrus.Warn(lw.buf.String())
			}
		}

		logrus.Infof("request finished [%s] '%s'| Status code: %v | TraceId: %v | Latency: %v",
			r.Method,
			r.URL.Path,
			lw.statusCode,
			traceID,
			time.Since(start),
		)
	})
}

//-------------------------------------------------------------------------------------

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check panic
		defer func() {
			if rec := recover(); rec != nil {
				trace := debug.Stack()
				logrus.Errorf("PANIC [%v] TRACE[%s]", rec, string(trace))
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
