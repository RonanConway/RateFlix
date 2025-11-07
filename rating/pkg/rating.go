package model

// defines a record Id. Together with RecordType identifies unique records across all types
type RecordID string

type RecordType string

// existing record types
const (
	RecordTypeMovie = RecordType("movie")
)

// UserID defines a user id
type UserID string

// RatingValue defines a value of a rating record
type RatingValue int

// Rating defines an individual rating created by a user for some record.
type Rating struct {
	RecordID   string      `json:"recordId"`
	RecordType string      `json:"recordType"`
	UserID     UserID      `json:"userId"`
	Value      RatingValue `json:"value"`
}
