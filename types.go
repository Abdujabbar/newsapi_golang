package newsapi_golang

//APIResponseArticles struct
type APIResponseArticles struct {
	Code         string    `json:"code"`
	Message      string    `json:"message"`
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

//APIResponseSources struct
type APIResponseSources struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Status  string   `json:"status"`
	Sources []Source `json:"sources"`
}

//Source ds
type Source struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}

//Article ds
type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
}
