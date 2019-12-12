package book

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Book domain model.
type Book struct {
	ID                string      `json:"id" db:"id"`
	Title             string      `json:"title" db:"title"`
	Publisher         string      `json:"publisher" db:"publisher"`
	YearPublished     int         `json:"yearPublished" db:"year_published"`
	CallNumber        string      `json:"callNumber" db:"call_number"`
	CoverPicture      string      `json:"coverPicture" db:"cover_picture"`
	ISBN              string      `json:"isbn" db:"isbn"`
	Collation         string      `json:"collation" db:"book_collation"`
	Edition           int         `json:"edition" db:"edition"`
	Description       string      `json:"description" db:"description"`
	LOCClassification string      `json:"locClassification" db:"loc_classification"`
	Subject           StringSlice `json:"subject" db:"subject"`
	Author            StringSlice `json:"author" db:"author"`
	Quantity          int         `json:"quantity" db:"quantity"`
	AddedAt           time.Time   `json:"addedAt" db:"added_at"`
}

// NewBook creates a new instance of Book domain model.
func NewBook(id string, title string, publisher string, yearPublished int, callNumber string, coverPicture string, isbn string, collation string, edition int, description string, locClassification string, subject []string, author []string, quantity int, addedAt time.Time) *Book {
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
		Subject:           subject,
		Author:            author,
		Quantity:          quantity,
		AddedAt:           addedAt,
	}
}

// StringSlice is used as a slice of string to convert the particular type in every struct
// of domain model to string for querying compatibility.
// Inspiration: https://www.codementor.io/tamizhvendan/leveraging-interfaces-in-golang-part-2-9skptyl78
type StringSlice []string

// Value is used to convert the slice of string to joined string with "," delimiter.
func (stringSlice StringSlice) Value() (driver.Value, error) {
	var quotedStrings []string
	for _, str := range stringSlice {
		quotedStrings = append(quotedStrings, strconv.Quote(str))
	}
	value := fmt.Sprintf("{ %s }", strings.Join(quotedStrings, ","))
	return value, nil
}
