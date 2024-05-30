package data

import (
	"github.com/jmoiron/sqlx"
)

func CreateSchema(db *sqlx.DB) {
	schema := `
	
		CREATE TABLE IF NOT EXISTS worlds (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			deleted_at DATETIME,
			name VARCHAR(120)
		);

		CREATE TABLE IF NOT EXISTS location_categories (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			deleted_at DATETIME,
			category VARCHAR(120) UNIQUE,
			description VARCHAR(120)
		);

		CREATE TABLE IF NOT EXISTS locations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			deleted_at DATETIME,
			world_id INTEGER NOT NULL,
			category_id INTEGER NOT NULL,
			name VARCHAR(120) NOT NULL,
			x VARCHAR(50) NOT NULL,
			y VARCHAR(50) NOT NULL,
			z VARCHAR(50) NOT NULL,
			info VARCHAR(500),
			FOREIGN KEY (world_id) REFERENCES worlds(id),
			FOREIGN KEY (category_id) REFERENCES location_categories(id)
		);
	`

	db.MustExec(schema)
}
