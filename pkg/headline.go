package pkg

// Headline represents a news headline
type Headline struct {
	URL   *string `json:"url"`
	Title *string `json:"title"`
	Blurb *string `json:"blurb,omitempty"`
}
