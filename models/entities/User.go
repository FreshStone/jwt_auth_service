package entities

type User struct {
	Email                  string `json:"user_email"`
	Password               string `json:"user_password"`
	AccessToken            string `json:"access_token"`
	RefreshToken           string `json:"refresh_token"`
}
