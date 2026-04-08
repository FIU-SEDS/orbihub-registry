package store

import (
	"context"

	"github.com/erielC/orbihub-registry/internal/model"
	"github.com/jackc/pgx/v5"
)

type AppsStore struct {
	Conn *pgx.Conn
}

func NewAppsStore(conn *pgx.Conn) AppsStore {
	appStore := AppsStore{Conn: conn}
	return appStore
}

func (as AppsStore) GetApps() ([]model.App, error) {
	var appsList []model.App

	rows, err := as.Conn.Query(context.Background(), "SELECT id, name, description, version, repo, author, image, created_at FROM apps")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var app model.App
		if err := rows.Scan(
			&app.ID,
			&app.Name,
			&app.Description,
			&app.Version,
			&app.Repo,
			&app.Author,
			&app.Image,
			&app.CreatedAt,
		); err != nil {
			return nil, err
		}
		appsList = append(appsList, app)
	}

	return appsList, nil
}

func (as AppsStore) GetAppByID(id string) (model.App, error) {
	var app model.App
	err := as.Conn.QueryRow(context.Background(), "SELECT id, name, description, version, repo, author, image, created_at FROM apps WHERE id = $1", id).Scan(
		&app.ID,
		&app.Name,
		&app.Description,
		&app.Version,
		&app.Repo,
		&app.Author,
		&app.Image,
		&app.CreatedAt,
	)
	if err != nil {
		return app, err
	}
	return app, nil
}

func (as AppsStore) CreateApp(app model.App) error {
	_, err := as.Conn.Exec(context.Background(), "INSERT INTO apps (id, name, description, version, repo, author, image, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", app.ID, app.Name, app.Description, app.Version, app.Repo, app.Author, app.Image, app.CreatedAt)

	return err
}

func (as AppsStore) UpdateApp(id string, app model.App) (model.App, error) {
	var updated model.App
	err := as.Conn.QueryRow(
		context.Background(),
		"UPDATE apps SET name=$1, description=$2, version=$3, repo=$4, author=$5, image=$6 WHERE id=$7 RETURNING id, name, description, version, repo, author, image, created_at",
		app.Name, app.Description, app.Version, app.Repo, app.Author, app.Image, id,
	).Scan(
		&updated.ID,
		&updated.Name,
		&updated.Description,
		&updated.Version,
		&updated.Repo,
		&updated.Author,
		&updated.Image,
		&updated.CreatedAt,
	)
	return updated, err
}

func (as AppsStore) DeleteApp(id string) error {
	tag, err := as.Conn.Exec(context.Background(), "DELETE FROM apps WHERE id=$1", id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
