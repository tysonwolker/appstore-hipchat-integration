package models

type Review struct {
	Author struct {
		URI struct {
			Label string `json:"label"`
		} `json:"uri"`
		Name struct {
			Label string `json:"label"`
		} `json:"name"`
		Label string `json:"label"`
	} `json:"author"`
	Version struct {
		Label string `json:"label"`
	} `json:"im:version"`
	Rating struct {
		Label string `json:"label"`
	} `json:"im:rating"`
	ID struct {
		Label string `json:"label"`
	} `json:"id"`
	Title struct {
		Label string `json:"label"`
	} `json:"title"`
	Content struct {
		Label string `json:"label"`
		Attributes struct {
			Type string `json:"type"`
		} `json:"attributes"`
	} `json:"content"`
	Link struct {
		Attributes struct {
			Rel string `json:"rel"`
			Href string `json:"href"`
		} `json:"attributes"`
	} `json:"link"`
	VoteSum struct {
		Label string `json:"label"`
	} `json:"im:voteSum"`
	ContentType struct {
		Attributes struct {
			Term string `json:"term"`
			Label string `json:"label"`
		} `json:"attributes"`
	} `json:"im:contentType"`
	VoteCount struct {
		Label string `json:"label"`
	} `json:"im:voteCount"`
}
