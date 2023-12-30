package app

type CreatedUserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}
