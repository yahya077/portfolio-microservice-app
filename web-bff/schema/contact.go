package schema

type Contact struct {
	AboutDescription string `json:"about_description"`
	Address          string `json:"address:"`
	Degree           string `json:"degree"`
	Email            string `json:"email"`
	FirstName        string `json:"first_name"`
	FullName         string `json:"full_name"`
	JobTitle         string `json:"job_title"`
	LastName         string `json:"last_name"`
	MapFrameURL      string `json:"map_frame_url"`
	Phone            string `json:"phone"`
}

type ContactResponse struct {
	Message string  `json:"message"`
	Data    Contact `json:"data"`
}
