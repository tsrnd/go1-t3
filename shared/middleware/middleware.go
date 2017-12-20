package middleware

import (
	"net/http"

	"github.com/monstar-lab/fr-circle-api/infrastructure"
	"github.com/sirupsen/logrus"
)

// Logger is log middleware.
func Logger(logger *infrastructure.LoggerHandler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// output headers.
			for i, v := range r.Header {
				logger.Log.WithFields(logrus.Fields{
					"header": i,
					"value": v,
				}).Info("")
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

// Header is log middleware.
func Header(logger *infrastructure.LoggerHandler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
