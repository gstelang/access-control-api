package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	client "github.com/openfga/go-sdk/client"
)

func DomainAuthorization(w http.ResponseWriter, r *http.Request) {
	domainName := r.PathValue("domainName")

	action := r.URL.Query().Get("action")
	if action == "" {
		http.Error(w, "Missing action parameter", http.StatusBadRequest)
		return
	}

	user := r.URL.Query().Get("user")
	if action == "" {
		http.Error(w, "Missing user parameter", http.StatusBadRequest)
		return
	}

	fgaClient, ok := r.Context().Value("fgaClient").(*client.OpenFgaClient)
	if !ok {
		http.Error(w, "FGA client not found", http.StatusInternalServerError)
		return
	}

	allowed, err := performAuthCheck(r.Context(), fgaClient, user, action, "domain:"+domainName)
	if err != nil {
		http.Error(w, "Error checking authorization", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"allowed": allowed})
}

func performAuthCheck(ctx context.Context, fgaClient *client.OpenFgaClient, user, relation, object string) (bool, error) {

	body := client.ClientCheckRequest{
		User:     "user:" + user,
		Relation: relation,
		Object:   object,
	}

	response, err := fgaClient.Check(ctx).Body(body).Execute()
	if err != nil {
		fmt.Print(err)
		return false, fmt.Errorf("error checking authorization: %v", err)
	}

	return response.GetAllowed(), nil
}
