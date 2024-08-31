package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/guluzadehh/kode_test/utils"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	sr *AuthService
}

func NewHandler(s *AuthService) *Handler {
	return &Handler{sr: s}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/users", h.handleUsers)
}

func (h *Handler) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.listUsers(w, r)
	case "POST":
		h.registerUser(w, r)
	}
}

func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	payload := &CreateUserRequest{}
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, fmt.Errorf("wrong json format"))
		return
	}

	if _, exists := h.sr.Store.GetByUsername(payload.Username); exists {
		utils.ErrorJSON(w, http.StatusBadRequest, fmt.Errorf("username %s is already taken", payload.Username))
		return
	}

	if payload.Password != payload.ConfirmPassword {
		utils.ErrorJSON(w, http.StatusBadRequest, fmt.Errorf("passwords do not match"))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Password encrypt error: %s\n", err)
		utils.ErrorJSON(w, http.StatusInternalServerError, fmt.Errorf("error happened"))
		return
	}

	user, err := h.sr.Store.Create(payload.Username, string(hashedPassword))
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, fmt.Errorf("error while creating the user"))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, user)
}

func (h *Handler) listUsers(w http.ResponseWriter, r *http.Request) {
	users := h.sr.Store.GetAll()
	utils.WriteJSON(w, http.StatusOK, users)
}
