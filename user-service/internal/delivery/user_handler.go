package delivery

import (
	"encoding/json"
	"net/http"

	"user-service/internal/entity"
	"user-service/internal/usecase"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)
	h.userUsecase.Register(r.Context(), user)
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req entity.User
	json.NewDecoder(r.Body).Decode(&req)

	token, err := h.userUsecase.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
