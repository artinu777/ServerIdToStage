package apiv1

import (
	"github.com/go-chi/chi"
	"github.com/sungora/sample/internal/middlep"
	"net/http"
)

func StageRoute() http.Handler  {
	r := chi.NewRouter()
	r.Use(middlep.SampleOne)
	r.Route("/", func(r chi.Router) {
		r.Use(middlep.SampleTwo)
		r.Get("/{id}/stage/{stage_id}", Stage1Handler)
	})
	return r
}