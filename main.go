package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type user struct {
	Name string
	Age  int
}

func handleRequests() {
	http.HandleFunc("/Hello", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		fmt.Println(t.Month())
		fmt.Println(t.Day())
		fmt.Println(t.Year())
		ua := r.Header.Get("User-Agent")
		fmt.Printf("user agent is: %s \n", ua)
		invocationid := r.Header.Get("X-Azure-Functions-InvocationId")
		fmt.Printf("invocationid is: %s \n", invocationid)
		json.NewEncoder(w).Encode(user{
			Name: "Kazuki from /hello",
			Age:  39,
		})
	})

	httpInvokerPort, exists := os.LookupEnv("FUNCTIONS_HTTPWORKER_PORT")
	if exists {
		fmt.Println("FUNCTIONS_HTTPWORKER_PORT: " + httpInvokerPort)
	}
	log.Println("Go server Listening...on httpInvokerPort:", httpInvokerPort)
	log.Fatal(http.ListenAndServe(":"+httpInvokerPort, nil))
}

func main() {
	handleRequests()
}
