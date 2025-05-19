package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type calcMiddleware struct {
	logger *logrus.Logger
}

func New(logger *logrus.Logger) *calcMiddleware {
	return &calcMiddleware{
		logger: logger,
	}
}

type responseWriter struct {
	http.ResponseWriter
	code int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.code = code
	rw.ResponseWriter.WriteHeader(code)
}

func (c calcMiddleware) WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{ResponseWriter: w, code: http.StatusOK}

		timeStart := time.Now()
		next.ServeHTTP(rw, r)
		code := rw.code

		if code >= http.StatusBadRequest && code <= http.StatusNetworkAuthenticationRequired {
			c.logger.Errorf("%d %s %s %s", code, r.Method, r.RequestURI, time.Since(timeStart))
		} else {
			c.logger.Infof("%d %s %s %s", code, r.Method, r.RequestURI, time.Since(timeStart))
		}
	}
}

func (c calcMiddleware) WithAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func (c calcMiddleware) WithRecover(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				c.logger.Errorf("Recovered: %v", r)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}
}
