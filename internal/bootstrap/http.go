package bootstrap

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/bddjr/hlfhr"
	"github.com/go-chi/chi/v5"
	"github.com/knadh/koanf/v2"

	"github.com/TheTNB/panel/internal/http/middleware"
	"github.com/TheTNB/panel/internal/route"
)

func NewRouter(middlewares *middleware.Middlewares, http *route.Http, ws *route.Ws) (*chi.Mux, error) {
	r := chi.NewRouter()

	// add middleware
	r.Use(middlewares.Globals(r)...)
	// add http route
	http.Register(r)
	// add ws route
	ws.Register(r)

	return r, nil
}

func NewHttp(conf *koanf.Koanf, r *chi.Mux) (*hlfhr.Server, error) {
	srv := hlfhr.New(&http.Server{
		Addr:           fmt.Sprintf(":%d", conf.MustInt("http.port")),
		Handler:        http.AllowQuerySemicolons(r),
		MaxHeaderBytes: 2048 << 20,
	})
	srv.HttpOnHttpsPortErrorHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hlfhr.RedirectToHttps(w, r, http.StatusTemporaryRedirect)
	})

	if conf.Bool("http.tls") {
		srv.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS13,
		}
	}

	return srv, nil
}
