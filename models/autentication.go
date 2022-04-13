package models

type DataAutentication struct {
	Name      string `json:"name"`
	AccountID string `json:"id"`
	Token     string `json:"token"`
}
