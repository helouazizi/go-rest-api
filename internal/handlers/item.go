package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go-rest-api/internal/services"
	"go-rest-api/pkg/logger"
)

type ItemHandler struct {
	Service *services.ItemService
}

func NewItemHandler(itemService *services.ItemService) *ItemHandler {
	return &ItemHandler{Service: itemService}
}

func (h *ItemHandler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	items := h.Service.GetAllItems()
	json.NewEncoder(w).Encode(items)
}

func (h *ItemHandler) GetItemById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/items/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.LogWithDetails(err)
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	item, found := h.Service.GetItemById(id)
	if !found {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.Name == "" {
		logger.LogWithDetails(err)
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	item := h.Service.CraeteItem(input.Name)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/items/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.LogWithDetails(err)
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	if !h.Service.DeleteItem(id) {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
