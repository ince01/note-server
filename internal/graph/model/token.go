package model

import "time"

type Token struct {
	TokenType    TokenType `json:"tokenType"`
	AccessToken  string    `json:"accessToken"`
	ExpiresIn    time.Time `json:"expiresIn"`
	RefreshToken *string   `json:"refreshToken"`
}

type UserCredential struct {
	GrantType GrantType `json:"grantType"`
	UserName  string    `json:"userName"`
	Password  string    `json:"password"`
	Scope     *string   `json:"scope"`
}
