package common

import (
	"net/http"

	sources "github.com/SabianF/htmx-playground/modules/common/data/sources"
	common_use_cases "github.com/SabianF/htmx-playground/modules/common/domain/use_cases"
)

const ROUTE_PAGE string = "/"

func AddMiddleware(handler http.Handler, middlewares ...sources.Adapter) http.Handler {

	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}

func Log(format string, v ...any) {
	sources.Printf(format, v...)
}

func getPage(w http.ResponseWriter, r *http.Request) {

	list_of_pages := GetAllPages()

	homePage := common_use_cases.GetHomePage(list_of_pages)

	homePage.Render(r.Context(), w)
}

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc(ROUTE_PAGE, getPage)
}
