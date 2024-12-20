package click_me

import (
	"github.com/a-h/templ"
	click_me "github.com/SabianF/htmx-playground/modules/click_me/presentation/pages"
)

func LoadPage(route_clicked string) templ.Component {
	return click_me.ClickMePage(route_clicked)
}
