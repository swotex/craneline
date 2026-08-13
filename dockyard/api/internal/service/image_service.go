package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	db "octopus/internal/db/generated"
)

type ParamInput struct {
	EnvVarName   string
	Type         string
	DefaultValue string
	Required     bool
	Description  string
}

type CreateFullImageInput struct {
	Name        string
	Registry    string
	Description string
	SourceURL   string
	LogoURL     string
	Tag         string
	Digest      string
	Params      []ParamInput
}

type FullImage struct {
	Image    db.Image             `json:"image"`
	Versions []FullImageVersion   `json:"versions"`
}

type FullImageVersion struct {
	db.ImageVersion
	Parameters []db.Parameter `json:"parameters"`
}

func CreateFullImage(ctx context.Context, pool *pgxpool.Pool, in CreateFullImageInput) (db.Image, db.ImageVersion, error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return db.Image{}, db.ImageVersion{}, err
	}
	defer tx.Rollback(ctx) // no-op si commit réussi

	q := db.New(tx)

	image, err := q.CreateImage(ctx, db.CreateImageParams{
		Name:        in.Name,
		Registry:    in.Registry,
		Description: pgtype.Text{String: in.Description, Valid: in.Description != ""},
		SourceUrl:   pgtype.Text{String: in.SourceURL, Valid: in.SourceURL != ""},
		LogoUrl:     pgtype.Text{String: in.LogoURL, Valid: in.LogoURL != ""},
	})
	if err != nil {
		return db.Image{}, db.ImageVersion{}, err
	}

	version, err := q.CreateImageVersion(ctx, db.CreateImageVersionParams{
		ImageID:  image.ID,
		Tag:      in.Tag,
		Digest:   pgtype.Text{String: in.Digest, Valid: in.Digest != ""},
		IsLatest: true,
	})
	if err != nil {
		return db.Image{}, db.ImageVersion{}, err
	}

	for _, p := range in.Params {
		_, err := q.CreateParameter(ctx, db.CreateParameterParams{
			ImageVersionID: version.ID,
			EnvVarName:     p.EnvVarName,
			Type:           p.Type,
			DefaultValue:   pgtype.Text{String: p.DefaultValue, Valid: p.DefaultValue != ""},
			Required:       p.Required,
			Description:    pgtype.Text{String: p.Description, Valid: p.Description != ""},
		})
		if err != nil {
			return db.Image{}, db.ImageVersion{}, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return db.Image{}, db.ImageVersion{}, err
	}

	return image, version, nil
}

func GetFullImage(ctx context.Context, pool *pgxpool.Pool, imageID int64) (FullImage, error) {
	q := db.New(pool)

	image, err := q.GetImageByID(ctx, imageID)
	if err != nil {
		return FullImage{}, err
	}

	versions, err := q.ListVersionsByImageID(ctx, imageID)
	if err != nil {
		return FullImage{}, err
	}

	result := FullImage{Image: image}
	for _, v := range versions {
		params, err := q.ListParametersByVersionID(ctx, v.ID)
		if err != nil {
			return FullImage{}, err
		}
		result.Versions = append(result.Versions, FullImageVersion{
			ImageVersion: v,
			Parameters:   params,
		})
	}

	return result, nil
}