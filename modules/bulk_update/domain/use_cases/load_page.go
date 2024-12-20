package bulk_update

import (
	bulk_update_pages "github.com/SabianF/htmx-playground/modules/bulk_update/presentation/pages"
	"github.com/a-h/templ"
)

func LoadPage(route_update string) templ.Component {
	return bulk_update_pages.BulkUpdatePage(route_update)
}
