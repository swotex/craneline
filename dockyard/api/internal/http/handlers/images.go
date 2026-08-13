package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/go-chi/chi/v5"

	db "octopus/internal/db/generated"
	"octopus/internal/service"
)

// ImageHandler regroupe les dépendances nécessaires aux routes /images.
type ImageHandler struct {
	Queries *db.Queries
	Pool    *pgxpool.Pool
}

type paramRequest struct {
	EnvVarName   string `json:"env_var_name"`
	Type         string `json:"type"`
	DefaultValue string `json:"default_value"`
	Required     bool   `json:"required"`
	Description  string `json:"description"`
}

type createFullImageRequest struct {
	Name        string         `json:"name"`
	Registry    string         `json:"registry"`
	Description string         `json:"description"`
	SourceURL   string         `json:"source_url"`
	LogoURL     string         `json:"logo_url"`
	Tag         string         `json:"tag"`
	Digest      string         `json:"digest"`
	Params      []paramRequest `json:"params"`
}

func NewImageHandler(q *db.Queries, pool *pgxpool.Pool) *ImageHandler {
	return &ImageHandler{Queries: q, Pool: pool}
}

// GET /api/v1/images?limit=20&offset=0
func (h *ImageHandler) ListImages(w http.ResponseWriter, r *http.Request) {
	limit := int32(20)
	offset := int32(0)

	if v := r.URL.Query().Get("limit"); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil && parsed > 0 {
			limit = int32(parsed)
		}
	}
	if v := r.URL.Query().Get("offset"); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil && parsed >= 0 {
			offset = int32(parsed)
		}
	}

	images, err := h.Queries.ListImages(r.Context(), db.ListImagesParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		respondError(w, http.StatusInternalServerError, "failed to fetch images")
		return
	}

	respondJSON(w, http.StatusOK, images)
}

type createImageRequest struct {
	Name        string `json:"name"`
	Registry    string `json:"registry"`
	Description string `json:"description"`
	SourceURL   string `json:"source_url"`
	LogoURL     string `json:"logo_url"`
}

// POST /api/v1/images
func (h *ImageHandler) CreateImage(w http.ResponseWriter, r *http.Request) {
	var req createImageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	if req.Name == "" {
		respondError(w, http.StatusBadRequest, "name is required")
		return
	}
	if req.Registry == "" {
		req.Registry = "docker.io"
	}

	image, err := h.Queries.CreateImage(r.Context(), db.CreateImageParams{
		Name:        req.Name,
		Registry:    req.Registry,
		Description: pgtype.Text{String: req.Description, Valid: req.Description != ""},
		SourceUrl:   pgtype.Text{String: req.SourceURL, Valid: req.SourceURL != ""},
		LogoUrl:     pgtype.Text{String: req.LogoURL, Valid: req.LogoURL != ""},
	})
	if err != nil {
		// Cas fréquent : violation de la contrainte UNIQUE(registry, name)
		respondError(w, http.StatusConflict, "image already exists or invalid data")
		return
	}

	respondJSON(w, http.StatusCreated, image)
}

// POST /api/v1/images/full
func (h *ImageHandler) CreateFullImage(w http.ResponseWriter, r *http.Request) {
	var req createFullImageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Name == "" || req.Tag == "" {
		respondError(w, http.StatusBadRequest, "name and tag are required")
		return
	}
	if req.Registry == "" {
		req.Registry = "docker.io"
	}

	params := make([]service.ParamInput, 0, len(req.Params))
	for _, p := range req.Params {
		params = append(params, service.ParamInput{
			EnvVarName:   p.EnvVarName,
			Type:         p.Type,
			DefaultValue: p.DefaultValue,
			Required:     p.Required,
			Description:  p.Description,
		})
	}

	image, version, err := service.CreateFullImage(r.Context(), h.Pool, service.CreateFullImageInput{
		Name:        req.Name,
		Registry:    req.Registry,
		Description: req.Description,
		SourceURL:   req.SourceURL,
		LogoURL:     req.LogoURL,
		Tag:         req.Tag,
		Digest:      req.Digest,
		Params:      params,
	})
	if err != nil {
		respondError(w, http.StatusInternalServerError, "failed to create image")
		return
	}

	respondJSON(w, http.StatusCreated, map[string]interface{}{
		"image":   image,
		"version": version,
	})
}

// GET /api/v1/images/{id}
func (h *ImageHandler) GetImage(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	image, err := service.GetFullImage(r.Context(), h.Pool, id)
	if err != nil {
		respondError(w, http.StatusNotFound, "image not found")
		return
	}

	respondJSON(w, http.StatusOK, image)
}
