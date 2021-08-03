package resources

type ItemSearch struct {
	ID   int    `json:"ID"`
	Icon string `json:"Icon"`
	Name string `json:"Name"`
	Url  string `json:"Url"`
}

type ItemSearchResponse struct {
	Results []ItemSearch `json:"Results"`
}
