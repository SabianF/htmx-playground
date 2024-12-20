package bulk_update

import (
	"net/http"

	bulk_update_components "github.com/SabianF/htmx-playground/modules/bulk_update/presentation/components"
)

const ROUTE_UPDATE string = "/bulk-update/submit"

func Update(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data := map[string][]string(r.PostForm)

	bulk_update_components.BulkUpdateToast(data).Render(r.Context(), w)
}
