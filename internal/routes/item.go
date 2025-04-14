package routes

import (
	"myapi/internal/handlers"

	"github.com/gorilla/mux"
)

func ItemRoutes(r *mux.Router) {
	r.HandleFunc("/api/itens", handlers.ListItems).Methods("GET")
	r.HandleFunc("/api/itens/{id}", handlers.GetItem).Methods("GET")
	r.HandleFunc("/api/itens/codigo/{codigo}", handlers.GetItemByCode).Methods("GET")
	r.HandleFunc("/api/itens", handlers.CreateItem).Methods("POST")
	r.HandleFunc("/api/itens", handlers.UpdateItem).Methods("PUT")
	r.HandleFunc("/api/itens/{id}", handlers.DeleteItem).Methods("DELETE")
}
