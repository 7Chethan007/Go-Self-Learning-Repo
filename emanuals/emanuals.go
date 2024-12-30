package main

import (
	"fmt"
	"net/http"
)

// handler handles HTTP requests and sends a response.
// It takes two parameters:
// - w: an http.ResponseWriter which is used to send the response back to the client.
// - r: an http.Request which contains all the information about the request made by the client.
//
// Inside the function:
// - It prints the received URL path to the console using fmt.Println.
// - It sends an HTML response with a welcome message using fmt.Fprintf.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieved: " + r.URL.Path)
	fmt.Fprintf(w, "<h1>Welocome To Night Owl</h1>")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port 3000....")
	http.ListenAndServe(":3000", nil)
}
