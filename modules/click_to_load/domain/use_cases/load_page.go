package click_to_load

import (
	"github.com/a-h/templ"

	pages "github.com/SabianF/htmx-playground/modules/click_to_load/presentation/pages"
)

func LoadPage(
	route_get_users string,
	next_page string,
	userProps []string,
	usersStringified [][]string,
) templ.Component {
	return pages.ClickToLoad(
		route_get_users,
		next_page,
		userProps,
		usersStringified,
	)
}
