package schema

type Skill struct {
	Category     string `json:"category"`
	CategorySlug string `json:"category_slug"`
	Name         string `json:"name"`
	Value        string `json:"value"`
}

type SkillResponse struct {
	Message string  `json:"message"`
	Data    []Skill `json:"data"`
}
