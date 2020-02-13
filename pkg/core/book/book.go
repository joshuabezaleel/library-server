package book

import (
	"time"

	"github.com/lib/pq"
)

//
// const (
// 	A = "GENERAL WORKS"
// 	B = "PHILOSOPHY, PSYCHOLOGY, RELIGION"
// 	C = "AUXILIARY SCIENCES OF HISTORY"
// 	D = "GENERAL AND OLD WORLD HISTORY"
// 	E = "HISTORY OF THE AMERICAS"
// 	F = "HISTORY OF THE AMERICAS"
// 	G = "GEOGRAPHY, ANTHROPOLOGY, AND RECREATION"
// 	H = "SOCIAL SCIENCES"
// 	J = "POLITICAL SCIENCES"
// 	K = "LAW"
// 	L = "EDUCATION"
// 	M = "MUSIC"
// 	N = "FINE ARTS"
// 	P = "LANGUAGE AND LITERATURE"
// 	Q = "SCIENCE"
// 	R = "MEDICINE"
// 	S = "AGRICULTURE"
// 	T = "TECHNOLOGY"
// 	U = "MILITARY SCIENCE"
// 	V = "NAVAL SCIENCE"
// 	Z = "BIBLIOGRAPHY, LIBRARY SCIENCE, AND GENERAL INFORMATION RESOURCES"
// )

// Book domain model.
type Book struct {
	ID                string         `json:"id" db:"id"`
	Title             string         `json:"title" db:"title"`
	Publisher         string         `json:"publisher" db:"publisher"`
	YearPublished     int            `json:"yearPublished" db:"year_published"`
	CallNumber        string         `json:"callNumber" db:"call_number"`
	CoverPicture      string         `json:"coverPicture" db:"cover_picture"`
	ISBN              string         `json:"isbn" db:"isbn"`
	Collation         string         `json:"collation" db:"book_collation"`
	Edition           int            `json:"edition" db:"edition"`
	Description       string         `json:"description" db:"description"`
	LOCClassification string         `json:"locClassification" db:"loc_classification"`
	Subject           pq.StringArray `json:"subject" db:"subject"`
	Author            pq.StringArray `json:"author" db:"author"`
	Quantity          int            `json:"quantity" db:"quantity"`
	AddedAt           time.Time      `json:"addedAt" db:"added_at"`
}

// NewBook creates a new instance of Book domain model.
func NewBook(id string, title string, publisher string, yearPublished int, callNumber string, coverPicture string, isbn string, collation string, edition int, description string, locClassification string, author []string, quantity int, addedAt time.Time) *Book {
	return &Book{
		ID:                id,
		Title:             title,
		Publisher:         publisher,
		YearPublished:     yearPublished,
		CallNumber:        callNumber,
		CoverPicture:      coverPicture,
		ISBN:              isbn,
		Collation:         collation,
		Edition:           edition,
		Description:       description,
		LOCClassification: locClassification,
		Author:            author,
		Quantity:          quantity,
		AddedAt:           addedAt,
	}
}
