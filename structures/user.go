package structures

type SignUpLoginRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	DesertRef string `json:"desertref"`
}

type UpdateUserRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Image     string `json:"image"`
	DesertRef string `json:"desertref"`
}
