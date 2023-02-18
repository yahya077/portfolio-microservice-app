package schema

type Social struct {
	IconURL        string `json:"icon_url"`
	Name           string `json:"name"`
	PrettifiedName string `json:"prettified_name"`
	ProfilePath    string `json:"profile_path"`
	RootPath       string `json:"root_path"`
	Slug           string `json:"slug"`
	Username       string `json:"username"`
}

type SocialResponse struct {
	Message string   `json:"message"`
	Data    []Social `json:"data"`
}
