package controllers

type CharactersData struct {
	Characters []Character `json:"characters"`
}

type Character struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Images      []string    `json:"images"`
	Debut       Debut       `json:"debut"`
	Family      Family      `json:"family"`
	Jutsu       []string    `json:"jutsu"`
	NatureType  []string    `json:"natureType"`
	Personal    Personal    `json:"personal"`
	Rank        Rank        `json:"rank"`
	Tools       []string    `json:"tools"`
	VoiceActors VoiceActors `json:"voiceActors"`
}

type Debut struct {
	Manga     string `json:"manga"`
	Anime     string `json:"anime"`
	Novel     string `json:"novel"`
	Movie     string `json:"movie"`
	Game      string `json:"game"`
	OVA       string `json:"ova"`
	AppearsIn string `json:"appearsIn"`
}

type Family struct {
	Father      string `json:"father"`
	Mother      string `json:"mother"`
	Son         string `json:"son,omitempty"`
	Daughter    string `json:"daughter,omitempty"`
	Wife        string `json:"wife,omitempty"`
	AdoptiveSon string `json:"adoptive son,omitempty"`
	Godfather   string `json:"godfather,omitempty"`
	Brother     string `json:"brother,omitempty"`
}

type Age struct {
	PartI           string `json:"Part I"`
	PartII          string `json:"Part II"`
	AcademyGraduate string `json:"Academy Graduate"`
}

type Height struct {
	PartI       string `json:"Part I"`
	PartII      string `json:"Part II"`
	BlankPeriod string `json:"Blank Period"`
}

type Weight struct {
	PartI  string `json:"Part I"`
	PartII string `json:"Part II"`
}

type Personal struct {
	Birthdate      string      `json:"birthdate"`
	Sex            string      `json:"sex"`
	Age            Age         `json:"age"`
	Height         Height      `json:"height"`
	Weight         Weight      `json:"weight"`
	BloodType      string      `json:"bloodType"`
	KekkeiGenkai   interface{} `json:"kekkeiGenkai"`
	Classification interface{} `json:"classification"`
	TailedBeast    string      `json:"tailedBeast"`
	Occupation     interface{} `json:"occupation"`
	Affiliation    interface{} `json:"affiliation"`
	Team           interface{} `json:"team"`
	Clan           interface{} `json:"clan"`
	Titles         []string    `json:"titles"`
}

type NinjaRank struct {
	PartI  string `json:"Part I"`
	Gaiden string `json:"Gaiden,omitempty"`
}

type Rank struct {
	NinjaRank         NinjaRank `json:"ninjaRank"`
	NinjaRegistration string    `json:"ninjaRegistration"`
}

type VoiceActors struct {
	Japanese interface{} `json:"japanese"`
	English  interface{} `json:"english"`
}

type ClansData struct {
	Clans []Clan `json:"clans"`
}

type Clan struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Characters []int  `json:"characters"`
}

type TailedBeastsData struct {
	TailedBeasts []TailedBeast `json:"tailed-beasts"`
}

type TailedBeast struct {
	ID           int                 `json:"id"`
	Name         string              `json:"name"`
	Images       []string            `json:"images"`
	Debut        Debut               `json:"debut"`
	Family       TailedBeastFamily   `json:"family"`
	Jutsu        []string            `json:"jutsu"`
	NatureType   []string            `json:"natureType"`
	Personal     TailedBeastPersonal `json:"personal"`
	UniqueTraits []string            `json:"uniqueTraits"`
}

type TailedBeastFamily struct {
	IncarnationWithTheGodTree string `json:"incarnation with the god tree"`
	DepoweredForm             string `json:"depowered form"`
}

type TailedBeastPersonal struct {
	Status         string      `json:"status"`
	KekkeiGenkai   string      `json:"kekkeiGenkai"`
	Classification interface{} `json:"classification"`
	Jinchuriki     []string    `json:"jinchūriki"`
	Titles         []string    `json:"titles"`
}

type KekkeiGenkaiData struct {
	KekkeiGenkai []KekkeiGenkai `json:"kekkei-genkai"`
}

type KekkeiGenkai struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Characters []int  `json:"characters"`
}
