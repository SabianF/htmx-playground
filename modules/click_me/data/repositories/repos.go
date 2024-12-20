package click_me

import (
	"net/http"

	click_me "github.com/SabianF/htmx-playground/modules/click_me/presentation/components"
	click_me_pages "github.com/SabianF/htmx-playground/modules/click_me/presentation/pages"
)

const ROUTE_PAGE string = "/click-me"
const ROUTE_CLICKED string = "/click-me/clicked"
const ROUTE_RESET string = "/click-me/reset"

func GetPage(w http.ResponseWriter, r *http.Request) {
	click_me_pages.ClickMePage(ROUTE_CLICKED).Render(r.Context(), w)
}

func GetClicked(w http.ResponseWriter, r *http.Request) {
	click_me.ClickMeClicked(ROUTE_RESET).Render(r.Context(), w)
}

func GetReset(w http.ResponseWriter, r *http.Request) {
	click_me.ClickMeButton(ROUTE_CLICKED).Render(r.Context(), w)
}
