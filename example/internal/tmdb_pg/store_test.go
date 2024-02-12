package tmdb_pg

import (
	"testing"

	models "github.com/aakash-rajur/example/internal/tmdb_pg/models"
	store "github.com/aakash-rajur/example/internal/tmdb_pg/store"
	"github.com/jmoiron/sqlx"
	"gotest.tools/assert"

	// "models"
	// "store"

	_ "github.com/lib/pq"
)

func TestMoviesFind(t *testing.T) {
	t.Run("without alt path", func(t *testing.T) {
		engine := "postgres"
		connectionUrl := "postgres://app:app@localhost:54320/app?sslmode=disable"

		db, err := sqlx.Open(engine, connectionUrl)
		if err != nil {
			t.Errorf("unable to connect to database: %v", err)
		}

		var id int32 = 24
		m := models.Movie{
			Id: &id,
		}

		// Count
		var countResult int64 = 1
		countMovies, err := store.Count[*models.Movie](db, &m)
		if err != nil {
			t.Errorf("unable : %v", err)
		}
		assert.Equal(t, countResult, *countMovies)

		// Find by PK
		killBillVol1, err := store.FindByPk[*models.Movie](db, &m)
		if err != nil {
			t.Errorf("unable : %v", err)
		}
		assert.Equal(t, "Kill Bill: Vol. 1", *killBillVol1.Title)
		assert.Equal(t, id, *killBillVol1.Id)

		// Find One
		var title string = "Kill Bill: Vol. 1"
		m2 := models.Movie{
			Title: &title,
		}
		killBillVol1, err = store.FindOne[*models.Movie](db, &m2)
		if err != nil {
			t.Errorf("unable : %v", err)
		}
		assert.Equal(t, "Kill Bill: Vol. 1", *killBillVol1.Title)
		assert.Equal(t, id, *killBillVol1.Id)

		// Find First
		killBillVol1, err = store.FindFirst[*models.Movie](db, &m2)
		if err != nil {
			t.Errorf("unable : %v", err)
		}
		assert.Equal(t, "Kill Bill: Vol. 1", *killBillVol1.Title)
		assert.Equal(t, id, *killBillVol1.Id)

		// Find Many
		var lang string = "zh"
		m3 := models.Movie{
			OriginalLanguage: &lang,
		}
		manyMovies, err := store.FindMany[*models.Movie](db, &m3)
		if err != nil {
			t.Errorf("unable : %v", err)
		}
		assert.Equal(t, 27, len(manyMovies))
		assert.Equal(t, "zh", *manyMovies[0].OriginalLanguage)

		println("store find done")

	})
}
