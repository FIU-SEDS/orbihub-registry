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
