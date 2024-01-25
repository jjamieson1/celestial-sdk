package models

type Token struct {

  AccessToken   string  `json:"access_token"`
  ExpiresIn     int32   `json:"expires_in"`
  RefreshToken  string  `json:"refresh_token"`

}
