package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/amulya.bl/demoCRUDApp/entities"
	"github.com/amulya.bl/demoCRUDApp/request"
	"github.com/amulya.bl/demoCRUDApp/response"
	"github.com/amulya.bl/demoCRUDApp/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}
func entityToResponse(userEntity entities.User) (responseUser response.User) {
	responseUser = response.User{
		ID:    userEntity.ID,
		Name:  userEntity.Name,
		Email: userEntity.Email,
	}
	return responseUser
}
func errorHandling(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	switch {
	case err.Error() == "Name should be at least 3 characters":
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	case err.Error() == "Invalid user, Enter valid userID":
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	case err.Error() == "record not found":
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
	default:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}

}
func (uc *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var requestUser request.CreateUser
	json.NewDecoder(r.Body).Decode(&requestUser)

	err := uc.Service.CreateUser(&requestUser)

	if err != nil {
		errorHandling(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User Created Sucessfully"})
	return

}

func (uc *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {

	urlInput := mux.Vars(r)
	var idStr = urlInput["id"]
	userId, _ := strconv.Atoi(idStr)

	userEntity, err := uc.Service.GetUser(userId)
	if err != nil {
		errorHandling(w, err)
		return
	}
	userResponse := entityToResponse(userEntity)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResponse)

}

func (uc *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	urlInput := mux.Vars(r)
	idStr := urlInput["id"]
	userId, err := strconv.Atoi(idStr)

	var updateUser request.UpdateUser
	err = json.NewDecoder(r.Body).Decode(&updateUser)

	err = uc.Service.UpdateUser(userId, updateUser)
	if err != nil {
		errorHandling(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User Updated Successfully"})
	fmt.Fprintf(w, "User updated successfully")
	return

}

func (uc *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	urlInput := mux.Vars(r)
	idStr := urlInput["id"]
	userId, err := strconv.Atoi(idStr)

	err = uc.Service.DeleteUser(userId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}
	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User Deleted Successfully"})
	fmt.Fprintf(w, "User deleted successfully")

}

func (uc *UserHandler) GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	users, err := uc.Service.GetAllUsers()
	if err != nil {
		http.Error(w, "Could not fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
