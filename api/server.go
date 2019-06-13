package server

import (
	"net/http"

	"github.com/Laurentus/poc-app/api/routes"
	"github.com/Laurentus/poc-app/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type CampaignResponse = models.Campaign
type CampaignListResponse []CampaignResponse

func Start() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)

	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})
	r.Mount("/campaigns", routes.CampaignRouter())

	http.ListenAndServe(":3333", r)
}
