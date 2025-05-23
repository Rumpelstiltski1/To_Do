package metricks

import (
	"net/http"
	"strconv"
	"time"
)

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func MetricsMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		start := time.Now()

		rr := &responseRecorder{
			ResponseWriter: writer,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(rr, request)

		duration := time.Since(start).Seconds()

		path := request.URL.Path
		method := request.Method
		status := strconv.Itoa(rr.statusCode)

		HttpRequestTotal.WithLabelValues(method, path).Inc()
		HttpRequestDuration.WithLabelValues(method, path).Observe(duration)
		HttpResponseStatus.WithLabelValues(status).Inc()
	})
}

/*
Если в нашем коде происходит ошибка
и вы вызываете w.WriteHeader(500), сработает метод WriteHeader у responseRecorder,
который сохранит статус 500 в rr.statusCode.
*/
func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}
