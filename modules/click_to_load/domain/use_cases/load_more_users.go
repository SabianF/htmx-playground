package click_to_load

import (
	"github.com/a-h/templ"

	click_to_load_components "github.com/SabianF/htmx-playground/modules/click_to_load/presentation/components"
)

func LoadMoreUsers(
	route_get_users string,
	userDataStrings [][]string,
	next_page string,
) templ.Component {
	return click_to_load_components.Rows(
		route_get_users,
		next_page,
		userDataStrings,
	)
}
