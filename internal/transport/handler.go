package transport

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Boh1mean/workmateTask/internal/service"
)

type TaskHandler interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
	GetTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	usecase service.TaskUsecase
}

func NewHandler(usecase service.TaskUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	task := h.usecase.NewTask()
	response := map[string]string{"id": task.ID}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	task, ok := h.usecase.GetTask(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	resp := map[string]interface{}{
		"id":         task.ID,
		"status":     task.Status,
		"created_at": task.CreatedAt.Format(time.RFC3339),
		"duration":   task.Duration.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	if err := h.usecase.DeleteTask(id); err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	resp := map[string]string{"message": "task deleted"}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
