package model

type Book struct {
	ID          int     `json:"id"`
	Title       string  `form:"title" binding:"required" json:"title"`
	Code        string  `form:"code" binding:"required" json:"code"`
	Author      string  `form:"author" binding:"required" json:"author"`
	PublishYear *int    `form:"publishYear" json:"publish-year"`
	Country     *string `form:"country,default=Indonesia" json:"country"`
}

type Index struct {
	LastId int
}
