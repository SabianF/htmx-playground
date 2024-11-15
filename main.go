
// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go

package main

import (
	"errors"
	"fmt"
	"github.com/a-h/templ"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/click-me", getClickMe)
	http.HandleFunc("/click-me/clicked", getClickMeClicked)
	http.HandleFunc("/click-me/reset", getClickMeReset)
	http.HandleFunc("/click-to-edit", getClickToEdit)
	http.HandleFunc("/click-to-edit/edit", getEdit)
	http.HandleFunc("/click-to-edit/save", getSave)
	http.HandleFunc("/click-to-edit/cancel", getCancel)

	component := hello("Stephen")
	http.Handle("/hello", templ.Handler(component))

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
	fmt.Printf("%s %s\n", r.Method, r.URL)

	component := readFile("index.html")

	io.WriteString(w, string(component) + "\n")
}

func getClickMe(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	io.WriteString(w, readFile("components/click-me/index.html"))
}

func getClickMeReset(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	io.WriteString(w, readFile("components/click-me/reset.html"))
}

func getClickMeClicked(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	io.WriteString(w, readFile("components/click-me/clicked.html"))
}

func getClickToEdit(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	component := readFile("components/click-to-edit/index.html")

	io.WriteString(w, string(component) + "\n")
}

func getEdit(w http.ResponseWriter, r *http.Request	) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	io.WriteString(w, readFile("components/click-to-edit/enabled.html"))
}

func getCancel(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	io.WriteString(w, readFile("components/click-to-edit/index.html"))
}

func getSave(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)
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
