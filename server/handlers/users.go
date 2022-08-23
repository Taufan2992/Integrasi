package handlers

import (
	"encoding/json"
	dto "golang/dto/result"
	usersdto "golang/dto/users"
	"golang/models"
	"golang/repositories"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Declare handler struct here ...
type handlerUser struct {
	UserRepository repositories.UserRepository
}

// Declare HandlerUser function here ...
func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

// Declare FindUsers method here ...
func (h *handlerUser) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: users}
	json.NewEncoder(w).Encode(response)
}

// Declare GetUser method here ...
func (h *handlerUser) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: user}
	json.NewEncoder(w).Encode(response)
}

// Declare CreateUser method here ...
func (h *handlerUser) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(usersdto.CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// validation := validator.New()
	// err := validation.Struct(request)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	data, err := h.UserRepository.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

// UPDATE USER
func (h *handlerUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	// user_id, _ := strconv.Atoi(r.FormValue("user_id"))
	request := usersdto.UpdateUserRequest{
		Name:  r.FormValue("name"),
		Email: r.FormValue("email"),
		Image: filename,
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.UserRepository.GetUser(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Image != "" {
		user.Image = request.Image
	}

	data, err := h.UserRepository.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// menampilkan gambar
	user.Image = path_file + user.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

// Declare DeleteUser method here ...
func (h *handlerUser) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

// Declare convertResponse function here ...
func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
