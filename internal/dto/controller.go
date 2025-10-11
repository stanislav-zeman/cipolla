package dto

type Controller struct {
	Name    string
	Methods []string

	Framework Framework `json:"-"`
	Import    Import    `json:"-"`
	Logger    Logger    `json:"-"`
}
