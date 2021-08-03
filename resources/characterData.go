package resources

type Character struct {
	Name               string             `json:"Name"`
	Avatar             string             `json:"Avatar"`
	Bio                string             `json:"Bio"`
	Gender             int                `json:"Gender"`
	ActiveClassJob     ActiveClassJob     `json:"ActiveClassJob"`
	ClassJobs          ClassJobs          `json:"ClassJobs"`
	ClassJobsBozjan    ClassJobsBozjan    `json:"ClassJobsBozjan"`
	ClassJobsElemental ClassJobsElemental `json:"ClassJobsElemental"`
	GearSet            GearSet            `json:"GearSet"`
	Dc                 string             `json:"DC"`
	FreeCompanyID      string             `json:"FreeCompanyId"`
	FreeCompanyName    string             `json:"FreeCompanyName"`
	GuardianDeity      int                `json:"GuardianDeity"`
	ID                 int                `json:"ID"`
	Lang               interface{}        `json:"Lang"`
	Nameday            string             `json:"Nameday"`
	ParseDate          int                `json:"ParseDate"`
	Portrait           string             `json:"Portrait"`
	PvPTeamID          interface{}        `json:"PvPTeamId"`
	Race               int                `json:"Race"`
	Server             string             `json:"Server"`
	Title              int                `json:"Title"`
	TitleTop           bool               `json:"TitleTop"`
	Town               int                `json:"Town"`
	Tribe              int                `json:"Tribe"`
}

type ActiveClassJob struct {
	ClassId       int           `json:"ClassID"`
	ExpLevel      int           `json:"ExpLevel"`
	ExpLevelMax   int           `json:"ExpLevelMax"`
	ExpLevelTogo  int           `json:"ExpLevelTogo"`
	IsSpecialised bool          `json:"IsSpecialised"`
	JobId         int           `json:"JobID"`
	Level         int           `json:"Level"`
	Name          string        `json:"Name"`
	UnlockedState UnlockedState `json:"UnlockedState"`
}

type UnlockedState struct {
	Id   int    `json:"ID"`
	Name string `json:"Name"`
}

type ClassJobs []struct {
	ClassID       int           `json:"ClassID"`
	ExpLevel      int           `json:"ExpLevel"`
	ExpLevelMax   int           `json:"ExpLevelMax"`
	ExpLevelTogo  int           `json:"ExpLevelTogo"`
	IsSpecialised bool          `json:"IsSpecialised"`
	JobID         int           `json:"JobID"`
	Level         int           `json:"Level"`
	Name          string        `json:"Name"`
	UnlockedState UnlockedState `json:"UnlockedState"`
}

type ClassJobsBozjan struct {
	Level  interface{} `json:"Level"`
	Mettle interface{} `json:"Mettle"`
	Name   string      `json:"Name"`
}

type ClassJobsElemental struct {
	ExpLevel     int    `json:"ExpLevel"`
	ExpLevelMax  int    `json:"ExpLevelMax"`
	ExpLevelTogo int    `json:"ExpLevelTogo"`
	Level        int    `json:"Level"`
	Name         string `json:"Name"`
}

type GearSet struct {
	Attributes struct {
		Num1  int `json:"1"`
		Num2  int `json:"2"`
		Num3  int `json:"3"`
		Num4  int `json:"4"`
		Num5  int `json:"5"`
		Num6  int `json:"6"`
		Num7  int `json:"7"`
		Num8  int `json:"8"`
		Num19 int `json:"19"`
		Num20 int `json:"20"`
		Num21 int `json:"21"`
		Num22 int `json:"22"`
		Num24 int `json:"24"`
		Num27 int `json:"27"`
		Num33 int `json:"33"`
		Num34 int `json:"34"`
		Num44 int `json:"44"`
		Num45 int `json:"45"`
		Num46 int `json:"46"`
	} `json:"Attributes"`
	ClassID int `json:"ClassID"`
	Gear    struct {
		Body struct {
			Creator interface{}   `json:"Creator"`
			Dye     interface{}   `json:"Dye"`
			ID      int           `json:"ID"`
			Materia []interface{} `json:"Materia"`
			Mirage  interface{}   `json:"Mirage"`
		} `json:"Body"`
		Bracelets struct {
			Creator interface{}   `json:"Creator"`
			Dye     interface{}   `json:"Dye"`
			ID      int           `json:"ID"`
			Materia []interface{} `json:"Materia"`
			Mirage  interface{}   `json:"Mirage"`
		} `json:"Bracelets"`
		Earrings struct {
			Creator interface{}   `json:"Creator"`
			Dye     interface{}   `json:"Dye"`
			ID      int           `json:"ID"`
			Materia []interface{} `json:"Materia"`
			Mirage  interface{}   `json:"Mirage"`
		} `json:"Earrings"`
		Feet struct {
			Creator interface{}   `json:"Creator"`
			Dye     interface{}   `json:"Dye"`
			ID      int           `json:"ID"`
			Materia []interface{} `json:"Materia"`
			Mirage  interface{}   `json:"Mirage"`
		} `json:"Feet"`
		Hands struct {
			Creator interface{}   `json:"Creator"`
			Dye     interface{}   `json:"Dye"`
			ID      int           `json:"ID"`
			Materia []interface{} `json:"Materia"`
			Mirage  interface{}   `json:"Mirage"`
		} `json:"Hands"`
		Head struct {
			Creator interface{}   `json:"Creator"`
			Dye     interface{}   `json:"Dye"`
			ID      int           `json:"ID"`
			Materia []interface{} `json:"Materia"`
			Mirage  interface{}   `json:"Mirage"`
		} `json:"Head"`
		Legs struct {
			Creator interface{}   `json:"Creator"`
			Dye     interface{}   `json:"Dye"`
			ID      int           `json:"ID"`
			Materia []interface{} `json:"Materia"`
			Mirage  interface{}   `json:"Mirage"`
		} `json:"Legs"`
		MainHand struct {
			Creator interface{}   `json:"Creator"`
			Dye     interface{}   `json:"Dye"`
			ID      int           `json:"ID"`
			Materia []interface{} `json:"Materia"`
			Mirage  interface{}   `json:"Mirage"`
		} `json:"MainHand"`
		Necklace struct {
			Creator interface{}   `json:"Creator"`
			Dye     interface{}   `json:"Dye"`
			ID      int           `json:"ID"`
			Materia []interface{} `json:"Materia"`
			Mirage  interface{}   `json:"Mirage"`
		} `json:"Necklace"`
		Ring1 struct {
			Creator interface{}   `json:"Creator"`
			Dye     interface{}   `json:"Dye"`
			ID      int           `json:"ID"`
			Materia []interface{} `json:"Materia"`
			Mirage  interface{}   `json:"Mirage"`
		} `json:"Ring1"`
		Ring2 struct {
			Creator interface{}   `json:"Creator"`
			Dye     interface{}   `json:"Dye"`
			ID      int           `json:"ID"`
			Materia []interface{} `json:"Materia"`
			Mirage  interface{}   `json:"Mirage"`
		} `json:"Ring2"`
		SoulCrystal struct {
			Creator interface{}   `json:"Creator"`
			Dye     interface{}   `json:"Dye"`
			ID      int           `json:"ID"`
			Materia []interface{} `json:"Materia"`
			Mirage  interface{}   `json:"Mirage"`
		} `json:"SoulCrystal"`
		Waist struct {
			Creator interface{}   `json:"Creator"`
			Dye     interface{}   `json:"Dye"`
			ID      int           `json:"ID"`
			Materia []interface{} `json:"Materia"`
			Mirage  interface{}   `json:"Mirage"`
		} `json:"Waist"`
	} `json:"Gear"`
	GearKey string `json:"GearKey"`
	JobID   int    `json:"JobID"`
	Level   int    `json:"Level"`
}

type GrandCompany struct {
	NameID int `json:"NameID"`
	RankID int `json:"RankID"`
}
