package click_me

import (
	click_me "github.com/SabianF/htmx-playground/modules/click_me/presentation/components"
	"github.com/a-h/templ"
)

func Reset(route_reset string) templ.Component {
	return click_me.ClickMeClicked(route_reset)
}
