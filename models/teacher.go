package models

type Teacher struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Age            string `json:"age"`
	Department     string `json:"department"`
	IsHead         bool   `json:"is_head"`
	SchoolCode     int    `json:"school_code"`
	IdentityNumber int    `json:"identityNumber"`
}

type TeacherInput struct {
	Name           string `json:"name"`
	Age            string `json:"age"`
	Department     string `json:"department"`
	IsHead         bool   `json:"is_head"`
	SchoolCode     int    `json:"school_code"`
	IdentityNumber int    `json:"identityNumber"`
}