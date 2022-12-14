package handlers

import (
	filmdto "BE-waysbuck-API/dto/film"
	dto "BE-waysbuck-API/dto/result"
	"BE-waysbuck-API/models"
	"BE-waysbuck-API/repositories"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerFilm struct {
	FilmRepository repositories.FilmRepository
}

func HandlerFilm(FilmRepository repositories.FilmRepository) *handlerFilm {
	return &handlerFilm{FilmRepository}
}

func (h *handlerFilm) GetFilms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	films, err := h.FilmRepository.GetFilms()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: films}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) GetFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	films, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: films}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) CreateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//get token user
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]

	if userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "not an admin"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//img
	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	price, _ := strconv.Atoi(r.FormValue("price"))
	category_id, _ := strconv.Atoi(r.FormValue("category_id"))
	request := filmdto.FilmRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		FilmUrl:     r.FormValue("filmurl"),
		CategoryID:  category_id,
		Price:       price,
		Thumbnail:   filename,
	}

	validation := validator.New()
	error := validation.Struct(request)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: error.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	//get data category by id
	category, _ := h.FilmRepository.GetCategoriesbyId(category_id)

	film := models.Film{
		Title:       request.Title,
		Category:    category,
		Price:       request.Price,
		Filmurl:     request.FilmUrl,
		Description: request.Description,
		Thumbnail:   os.Getenv("PATH_FILE") + request.Thumbnail,
	}

	film, err := h.FilmRepository.CreateFilm(film)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: film}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) UpdateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]

	if userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "not an admin"}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	//sesuai request
	price, _ := strconv.Atoi(r.FormValue("price"))
	request := filmdto.FilmRequest{
		Title:     r.FormValue("title"),
		Price:     price,
		Thumbnail: filename,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	film, err := h.FilmRepository.GetFilm(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	//form inputan masih kurang
	if request.Title != "" {
		film.Title = request.Title
	}

	if request.Price != 0 {
		film.Price = request.Price
	}

	if request.Thumbnail != "empty" {
		film.Thumbnail = request.Thumbnail
	}

	data, err := h.FilmRepository.UpdateFilm(film)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// FilmResponse := filmdto.FilmResponse{
	// 	ID:    data.ID,
	// 	Title: data.Title,
	// 	Price: data.Price,
	// 	// CategoryID:  data.CategoryID,
	// 	FilmUrl:     data.Filmurl,
	// 	Description: data.Description,
	// 	Thumbnail:   os.Getenv("PATH_FILE") + data.Thumbnail,
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]

	if userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "not an admin"}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.FilmRepository.DeleteFilm(film)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)

}

// func filmResponse(u models.Film) filmdto.FilmResponse {
// 	return filmdto.FilmResponse{
// 		ID:        u.ID,
// 		Title:     u.Title,
// 		Price:     u.Price,
// 		Thumbnail: u.Thumbnail,
// 	}

// }
