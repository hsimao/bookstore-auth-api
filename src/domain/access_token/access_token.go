package access_token

import (
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"` // 判斷訪問裝置用 brower、ios、android...
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	// now := time.Now().UTC()
	// expirationTime := time.Unix(at.Expires, 0)
	// return now.After(expirationTime)

	// 改寫成一行
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
