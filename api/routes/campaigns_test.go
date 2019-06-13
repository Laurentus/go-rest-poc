package routes_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/Laurentus/poc-app/api/routes"
	"github.com/Laurentus/poc-app/models"
)

func TestCampaigns(t *testing.T) {
	camp := models.Campaign{Name: "Test"}
	data := make(chan models.Campaign)
	data <- camp
	close(data)
	expect_data := []models.Campaign{camp}

	var result []byte
	routes.ListCampaignsResponse(func(values []byte) (int, error) {
		result = values
		return 0, nil
	}, data)
	expected_json, _ := json.Marshal(expect_data)
	if bytes.Compare(result, expected_json) != 0 {
		t.Errorf("Campaing results did not match!")
	}
}

func TestCampaignsWritesEmptyArray(t *testing.T) {
	data := make(chan models.Campaign)
	close(data)

	var result []byte
	routes.ListCampaignsResponse(func(values []byte) (int, error) {
		result = values
		return 0, nil
	}, data)
	expected_json, _ := json.Marshal([]string{})
	if bytes.Compare(result, expected_json) != 0 {
		t.Errorf("Campaign result did not write empty array")
	}
}
