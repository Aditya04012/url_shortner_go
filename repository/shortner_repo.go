package repository

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"math/big"
	"shortner-go/models"
	"time"
)

func generateShortURL() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 6)

	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			panic(err)
		}
		b[i] = chars[n.Int64()]
	}

	return string(b)
}

func Get_short_url(db *sql.DB, shortUrl string) (models.URL, error) {
    var u models.URL

    query := `
        SELECT "shortUrl", "longUrl", "expireAt"
        FROM shortner
        WHERE "shortUrl" = $1
          AND "expireAt" > NOW()
    `

    err := db.QueryRow(query, shortUrl).Scan(&u.ShortUrl,
        &u.LongUrl,
        &u.ExpireAt,)
	
    if err != nil {	
		fmt.Println("Database error:", err)
        return models.URL{}, err
		
    }
fmt.Println("hii",u.LongUrl)
    return u, nil
}

func Create_short(db *sql.DB,longUrl string)(models.URL,error){
	var u models.URL
	u.LongUrl= longUrl
	u.ExpireAt = time.Now().Add(90 * 24 * time.Hour)
	u.ShortUrl = generateShortURL()

	
	query := `
    INSERT INTO shortner ("shortUrl", "longUrl", "expireAt")
    VALUES ($1, $2, $3)
    RETURNING "shortUrl", "longUrl", "expireAt"
`
	err:=db.QueryRow(
		query,
		u.ShortUrl,
		u.LongUrl,
		u.ExpireAt,
	).Scan(&u.ShortUrl, &u.LongUrl, &u.ExpireAt)

	if err!=nil{
		return models.URL{},err	
	}
	return u,nil
}



