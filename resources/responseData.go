package resources

type CharResponse struct {
	Character          Character   `json:"Character"`
	FreeCompany        interface{} `json:"FreeCompany"`
	FreeCompanyMembers interface{} `json:"FreeCompanyMembers"`
	Friends            interface{} `json:"Friends"`
	FriendsPublic      interface{} `json:"FriendsPublic"`
	Minions            interface{} `json:"Minions"`
	Mounts             interface{} `json:"Mounts"`
	PvPTeam            interface{} `json:"PvPTeam"`
}

type CharacterSearchResponse struct {
	Results []Results `json:"Results"`
}

type Results struct {
	ID     int    `json:"ID"`
	Name   string `json:"Name"`
	Server string `json:"Server"`
}
