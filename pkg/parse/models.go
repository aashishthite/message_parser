package parse

type Link struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

type ParseResult struct {
	Mentions  []string `json:"mentions,omitempty"`
	Emoticons []string `json:"emoticons,omitempty"`
	Links     []*Link  `json:"links,omitempty"`
}
