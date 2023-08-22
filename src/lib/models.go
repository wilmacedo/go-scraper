package lib

type ApiNewsSource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ApiNewsArticle struct {
	Source      ApiNewsSource `json:"source"`
	Author      string        `json:"author"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	URL         string        `json:"url"`
	URLToImage  string        `json:"urlToImage"`
	PublishedAt string        `json:"publishedAt"`
	Content     string        `json:"content"`
}

type ApiNewResponse struct {
	Status       string           `json:"ok"`
	TotalResults int              `json:"totalResults"`
	Articles     []ApiNewsArticle `json:"articles"`
}
