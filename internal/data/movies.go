package data

import "time"

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
