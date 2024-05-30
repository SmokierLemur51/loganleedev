package data

import (
	"database/sql"
	// "strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

type World struct {
	ID         int          `db:"id"`
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  time.Time    `db:"updated_at"`
	DeletedAt  sql.NullTime `db:"deleted_at"`
	Name       string       `db:"name"`
	Categories []LocationCategory
	Locations  []Location
}

func PopulateWorlds(db *sqlx.DB) {
	db.MustExec("insert into worlds (name) values (?)", "Greek Mythology")
	db.MustExec("insert into worlds (name) values (?)", "First World Together")
	db.MustExec("insert into worlds (name) values (?)", "Hogsmead")
}

func LoadWorldByName(db *sqlx.DB, name string) (World, error) {
	var world World
	err := db.Get(&world, "select * from worlds where name = ?", name)
	if err != nil {
		return World{}, err
	}
	err = db.Select(&world.Categories, "select * from location_categories")
	if err != nil {
		return World{}, err
	}
	err = db.Select(&world.Locations, "select * from locations where world_id = ?", world.ID)
	if err != nil {
		return World{}, err
	}
	return world, nil
}
