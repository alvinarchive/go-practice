package data

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
	"movie.alvintanoto.id/internal/validator"
)

// change key in json object to json snakecase
// to control the visibility of data:
// use - to hide
// use omitempty to hide zero, null, or empty string value
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime"` // use runtime type instead of int32
	Genres    []string  `json:"genres"`
	Version   int32     `json:"version"` // incremented every time movie info got updated
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year > 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must at least contains 1 genres")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contains more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contains duplicate values")
}

// Define a model struct type
type MovieModel struct {
	DB *sql.DB
}

// Add a placeholder method for inserting a new record
func (m MovieModel) Insert(movie *Movie) error {
	query := `
	INSERT INTO movies (title, year, runtime, genres) 
	VALUES ($1, %2, %3, %4)
	RETURNING id, created_at, version
	`

	args := []interface{}{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres)}
	return m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

func (m MovieModel) Get(id int64) (*Movie, error) {
	return nil, nil
}

func (m MovieModel) Update(movie *Movie) error {
	return nil
}

func (m MovieModel) Delete(id int64) error {
	return nil
}
