package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/hyperx/backend/internal/identity/service"
)

type Handler struct {
	accountService *service.AccountService
	creatorService *service.CreatorService
}

func NewHandler(accountService *service.AccountService, creatorService *service.CreatorService) *Handler {
	return &Handler{
		accountService: accountService,
		creatorService: creatorService,
	}
}

func (h *Handler) Routes(r chi.Router) {
	r.Route("/accounts", func(r chi.Router) {
		r.Post("/", h.CreateAccount)
		r.Get("/{id}", h.GetAccount)
	})

	r.Route("/creators", func(r chi.Router) {
		r.Post("/", h.CreateCreator)
		r.Get("/{id}", h.GetCreator)
	})
}

type CreateAccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	acc, err := h.accountService.CreateAccount(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(acc)
}

func (h *Handler) GetAccount(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid account id", http.StatusBadRequest)
		return
	}

	acc, err := h.accountService.GetAccount(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
}

type CreateCreatorRequest struct {
	AccountID   uuid.UUID `json:"account_id"`
	Slug        string    `json:"slug"`
	DisplayName string    `json:"display_name"`
}

func (h *Handler) CreateCreator(w http.ResponseWriter, r *http.Request) {
	var req CreateCreatorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	cp, err := h.creatorService.CreateProfile(r.Context(), req.AccountID, req.Slug, req.DisplayName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cp)
}

func (h *Handler) GetCreator(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid creator id", http.StatusBadRequest)
		return
	}

	cp, err := h.creatorService.GetProfile(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cp)
}
