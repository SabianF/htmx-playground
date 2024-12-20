package click_to_load

import (
	"github.com/a-h/templ"

	pages "github.com/SabianF/htmx-playground/modules/click_to_load/presentation/pages"
)

type LoadPageArgs struct {
	Route_get_users string
	Next_page string
	UserProps []string
	UsersStringified [][]string
}

func LoadPage(args LoadPageArgs) templ.Component {
	return pages.ClickToLoad(
		args.Route_get_users,
		args.Next_page,
		args.UserProps,
		args.UsersStringified,
	)
}
