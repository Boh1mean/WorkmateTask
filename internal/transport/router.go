package transport

import "net/http"

func NewRouter(h TaskHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /task", h.CreateTask)
	mux.HandleFunc("GET /task/", h.GetTask)
	mux.HandleFunc("DELETE /task/", h.DeleteTask)
	return mux
}
