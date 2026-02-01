package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/hyperx/backend/internal/governance/service"
)

type Handler struct {
	adminService *service.AdminService
}

func NewHandler(adminService *service.AdminService) *Handler {
	return &Handler{adminService: adminService}
}

func (h *Handler) Routes(r chi.Router) {
	r.Route("/admin/actions", func(r chi.Router) {
		r.Post("/suspend", h.SuspendAccount)
		r.Post("/freeze-payout", h.FreezePayout)
		r.Post("/{id}/revoke", h.RevokeAction)
	})
}

type SuspendRequest struct {
	TargetAccountID uuid.UUID `json:"target_account_id"`
	AdminUserID     uuid.UUID `json:"admin_user_id"`
	Reason          string    `json:"reason"`
	DurationSeconds int       `json:"duration_seconds"`
}

func (h *Handler) SuspendAccount(w http.ResponseWriter, r *http.Request) {
	var req SuspendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	duration := time.Duration(req.DurationSeconds) * time.Second
	err := h.adminService.SuspendAccount(r.Context(), req.TargetAccountID, req.AdminUserID, req.Reason, duration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

type FreezeRequest struct {
	TargetAccountID uuid.UUID `json:"target_account_id"`
	AdminUserID     uuid.UUID `json:"admin_user_id"`
	Reason          string    `json:"reason"`
}

func (h *Handler) FreezePayout(w http.ResponseWriter, r *http.Request) {
	var req FreezeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := h.adminService.FreezePayout(r.Context(), req.TargetAccountID, req.AdminUserID, req.Reason)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) RevokeAction(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid action id", http.StatusBadRequest)
		return
	}

	err = h.adminService.RevokeAction(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
