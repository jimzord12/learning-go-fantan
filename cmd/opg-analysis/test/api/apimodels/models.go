package apimodels

import "time"

// Here we are trying to create structs that reflect the shape of the Incoming
// JSON data
type Attributes struct {
	PublishOn time.Time `json:"publishOn"`
	Title     string    `json:"title"`
}

type SeekingAlplaNews struct {
	Attributes Attributes `json:"attributes"`
}

type SeekingAlplaNewsResponse struct {
	Articles []SeekingAlplaNews `json:"data"`
}
