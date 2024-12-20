package click_to_edit

import (
	click_to_edit_pages "github.com/SabianF/htmx-playground/modules/click_to_edit/presentation/pages"
	"github.com/a-h/templ"
)

type LoadPageArgs struct {
	Route_edit string
	FirstName string
	LastName string
	Email string
}

func LoadPage(args LoadPageArgs) templ.Component {
	return click_to_edit_pages.ClickToEditPage(
		args.Route_edit,
		args.FirstName,
		args.LastName,
		args.Email,
	)
}
