package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"strings"
	"kasir-api/models"
	"kasir-api/services"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

// multiple item apa aja, quantity nya
func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req models.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	transaction, err := h.service.Checkout(req.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

func (h *TransactionHandler) HandleReportHariIni(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.ReportHariIni(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) ReportHariIni(w http.ResponseWriter, r *http.Request) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	end := start.AddDate(0, 0, 1) // end exclusive

	report, err := h.service.GetReport(start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

func (h *TransactionHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.ReportRange(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) ReportRange(w http.ResponseWriter, r *http.Request) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	q := r.URL.Query()

	startStr := strings.TrimSpace(q.Get("start_date"))
	endStr := strings.TrimSpace(q.Get("end_date"))

	// kalau tidak ada query -> default hari ini
	if startStr == "" && endStr == "" {
		now := time.Now().In(loc)
		start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
		end := start.AddDate(0, 0, 1)

		report, err := h.service.GetReport(start, end)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(report)
		return
	}

	// validasi query
	if startStr == "" || endStr == "" {
		http.Error(w, "start_date dan end_date harus diisi", http.StatusBadRequest)
		return
	}

	startDate, err := time.ParseInLocation("2006-01-02", startStr, loc)
	if err != nil {
		http.Error(w, "format start_date harus YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	endDate, err := time.ParseInLocation("2006-01-02", endStr, loc)
	if err != nil {
		http.Error(w, "format end_date harus YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	if endDate.Before(startDate) {
		http.Error(w, "end_date tidak boleh lebih kecil dari start_date", http.StatusBadRequest)
		return
	}

	start := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, loc)
	endExclusive := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, loc).AddDate(0, 0, 1)

	report, err := h.service.GetReport(start, endExclusive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
