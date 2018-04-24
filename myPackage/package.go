package myPackage


type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Browsers []string `json:"browsers"`
}
