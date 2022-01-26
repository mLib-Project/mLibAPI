package types

type Book struct {
	Family      string `form:"family" json:"family"`
	Category    string `form:"category" json:"category"`
	Subcategory string `form:"subcategory" json:"subcategory"`
	Name        string `form:"name" json:"name"`
	Author      string `form:"author" json:"author"`
	Source      string `form:"source" json:"source"`
	ID          string `form:"ID" json:"ID"`
}
