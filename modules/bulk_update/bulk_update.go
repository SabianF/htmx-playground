package bulk_update

import (
	bulk_update_components "github.com/SabianF/htmx-playground/modules/bulk_update/presentation/components"
	bulk_update_pages "github.com/SabianF/htmx-playground/modules/bulk_update/presentation/pages"
	"github.com/a-h/templ"
)

func BulkUpdate() templ.Component {
	return bulk_update_pages.BulkUpdatePage()
}

func BulkUpdateToast(updatedUsers map[string][]string) templ.Component {
	return bulk_update_components.BulkUpdateToast(updatedUsers)
}
