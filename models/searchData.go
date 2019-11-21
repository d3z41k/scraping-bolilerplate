package models

// SearchData is a struct to search data
type SearchData struct {
	URL     string `json:"url"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Phrase  string `json:"phrase"`
}

// ResultData is a struct to result data
type ResultData struct {
	URL     string
	Status	bool
	Results []map[string]map[string]interface{}
}
