package api

import "time"

type NewAuthor struct {
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

type Author struct {
	ID int64 `json:"id"`
	// Nested struct
	NewAuthor
}

type UpdateAuthor struct {
	Version int64 `json:"version"`
	// Nested struct
	Author
}

type NewBook struct {
	Title            string    `json:"title"`
	Slugs            []string  `json:"slugs"`
	AuthorId         int64     `json:"authorId"`
	DateOfPublishing time.Time `json:"dateOfPublishing"`
}

type Book struct {
	ID int64 `json:"id"`
	// nested struct
	NewBook
	// nested struct
	Author Author
}

type UpdateBook struct {
	Version int64 `json:"version"`
	// nested struct
	NewBook
}
