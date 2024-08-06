package handlers

import (
	"encoding/json"
	"net/http"
)

type Domain struct {
	Domain      string `json:"domain"`
	Description string `json:"description"`
}

func TransferHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: simulate http request read for a request
	domains := []Domain{
		{Domain: ".com", Description: "Can transfer"},
		{Domain: ".in", Description: "Can transfer"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(domains)
}
