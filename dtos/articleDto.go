package dtos

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ArticleReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
