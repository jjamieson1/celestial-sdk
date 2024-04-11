package models

type SearchResult struct {
	Name         string  `json:"name"`
	Manufacturer string  `json:"manufacturer,omitempty"`
	Id           string  `json:"id"`
	IsDisplayed  bool    `json:"isDisplayed"`
	Image        string  `json:"image"`
	Score        float64 `json:"score"`
}

type CMSSearchResult struct {
	Id        string  `json:"id"`
	Title     string  `json:"title"`
	Body      string  `json:"body"`
	ShortText string  `json:"shortText"`
	Score     float64 `json:"score"`
}
