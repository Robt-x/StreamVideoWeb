package defs

// requests
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// Data model
type VideoInfo struct {
	Id           string
	Author_id    int
	name         string
	DisplayCtime string
}
