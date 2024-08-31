package note

import (
	"fmt"
	"log"
	"net/http"

	"github.com/guluzadehh/kode_test/apps/auth"
	"github.com/guluzadehh/kode_test/speller"
	"github.com/guluzadehh/kode_test/utils"
)

type Handler struct {
	sr   *NoteService
	auth *auth.Handler
}

func NewHandler(s *NoteService, auth *auth.Handler) *Handler {
	return &Handler{sr: s, auth: auth}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/notes", h.auth.BasicAuthMiddleware(h.handleNotes))
}

func (h *Handler) handleNotes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.listNotes(w, r)
	case "POST":
		h.createNote(w, r)
	}
}

func (h *Handler) listNotes(w http.ResponseWriter, r *http.Request) {
	user := auth.GetUser(r)
	notes := h.sr.Store.GetAllOfUser(user.Id)
	utils.WriteJSON(w, http.StatusOK, notes)
}

func (h *Handler) createNote(w http.ResponseWriter, r *http.Request) {
	payload := &CreateNoteRequest{}
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, fmt.Errorf("wrong json format"))
		return
	}

	if len(payload.Text) == 0 {
		utils.ErrorJSON(w, http.StatusBadRequest, fmt.Errorf("text is empty"))
		return
	}

	if spellRes, err := speller.CheckText(payload.Text); err != nil {
		log.Println("Speller error: ", err)
		utils.ErrorJSON(w, http.StatusInternalServerError, fmt.Errorf("error while validating the text"))
		return
	} else if len(spellRes) > 0 {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"error":    "text contains mistakes",
			"mistakes": spellRes,
		})
		return
	}

	user := auth.GetUser(r)
	note, err := h.sr.Store.Create(payload.Text, user.Id)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, fmt.Errorf("error while creating the note"))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, note)
}
