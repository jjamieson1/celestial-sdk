package utilities

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func CorrelationIdMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		logger := Logger(r.Context())
		defer logger.Sync()

		correlationId := r.Header.Get("X-Trace-Id")
		if correlationId == "" {
			correlationId = uuid.New().String()
		}

		r = r.WithContext(context.WithValue(r.Context(), CorrelationIdConst, correlationId))
		handler.ServeHTTP(rw, r)
	})
}
