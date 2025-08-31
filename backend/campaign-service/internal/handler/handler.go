package handler

import (
	"fmt"
	"net/http"
)

func GetCampaigns(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List of campaigns")
}
