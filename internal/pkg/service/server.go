package service

import (
	"context"
	"io/fs"
	"log/slog"
	"net/http"
	"text/template"

	"github.com/marvinjwendt/httb/assets"
	"github.com/marvinjwendt/httb/internal/pkg/api"
	sloghttp "github.com/samber/slog-http"
)

func (s *Service) Start(ctx context.Context) error {
	slog.Info("starting httb service")

	service, err := NewService(s.config)
	if err != nil {
		return err
	}
	r := http.NewServeMux()
	h := api.HandlerFromMux(service, r)

	// landing page
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		_, _ = w.Write(assets.LandingPage)
	})

	// publish swagger ui under /docs
	swaggerUI, err := fs.Sub(assets.Swagger, "swagger-ui")
	if err != nil {
		return err
	}
	r.Handle("/docs/", http.StripPrefix("/docs", http.FileServer(http.FS(swaggerUI))))

	// parse openapi spec
	tmpl, err := template.New("openapi").Parse(assets.OpenAPISpec)
	// publish openapi spec
	openapiSpecHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")
		err := tmpl.Execute(w, *s.config)
		if err != nil {
			slog.Error("failed to render openapi spec", "error", err)
		}
	}
	r.HandleFunc("/openapi.yaml", openapiSpecHandler)
	r.HandleFunc("/openapi.yml", openapiSpecHandler)

	// middlewares
	h = DelayMiddleware(h)
	h = sloghttp.Recovery(h)
	h = sloghttp.New(slog.Default())(h)

	s.server = &http.Server{
		Handler: h,
		Addr:    s.config.Addr,
	}

	// Start server in a goroutine
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server stopped", "error", err)
		}
	}()

	// Wait for context cancellation
	<-ctx.Done()
	slog.Info("shutdown signal received, stopping server")

	return nil
}
