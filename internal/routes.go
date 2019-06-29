package internal

import (
	"fmt"
	"github.com/sungora/sample/internal/stage/apiv1"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sungora/app/request"
	"github.com/sungora/app/servhttp/middles"
	"github.com/swaggo/http-swagger"

	"github.com/sungora/sample/docs"
	"github.com/sungora/sample/internal/config"
)

func Routes(router *chi.Mux, cfg *config.Config) {

	router.Use(middles.TimeoutContext(cfg.Http.WriteTimeout))
	router.Use(middles.Cors(cfg.Http.Cors).Handler)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	// general
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", cfg.App.Domain, cfg.Http.Port)
	router.Get("/api/v1/*", httpSwagger.Handler()) // swagger
	router.Get("/api/v1/ping", handlerPing)        // check server work

	router.Mount("/api/v1/server", apiv1.StageRoute())
}

// handlerPing проверка доступности сервера
// @Tags General
// @Router /ping [get]
// @Summary проверка доступности сервера
// @Success 200 {string} string
// @Failure 500 {string} string
func handlerPing(w http.ResponseWriter, r *http.Request) {
	request.NewIn(w, r).Json("pong")
}
