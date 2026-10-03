package main

import (
	"fmt"
	"net/http"
	"shortner-go/database"
	"shortner-go/handlers"
	
)

/*
var redis1 redis.Client

func hello(w http.ResponseWriter,r *http.Request){
	key:=r.PathValue("key")
	value:=r.PathValue("value")
	redis1.Set(r.Context(),key,value,5*time.Minute)
	fmt.Fprintln(w,"Set sucess")

}

func hello2(w http.ResponseWriter,r *http.Request){
	key:=r.PathValue("key")
	val,err:=redis1.Get(r.Context(),key).Result()
	if err!=nil{
		log.Println(err)
		return
	}
	fmt.Fprintln(w,val)

}*/

func main(){
 db:=database.ConnectDB()
 defer db.Close()

 redis:=database.ConnectRedis()
 defer redis.Close()

 mux:=http.NewServeMux()

handler:=&handlers.Shortner{
	DB: db,
	R: redis,
}

  mux.HandleFunc("GET /v1/url/{shortUrl}",handler.GetUrl)
 mux.HandleFunc("POST /v1/url/",handler.CreateUrl)



 server:=&http.Server{
	Addr: ":8000",
	Handler: mux,
 }

 fmt.Println("Server is live at port 8000")
 server.ListenAndServe()


}