package click_to_load

import (
	"github.com/a-h/templ"

	click_to_load_components "github.com/SabianF/htmx-playground/modules/click_to_load/presentation/components"
)

type LoadMoreUsersArgs struct {
	Route_get_users string
	Next_page string
	UserDataStrings [][]string
}

func LoadMoreUsers(args LoadMoreUsersArgs) templ.Component {
	return click_to_load_components.Rows(
		args.Route_get_users,
		args.Next_page,
		args.UserDataStrings,
	)
}
