package main

import (
	"fmt"
	"html/template"
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

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("recieved: " + r.URL.Path)
// 	// fmt.Fprintf(w, "<h1>Welocome To Night Owl</h1>")
// 	t, _ := template.ParseFiles("templates/index.html")
// 	t.Execute(w, nil)
// }

// We are using a library called "html/template" to render HTML templates.
// The template.ParseFiles function reads the content of the file and returns a pointer to a template.
// The template.Execute function is used to render the template and send the output to the http.ResponseWriter.
// Right now no error handling is done, but in a real-world application, you should handle errors properly.

// Now lets handle the erros
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieved: " + r.URL.Path)
	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%v Server Error\n", http.StatusNotFound)
		fmt.Fprintf(w, "Description: %s\n", err)
		return
	}
	pages, _ := scandir("./manuals")
	fmt.Println(pages)
	// scandir is a function that reads the contents of a directory and returns a list of file names.

	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port 3000....")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
