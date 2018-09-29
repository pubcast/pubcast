package webfinger

// Link to another ActivityStreams object
type Link struct {
	Rel  string `json:"rel"`
	Type string `json:"type"`
	HREF string `json:"href"`
}

// Actor includes links to an ActivityStreams Actor object
type Actor struct {
	Subject string `json:"subject"`
	Links   []Link `json:"links"`
}
