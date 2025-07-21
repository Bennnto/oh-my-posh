package dsc

type Fonts []*Font

type Font struct {
	Name   string `json:"name,omitempty"`
	System bool   `json:"system,omitempty"`
}
