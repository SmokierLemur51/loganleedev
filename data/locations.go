package data

import (
	"database/sql"
	// "strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

type LocationCategory struct {
	ID          int            `db:"id"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
	DeletedAt   sql.NullTime   `db:"deleted_at"`
	Category    string         `db:"category"`
	Description sql.NullString `db:"description"`
}

type Location struct {
	ID         int            `db:"id"`
	CreatedAt  time.Time      `db:"created_at"`
	UpdatedAt  time.Time      `db:"updated_at"`
	DeletedAt  sql.NullTime   `db:"deleted_at"`
	WorldID    int            `db:"world_id"`
	CategoryID int            `db:"category_id"`
	Name       string         `db:"name"`
	X          string         `db:"x"`
	Y          string         `db:"y"`
	Z          string         `db:"z"`
	Info       sql.NullString `db:"info"`
}

func PopulateLocationCategories(db *sqlx.DB) {
	db.MustExec("insert into location_categories (category, description) values (?,?)",
		"Generated Villages", "Minecraft generated villages")
	db.MustExec("insert into location_categories (category, description) values (?,?)",
		"Biomes", "Generated or designed biomes.")
	db.MustExec("insert into location_categories (category, description) values (?,?)",
		"Ravines", "Naturally generated ravines")
	db.MustExec("insert into location_categories (category, description) values (?,?)",
		"Dungeons", "")
	db.MustExec("insert into location_categories (category, description) values (?,?)",
		"Pillager Outposts", "")
	db.MustExec("insert into location_categories (category, description) values (?,?)",
		"Mob Spawners", "Spawners, you should specify the mob that spawns from it.")
	db.MustExec("insert into location_categories (category, description) values (?,?)",
		"Abandoned Mineshafts", "")
	db.MustExec("insert into location_categories (category, description) values (?,?)",
		"Secret Stashes", "Secret stashes people shouldn't know about.")
	db.MustExec("insert into location_categories (category, description) values (?,?)",
		"Storage Rooms", "Shared storage rooms with other players.")
	db.MustExec("insert into location_categories (category, description) values (?,?)",
		"Player Bases", "Other players bases for intelligence gathering.")
}
