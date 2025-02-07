package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lakeside763/synapse-crm-pkg/models"
	"github.com/lakeside763/synapse-crm/api-gateway-service/internals/cores/domains/users/services"
)

type UserHandler struct {
	userService *services.UserServiceClient
}

func NewUserHandler(userService *services.UserServiceClient) *UserHandler {
	return &UserHandler{userService}
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Set content type json
	w.Header().Set("Content-Type", "application/json")

	// Parse the request body
	var input *models.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call the product service
	user, err := u.userService.CreateUser(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}