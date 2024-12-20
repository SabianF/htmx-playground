package bulk_update

import (
	"net/http"

	bulk_update_pages "github.com/SabianF/htmx-playground/modules/bulk_update/presentation/pages"
)

const ROUTE_PAGE string = "/bulk-update"

func GetPage(w http.ResponseWriter, r *http.Request) {
	bulk_update_pages.BulkUpdatePage(ROUTE_UPDATE).Render(r.Context(), w)
}
