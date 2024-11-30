package click_to_load

import (
	"fmt"
	"net/http"
	"strconv"

	click_to_load "github.com/SabianF/htmx-playground/modules/click_to_load/domain/repositories"
	click_to_load_pages "github.com/SabianF/htmx-playground/modules/click_to_load/presentation/pages"
)

func ServePageWithInitialData(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	const INITIAL_PAGE = 1
	const INITIAL_PAGINATION = 10

	userDataHeadings := click_to_load.GetUserProps()
	userDataRows := click_to_load.GetUsers(INITIAL_PAGE, INITIAL_PAGINATION)

	component := click_to_load_pages.ClickToLoad(userDataHeadings, userDataRows)

	component.Render(r.Context(), w)
}

func LoadMoreUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	const DEFAULT_PAGINATION = 20

	// Parse query
	pageNumQuery := r.URL.Query().Get("page")

	pageNum, err := strconv.Atoi(pageNumQuery)
	if err != nil {
		panic("Error parsing query: page")
	}

	// Make DB request
	userDataStrings := click_to_load.GetUsers(pageNum, DEFAULT_PAGINATION)

	// Render components
	nextPageNumString := strconv.Itoa(pageNum + 1)
	component := click_to_load_pages.ClickToLoadRows(userDataStrings, nextPageNumString)

	// Send response
	component.Render(r.Context(), w)
}
