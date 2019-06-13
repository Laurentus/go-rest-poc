package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Laurentus/poc-app/dao"
	"github.com/Laurentus/poc-app/models"
	"github.com/go-chi/chi"
)

func CampaignRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", listCampaigns)
	return r
}

func listCampaigns(w http.ResponseWriter, r *http.Request) {
	ListCampaignsResponse(w.Write, dao.Campaigns().All())
}

func ListCampaignsResponse(write func([]byte) (int, error), future chan models.Campaign) {
	listResponse := make([]models.Campaign, 0)
	for camp := range future {
		listResponse = append(listResponse, camp)
	}
	bytes, err := json.Marshal(listResponse)
	if err != nil {
		log.Fatal(err)
	}
	write(bytes)
}
