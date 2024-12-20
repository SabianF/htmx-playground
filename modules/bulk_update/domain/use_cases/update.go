package bulk_update

import (
	bulk_update_components "github.com/SabianF/htmx-playground/modules/bulk_update/presentation/components"
	"github.com/a-h/templ"
)

func Update(form_data map[string][]string) templ.Component {
	return bulk_update_components.BulkUpdateToast(form_data)
}
