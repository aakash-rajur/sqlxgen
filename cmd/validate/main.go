package main

import (
	"fmt"

	"github.com/aakash-rajur/sqlxgen/gen/tmdb_pg/models"
	"github.com/aakash-rajur/sqlxgen/gen/tmdb_pg/store"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", "postgres://app:app@localhost:54320/app?sslmode=disable")

	if err != nil {
		utils.ExitWithError(err)

		return
	}

	tx, err := db.Beginx()

	if err != nil {
		utils.ExitWithError(err)

		return
	}

	defer func(tx *sqlx.Tx) {
		err := tx.Rollback()

		if err != nil {
			utils.ExitWithError(err)

			return
		}
	}(tx)

	actor := models.Actor{
		//Id:   utils.PointerTo(int32(50000)),
		Name: utils.PointerTo("Aakash Rajur"),
	}

	err = store.Insert(tx, &actor)

	if err != nil {
		utils.ExitWithError(err)

		return
	}

	fmt.Println(actor)
}
