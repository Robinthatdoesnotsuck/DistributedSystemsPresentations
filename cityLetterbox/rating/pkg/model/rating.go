package model

type RecordID string

type RecordType string

const (
	RecordTypeMovie = RecordType("movie")
)

type UserID string
type RatingValue int

type Rating struct {
	RecordID   string      `json:"recordId"`
	RecordType string      `json:"recordType"`
	UserID     UserID      `json:"userId"`
	Value      RatingValue `json:"value"`
}
