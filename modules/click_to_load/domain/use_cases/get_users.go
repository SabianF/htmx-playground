package click_to_load

import (
	"fmt"
	"net/http"

	click_to_load "github.com/SabianF/htmx-playground/modules/click_to_load/domain/repositories"
	click_to_load_pages "github.com/SabianF/htmx-playground/modules/click_to_load/presentation/pages"
)

func ServePageWithInitialData(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	const INITIAL_PAGE = 1
	const DEFAULT_PAGINATION = 10

	userDataHeadings := click_to_load.GetUserProps()
	userDataRows := click_to_load.GetUsers(INITIAL_PAGE, DEFAULT_PAGINATION)

	component := click_to_load_pages.ClickToLoad(userDataHeadings, userDataRows)

	component.Render(r.Context(), w)
}
