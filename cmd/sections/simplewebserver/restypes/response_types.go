package restypes

import "time"

type Attributes struct {
	PublishOn time.Time `json::"publishOn"`
	Title     string    `json:"title"`
}

type DataItem struct {
	Attributes Attributes `json:"attributes"`
}

type ArticleResponse struct {
	Data []DataItem `json:"data"`
}
