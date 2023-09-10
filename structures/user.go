package structures

type SignUpLoginRequest struct {
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Name      string   `json:"name"`
	Company   string   `json:"company"`
	Location  string   `json:"location"`
	Roles     []string `json:"roles"`
	Image     string   `json:"image"`
	DesertRef string   `json:"desertref"`
}

type UpdateUserRequest struct {
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Company   string   `json:"company"`
	Location  string   `json:"location"`
	Roles     []string `json:"roles"`
	Active    bool     `json:"active"`
	Image     string   `json:"image"`
	DesertRef string   `json:"desertref"`
}
