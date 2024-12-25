package delete_row

import (
	"net/http"

	use_cases "github.com/SabianF/htmx-playground/modules/delete_row/domain/use_cases"
)

const ROUTE_PAGE string = "/delete-row"
const ROUTE_DELETE string = "/delete-row/delete"

var TABLE_HEADINGS = []string{
	"Name",
	"Email",
	"Status",
}

var table_rows = [][]string{
	{ "Nancy", "nancy@gmail.com", "Active" },
	{ "John", "john@gmail.com", "Inactive" },
	{ "Stephanie", "steph@gmail.com", "Active" },
	{ "George", "gg@gmail.com", "Active" },
}

func GetPage(w http.ResponseWriter, r *http.Request) {

	page := use_cases.LoadPage(
		TABLE_HEADINGS,
		table_rows,
		ROUTE_DELETE,
	)

	page.Render(r.Context(), w)
}

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc(ROUTE_PAGE, GetPage)
}
