package middleware

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
)

type Middleware struct {
	requests *prometheus.CounterVec
}

func NewPrometheus() func(next http.Handler) http.Handler {
	var m Middleware
	m.requests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests_total",
			Help: "Number of requests",
		},
		[]string{"status"},
	)
	prometheus.MustRegister(m.requests)

	return m.handler
}

func (c Middleware) handler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)
		c.requests.WithLabelValues(strconv.Itoa(ww.Status())).Inc()
	}
	return http.HandlerFunc(fn)
}
