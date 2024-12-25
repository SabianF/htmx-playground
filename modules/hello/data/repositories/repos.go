package helloExample

import (
	"net/http"

	use_cases "github.com/SabianF/htmx-playground/modules/hello/domain/use_cases"
)

const ROUTE_PAGE string = "/hello"

const DEFAULT_NAME string = "Steven"

func GetPage(w http.ResponseWriter, r *http.Request) {
	use_cases.LoadPage(DEFAULT_NAME).Render(r.Context(), w)
}

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc(ROUTE_PAGE, GetPage)
}
