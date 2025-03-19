package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"blog-api/internal/models"
	"blog-api/internal/repositories"
	"blog-api/internal/utils"
	"github.com/go-chi/chi/v5"
)

type PostHandler struct {
	repo *repositories.PostRepository
}

func NewPostHandler(repo *repositories.PostRepository) *PostHandler {
	return &PostHandler{repo: repo}
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.repo.Create(&post); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create post")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, post)
}

func (h *PostHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	post, err := h.repo.GetByID(uint(id))
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, post)
}

func (h *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")
	posts, err := h.repo.GetAll(term)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch posts")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, posts)
}

func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var updatedPost models.Post
	if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.repo.Update(uint(id), &updatedPost); err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, updatedPost)
}

func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
