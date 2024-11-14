
// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go

package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/click-to-edit", getClickToEdit)
	http.HandleFunc("/click-to-edit/edit", getEdit)
	http.HandleFunc("/click-to-edit/save", getSave)
	http.HandleFunc("/click-to-edit/cancel", getCancel)

	const port = ":3333"
	fmt.Printf("Starting server on port %s ...\n", port)
	err := http.ListenAndServe(port, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")

	} else if err != nil {
		fmt.Printf("Error with server: %s\n", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET /\n")

	component := readFile("index.html")

	io.WriteString(w, string(component) + "\n")
}

func getClickToEdit(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET /click-to-edit\n")

	component := readFile("components/click-to-edit/index.html")

	io.WriteString(w, string(component) + "\n")
}

func getEdit(w http.ResponseWriter, r *http.Request	) {
	fmt.Printf("GET /click-to-edit/edit\n")

	io.WriteString(w, readFile("components/click-to-edit/enabled.html"))
}

func getCancel(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET /click-to-edit/save\n")

	io.WriteString(w, readFile("components/click-to-edit/index.html"))
}

func getSave(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET /click-to-edit/save\n")
	fmt.Printf("Request.body: %#v\n", r.Body)

	io.WriteString(w, readFile("components/click-to-edit/index.html"))
}

func readFile(filepathRelative string) string {
	component, err := os.ReadFile(filepathRelative)
	if err != nil {
		panic(err)
	}
	return string(component)
}
