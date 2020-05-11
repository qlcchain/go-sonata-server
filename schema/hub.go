package schema

type HubSubscriber struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Query    string `json:"query"`
	Callback string `json:"callback"`
}
