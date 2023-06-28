package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

func main() {

	http.HandleFunc("/", pageVists)
	fmt.Println("Http server started on port 9003")
	http.ListenAndServe(":9003", nil)

}

func pageVists(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	//Declare and Initialise page views

	client := redis.NewClient(&redis.Options{
		Addr:     "redis-server:6379",
		Password: "",
		DB:       0,
	})
	//data, _ := client.Get("visits").Result()

	data := client.Incr("visits")
	fmt.Fprintf(w, "Page views are %s", data)

}
