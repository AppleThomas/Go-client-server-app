package models

type Album struct {
	ID    string `json:"ObjectID"`
	Name  string `json:"name"`
	Group string `json:"group"`
	Songs string `json:"songs"`
	Year  string `json:"year"`
	Img   string `json:"img"`
}
