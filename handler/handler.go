package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vieiravitor/rest-server/domain"
)

type Handler struct{}

func (h *Handler) SaveTransaction(w http.ResponseWriter, r *http.Request) {
	var paymentRequest domain.PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&paymentRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(paymentRequest)
	var paymentResponse domain.PaymentResponse
	paymentResponse.ID = paymentRequest.ID
	paymentResponse.Response = "Payment completed"
	respondJson(w, http.StatusOK, paymentResponse)
}

func respondJson(w http.ResponseWriter, statusCode int, paymentResponse domain.PaymentResponse) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paymentResponse)
}
