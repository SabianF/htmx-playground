package delete_row

import (
	pages "github.com/SabianF/htmx-playground/modules/delete_row/presentation/pages"
	"github.com/a-h/templ"
)

func LoadPage(
	table_headings []string,
	table_rows [][]string,
	button_delete_route string,
) templ.Component {

	return pages.DeleteRowPage(
		table_headings,
		table_rows,
		button_delete_route,
	)
}
