// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go

// TODO: Impl hot reloading: https://medium.com/ostinato-rigore/go-htmx-templ-tailwind-complete-project-setup-hot-reloading-2ca1ba6c28be
// TODO: Impl middleware: https://drstearns.github.io/tutorials/gomiddleware/

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
	http.HandleFunc("/click-me", getClickMe)
	http.HandleFunc("/click-me/clicked", getClickMeClicked)
	http.HandleFunc("/click-me/reset", getClickMeReset)
	http.HandleFunc("/click-to-edit", getClickToEdit)
	http.HandleFunc("/click-to-edit/edit", getEdit)
	http.HandleFunc("/click-to-edit/save", getSave)
	http.HandleFunc("/click-to-edit/cancel", getCancel)
	http.HandleFunc("/hello", getHello)

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

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	hello("Stephen").Render(r.Context(), w)
}

func getClickMe(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	clickMePage().Render(r.Context(), w)
}

func getClickMeReset(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	clickMeButton().Render(r.Context(), w)
}

func getClickMeClicked(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	clickMeClicked().Render(r.Context(), w)
}

var clickToEditData = map[string]string{
	"firstName"	: "Joe",
	"lastName"	: "Blow",
	"email"			: "joe@blow.com",
}

func getClickToEdit(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	clickToEditPage(
		clickToEditData["firstName"],
		clickToEditData["lastName"],
		clickToEditData["email"],
	).Render(r.Context(), w)
}

func getEdit(w http.ResponseWriter, r *http.Request	) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	clickToEditForm(
		clickToEditData["firstName"],
		clickToEditData["lastName"],
		clickToEditData["email"],
	).Render(r.Context(), w)
}

func getCancel(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	clickToEditText(
		clickToEditData["firstName"],
		clickToEditData["lastName"],
		clickToEditData["email"],
	).Render(r.Context(), w)
}

func getSave(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	firstName := r.PostFormValue("firstName")
	if firstName != "" && firstName != clickToEditData["firstName"] {
		clickToEditData["firstName"] = firstName
	}

	lastName := r.PostFormValue("lastName")
	if lastName != "" && lastName != clickToEditData["lastName"] {
		clickToEditData["lastName"] = lastName
	}

	email := r.PostFormValue("email")
	if email != "" && email != clickToEditData["email"] {
		clickToEditData["email"] = email
	}

	clickToEditText(
		clickToEditData["firstName"],
		clickToEditData["lastName"],
		clickToEditData["email"],
	).Render(r.Context(), w)
}

func readFile(filepathRelative string) string {
	component, err := os.ReadFile(filepathRelative)
	if err != nil {
		panic(err)
	}
	return string(component)
}
