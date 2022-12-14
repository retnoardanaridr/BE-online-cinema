package routes

import (
	"BE-waysbuck-API/handlers"
	"BE-waysbuck-API/pkg/middleware"
	"BE-waysbuck-API/pkg/mysql"
	"BE-waysbuck-API/repositories"

	"github.com/gorilla/mux"
)

func FilmRoutes(r *mux.Router) {
	filmRepository := repositories.RespositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(filmRepository)

	r.HandleFunc("/films", h.GetFilms).Methods("GET")
	r.HandleFunc("/film/{id}", h.GetFilm).Methods("GET")
	r.HandleFunc("/film", middleware.Auth(middleware.UploadFile(h.CreateFilm))).Methods("POST")
	r.HandleFunc("/film/{id}", middleware.Auth(middleware.UploadFile(h.UpdateFilm))).Methods("PATCH")
	r.HandleFunc("/film/{id}", middleware.Auth(h.DeleteFilm)).Methods("DELETE")

}
