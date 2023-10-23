package model

type ApiKeys struct {
	ApiKey       string `gorm:"primaryKey" json:"apiKey"`
	PublicKeyPEM string `json:"publicKeyPEM"`
	Hash         string `json:"hash"`
	ClientID     string `json:"clientId"`
}
