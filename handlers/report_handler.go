package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"kasir-api/models"
	"kasir-api/repositories"
)

type ReportHandler struct {
	transactionRepo repositories.TransactionRepositoryInterface
}

func NewReportHandler(transactionRepo repositories.TransactionRepositoryInterface) *ReportHandler {
	return &ReportHandler{transactionRepo: transactionRepo}
}

func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")

	var startDate, endDate time.Time
	var err error

	// Default to today if dates are not provided
	if startDateStr == "" || endDateStr == "" {
		now := time.Now()
		startDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		endDate = startDate
	} else {
		// Parse provided dates
		layout := "2006-01-02"
		startDate, err = time.Parse(layout, startDateStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrInternal, "Invalid start_date format. Use YYYY-MM-DD"))
			return
		}

		endDate, err = time.Parse(layout, endDateStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrInternal, "Invalid end_date format. Use YYYY-MM-DD"))
			return
		}
	}

	reportData, err := h.transactionRepo.GetReport(startDate, endDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Error(http.StatusInternalServerError, models.ErrInternal, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(models.Success(http.StatusOK, "Report retrieved successfully", reportData))
}
