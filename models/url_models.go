package models

import "time"

type URL struct {
	LongUrl  string `json:"longUrl"`
	ShortUrl string `json:"shortUrl"`
	ExpireAt time.Time `json:"expireAt"`
}