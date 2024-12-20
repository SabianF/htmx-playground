package bulk_update

import (
	"net/http"

	bulk_update_use_cases "github.com/SabianF/htmx-playground/modules/bulk_update/domain/use_cases"
)

const ROUTE_PAGE string = "/bulk-update"

func GetPage(w http.ResponseWriter, r *http.Request) {
	bulk_update_use_cases.LoadPage(ROUTE_UPDATE).Render(r.Context(), w)
}

const ROUTE_UPDATE string = "/bulk-update/submit"

func GetUpdate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var form_data map[string][]string = r.PostForm
	bulk_update_use_cases.Update(form_data).Render(r.Context(), w)
}
