// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go

// - HTMX + Go + Templ: https://templ.guide/server-side-rendering/htmx/#installation
// - HTMX + Go
// 	- Part 1 (Templating): https://medium.com/gravel-engineering/this-blogpost-also-posted-in-my-personal-blog-which-you-can-access-here-dd856c61001
// 	- Part 2 (TailwindCSS): https://medium.com/gravel-engineering/personal-blog-with-htmx-go-part-2-integrating-tailwindcss-412ebc4dcc97
// 	- Part 3 (Server-rendered markdown):https://medium.com/@mwyndham/personal-blog-with-htmx-go-part-3-server-rendered-markdown-75c80cc5f470
// - HTML custom tags: https://matthewjamestaylor.com/custom-tags
// - TailwindCSS UI Kit: https://www.hyperui.dev/

// Go - Serving static files: https://stackoverflow.com/a/43425767

// TODO: GoTTH stack: https://www.youtube.com/watch?v=k00jVJeZxrs
// TODO: Impl middleware: https://drstearns.github.io/tutorials/gomiddleware/

package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	bulk_update_data_repos "github.com/SabianF/htmx-playground/modules/bulk_update/data/repositories"
	click_me_components "github.com/SabianF/htmx-playground/modules/click_me/presentation/components"
	click_me_pages "github.com/SabianF/htmx-playground/modules/click_me/presentation/pages"
	click_to_edit_components "github.com/SabianF/htmx-playground/modules/click_to_edit/presentation/components"
	click_to_edit_pages "github.com/SabianF/htmx-playground/modules/click_to_edit/presentation/pages"
	click_to_load_use_cases "github.com/SabianF/htmx-playground/modules/click_to_load/domain/use_cases"
	common_handlers "github.com/SabianF/htmx-playground/modules/common/data/repositories"
	hello_pages "github.com/SabianF/htmx-playground/modules/hello/presentation/pages"

	"net/http"
)

func main() {
	handleGracefulExit()
	exposeEndpoints()
	exposeResources()
	startServer()
}

func handleGracefulExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("Received SIGTERM. Exiting...\n")
		os.Exit(1)
	}()
}

func exposeEndpoints() {

	// TODO: Create routes in each module, then import all here to expose

	http.HandleFunc("/", common_handlers.GetRoot)
	http.HandleFunc(bulk_update_data_repos.ROUTE_PAGE, bulk_update_data_repos.GetPage)
	http.HandleFunc(bulk_update_data_repos.ROUTE_UPDATE, bulk_update_data_repos.GetUpdate)
	http.HandleFunc("/click-me", getClickMe)
	http.HandleFunc("/click-me/clicked", getClickMeClicked)
	http.HandleFunc("/click-me/reset", getClickMeReset)
	http.HandleFunc("/click-to-edit", getClickToEdit)
	http.HandleFunc("/click-to-edit/edit", getEdit)
	http.HandleFunc("/click-to-edit/save", getSave)
	http.HandleFunc("/click-to-edit/cancel", getCancel)
	http.HandleFunc("/click-to-load", click_to_load_use_cases.ServePageWithInitialData)
	http.HandleFunc("/click-to-load/", click_to_load_use_cases.LoadMoreUsers)
	http.HandleFunc("/hello", getHello)
}

func exposeResources() {
	http.Handle("/modules/common/data/sources/assets/", http.StripPrefix("/modules/", http.FileServer(http.Dir("modules"))))
}

func startServer() {
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

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	hello_pages.Hello("Stephen").Render(r.Context(), w)
}

func getClickMe(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	click_me_pages.ClickMePage().Render(r.Context(), w)
}

func getClickMeReset(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	click_me_components.ClickMeButton().Render(r.Context(), w)
}

func getClickMeClicked(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	click_me_components.ClickMeClicked().Render(r.Context(), w)
}

var clickToEditData = map[string]string{
	"firstName"	: "Joe",
	"lastName"	: "Blow",
	"email"			: "joe@blow.com",
}

func getClickToEdit(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	click_to_edit_pages.ClickToEditPage(
		clickToEditData["firstName"],
		clickToEditData["lastName"],
		clickToEditData["email"],
	).Render(r.Context(), w)
}

func getEdit(w http.ResponseWriter, r *http.Request	) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	click_to_edit_components.ClickToEditForm(
		clickToEditData["firstName"],
		clickToEditData["lastName"],
		clickToEditData["email"],
	).Render(r.Context(), w)
}

func getCancel(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	click_to_edit_pages.ClickToEditText(
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

	click_to_edit_pages.ClickToEditText(
		clickToEditData["firstName"],
		clickToEditData["lastName"],
		clickToEditData["email"],
	).Render(r.Context(), w)
}
