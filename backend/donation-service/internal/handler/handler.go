package handler

import (
	"fmt"
	"net/http"
)

func GetDonations(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List of donations")
}
