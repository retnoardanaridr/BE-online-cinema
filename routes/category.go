package routes

import (
	"BE-waysbuck-API/handlers"
	"BE-waysbuck-API/pkg/mysql"
	"BE-waysbuck-API/repositories"

	"github.com/gorilla/mux"
)

func CategoryRoutes(r *mux.Router) {
	categoryRepository := repositories.RepositoryCategory(mysql.DB)
	h := handlers.HandlerCategory(categoryRepository)

	r.HandleFunc("/categories", h.GetCategories).Methods("GET")
	r.HandleFunc("/category/{id}", h.GetCategory).Methods("GET")
	r.HandleFunc("/category", h.CreateCategory).Methods("POST")
	r.HandleFunc("/category/{id}", h.UpdateCategory).Methods("PATCH")
	r.HandleFunc("/category/{id}", h.DeleteCategory).Methods("DELETE")
}
