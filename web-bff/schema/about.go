package schema

type About struct {
	Contact Contact  `json:"contact"`
	Skills  []Skill  `json:"skills"`
	Socials []Social `json:"socials"`
}
