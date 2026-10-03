package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shortner-go/models"
	"shortner-go/repository"
	"time"

	"github.com/redis/go-redis/v9"
)
type Shortner struct {
	DB *sql.DB
	R  *redis.Client
}
//short url--->it will give u longurl+redirect
func (h Shortner)GetUrl(w http.ResponseWriter,r *http.Request){
  shortUrl:=r.PathValue("shortUrl")

  cached,err:=h.R.Get(r.Context(),shortUrl).Result()
  fmt.Println("cahched",err)
  if err==nil{
	 w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(cached)
	 return
  }

 result,err1:=repository.Get_short_url(h.DB,shortUrl)

 if err1!=nil{
	http.Error(w,"No Result found",http.StatusNotFound)
	return
 }

 w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(result)

}

func (h Shortner)CreateUrl(w http.ResponseWriter,r *http.Request){
var u models.URL
json.NewDecoder(r.Body).Decode(&u)

result,err:=repository.Create_short(h.DB,u.LongUrl)
if err!=nil{
	fmt.Println(err)
	http.Error(w,"Error ",http.StatusInternalServerError)
	return
}


err=h.R.Set(r.Context(),result.ShortUrl,u.LongUrl,5*time.Minute).Err()
if err!=nil{
	log.Printf("Redis cache error: %v", err)
	return
}
w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(result)
}