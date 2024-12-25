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
// TODO: Easy backend solution: https://appwrite.io/
// TODO: Simple motion animations https://auto-animate.formkit.com/#usage

package main

import (
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	data_repos "github.com/SabianF/htmx-playground/modules/common/data/repositories"
	sources "github.com/SabianF/htmx-playground/modules/common/data/sources"

	bulk_update_data_repos "github.com/SabianF/htmx-playground/modules/bulk_update/data/repositories"
	click_me_data_repos "github.com/SabianF/htmx-playground/modules/click_me/data/repositories"
	click_to_edit_data_repos "github.com/SabianF/htmx-playground/modules/click_to_edit/data/repositories"
	click_to_load_data_repos "github.com/SabianF/htmx-playground/modules/click_to_load/data/repositories"
	hello_example_data_repos "github.com/SabianF/htmx-playground/modules/hello/data/repositories"
)

const ROUTE_ROOT string = "/"

func main() {
	handleGracefulExit()

	mux := http.NewServeMux()

	exposeEndpoints(mux)
	exposeResources(mux)

	logger := sources.Notify()

	mux_wrapped := data_repos.AddMiddleware(
		mux,
		logger,
	)

	startServer(mux_wrapped)
}

func handleGracefulExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		data_repos.Log("Received SIGTERM. Exiting...\n")
		os.Exit(1)
	}()
}

func exposeEndpoints(mux *http.ServeMux) {
	mux.HandleFunc(ROUTE_ROOT, data_repos.GetPage)
	bulk_update_data_repos.AddRoutes(mux)
	click_me_data_repos.AddRoutes(mux)
	click_to_edit_data_repos.AddRoutes(mux)
	mux.HandleFunc(click_to_load_data_repos.ROUTE_PAGE, click_to_load_data_repos.GetPage)
	mux.HandleFunc(click_to_load_data_repos.ROUTE_GET_USERS, click_to_load_data_repos.GetUsers)
	mux.HandleFunc(hello_example_data_repos.ROUTE_PAGE, hello_example_data_repos.GetPage)
}

func exposeResources(mux *http.ServeMux) {
	mux.Handle("/modules/common/data/sources/assets/", http.StripPrefix("/modules/", http.FileServer(http.Dir("modules"))))
}

func startServer(mux http.Handler) {
	const port = ":3333"
	data_repos.Log("Starting server on port %s ...\n", port)
	err := http.ListenAndServe(port, mux)

	if errors.Is(err, http.ErrServerClosed) {
		data_repos.Log("Server closed\n")

	} else if err != nil {
		data_repos.Log("Error with server: %s\n", err)
		os.Exit(1)
	}
}
