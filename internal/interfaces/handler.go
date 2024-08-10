package interfaces

import (
	"encoding/json"
	"github.com/DemchukM/api-gateway/internal/domain"
	"github.com/DemchukM/api-gateway/internal/usecase"
	"net/http"
	"strconv"
)

type UserHandler struct {
	useCase *usecase.UserUseCase
}

func NewUserHandler(useCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	ID, _ := strconv.Atoi(idStr)

	user, err := h.useCase.GetUserById(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	json.NewDecoder(r.Body).Decode(&user)

	id, err := h.useCase.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(id)))
}
