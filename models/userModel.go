package models

type User struct {
	IDUser     int    `json:"id_user"`
	NameUser   string `json:"name_user"`
	EmailUser  string `json:"email_user"`
	SecretUser string `json:"secret_user"`
}
