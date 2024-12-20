package click_to_edit

import (
	click_to_edit_pages "github.com/SabianF/htmx-playground/modules/click_to_edit/presentation/pages"
	"github.com/a-h/templ"
)

func LoadPage(
	route_edit string,
	firstName string,
	lastName string,
	email string,
) templ.Component {
	return click_to_edit_pages.ClickToEditPage(
		route_edit,
		firstName,
		lastName,
		email,
	)
}
