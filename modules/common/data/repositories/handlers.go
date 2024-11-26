package common_handlers

import (
	"fmt"
	"net/http"

	"github.com/SabianF/htmx-playground/modules/common/presentation/pages"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	common_pages.HomePage().Render(r.Context(), w)
}
