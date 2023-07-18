package models

type InlineResponse201 struct {
	Change int    `json:"change"`
	Thanks string `json:"thanks,omitempty"`
}
