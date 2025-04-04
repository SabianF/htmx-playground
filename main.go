// TODO: Implement basic DB for table data

// Backend solution: https://appwrite.io/

// HTML custom tags: https://matthewjamestaylor.com/custom-tags

// Simple motion animations https://auto-animate.formkit.com/#usage

package main

import (
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	data_repos "github.com/SabianF/htmx-playground/modules/common/data/repositories"
	sources "github.com/SabianF/htmx-playground/modules/common/data/sources"
	delete_row_data_repos "github.com/SabianF/htmx-playground/modules/delete_row/data/repositories"

	bulk_update_data_repos "github.com/SabianF/htmx-playground/modules/bulk_update/data/repositories"
	click_me_data_repos "github.com/SabianF/htmx-playground/modules/click_me/data/repositories"
	click_to_edit_data_repos "github.com/SabianF/htmx-playground/modules/click_to_edit/data/repositories"
	click_to_load_data_repos "github.com/SabianF/htmx-playground/modules/click_to_load/data/repositories"
	hello_example_data_repos "github.com/SabianF/htmx-playground/modules/hello/data/repositories"
	web_socket_data_repos "github.com/SabianF/htmx-playground/modules/web_socket/data/repositories"

	"github.com/joho/godotenv"
)

func main() {
	handleGracefulExit()
	initEnv()

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

func initEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		data_repos.Log("Error loading .env: %v\n", err)
		return
	}

	data_repos.Log("Loaded environment variables.")
}

func exposeEndpoints(mux *http.ServeMux) {
	data_repos.AddRoutes(mux)
	bulk_update_data_repos.AddRoutes(mux)
	click_me_data_repos.AddRoutes(mux)
	click_to_edit_data_repos.AddRoutes(mux)
	click_to_load_data_repos.AddRoutes(mux)
	delete_row_data_repos.AddRoutes(mux)
	hello_example_data_repos.AddRoutes(mux)
	web_socket_data_repos.AddRoutes(mux)
}

func exposeResources(mux *http.ServeMux) {
	mux.Handle("/modules/common/data/sources/assets/", http.StripPrefix("/modules/", http.FileServer(http.Dir("modules"))))
}

func startServer(mux http.Handler) {
	port := ":" + os.Getenv("APP_PORT")
	data_repos.Log("Starting server on port %s ...\n", port)
	err := http.ListenAndServe(port, mux)

	if errors.Is(err, http.ErrServerClosed) {
		data_repos.Log("Server closed\n")

	} else if err != nil {
		data_repos.Log("Error with server: %s\n", err)
		os.Exit(1)
	}
}
