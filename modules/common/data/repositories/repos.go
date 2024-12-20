package common

import (
	"net/http"

	common_use_cases "github.com/SabianF/htmx-playground/modules/common/domain/use_cases"
)

func GetPage(w http.ResponseWriter, r *http.Request) {
	homePage := common_use_cases.GetHomePage()

	homePage.Render(r.Context(), w)
}
