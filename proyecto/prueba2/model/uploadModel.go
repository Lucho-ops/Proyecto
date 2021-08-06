package model

type Upload struct {
	Organization string `json:"organization"`
	User         User   `json:"users"`
}

type User struct {
	Name  string   `json:"username"`
	Roles []string `json:"roles"`
}
