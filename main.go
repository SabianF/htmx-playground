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
	"net/http"
	"os"
	"os/signal"
	"syscall"

	bulk_update_data_repos "github.com/SabianF/htmx-playground/modules/bulk_update/data/repositories"
	click_me_data_repos "github.com/SabianF/htmx-playground/modules/click_me/data/repositories"
	click_to_edit_data_repos "github.com/SabianF/htmx-playground/modules/click_to_edit/data/repositories"
	click_to_load_data_repos "github.com/SabianF/htmx-playground/modules/click_to_load/data/repositories"
	common_handlers "github.com/SabianF/htmx-playground/modules/common/data/repositories"
	hello_example_data_repos "github.com/SabianF/htmx-playground/modules/hello/data/repositories"
)

const ROUTE_ROOT string = "/"

func main() {
	handleGracefulExit()

	mux := http.NewServeMux()

	exposeEndpoints(mux)
	exposeResources(mux)
	startServer(mux)
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

func exposeEndpoints(mux *http.ServeMux) {
	mux.HandleFunc(ROUTE_ROOT, common_handlers.GetPage)
	mux.HandleFunc(bulk_update_data_repos.ROUTE_PAGE, bulk_update_data_repos.GetPage)
	mux.HandleFunc(bulk_update_data_repos.ROUTE_UPDATE, bulk_update_data_repos.GetUpdate)
	mux.HandleFunc(click_me_data_repos.ROUTE_PAGE, click_me_data_repos.GetPage)
	mux.HandleFunc(click_me_data_repos.ROUTE_CLICKED, click_me_data_repos.GetClicked)
	mux.HandleFunc(click_me_data_repos.ROUTE_RESET, click_me_data_repos.GetReset)
	mux.HandleFunc(click_to_edit_data_repos.ROUTE_PAGE, click_to_edit_data_repos.GetPage)
	mux.HandleFunc(click_to_edit_data_repos.ROUTE_EDIT, click_to_edit_data_repos.GetEdit)
	mux.HandleFunc(click_to_edit_data_repos.ROUTE_SAVE, click_to_edit_data_repos.GetSave)
	mux.HandleFunc(click_to_edit_data_repos.ROUTE_CANCEL, click_to_edit_data_repos.GetCancel)
	mux.HandleFunc(click_to_load_data_repos.ROUTE_PAGE, click_to_load_data_repos.GetPage)
	mux.HandleFunc(click_to_load_data_repos.ROUTE_GET_USERS, click_to_load_data_repos.GetUsers)
	mux.HandleFunc(hello_example_data_repos.ROUTE_PAGE, hello_example_data_repos.GetPage)
}

func exposeResources(mux *http.ServeMux) {
	mux.Handle("/modules/common/data/sources/assets/", http.StripPrefix("/modules/", http.FileServer(http.Dir("modules"))))
}

func startServer(mux *http.ServeMux) {
	const port = ":3333"
	fmt.Printf("Starting server on port %s ...\n", port)
	err := http.ListenAndServe(port, mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")

	} else if err != nil {
		fmt.Printf("Error with server: %s\n", err)
		os.Exit(1)
	}
}
