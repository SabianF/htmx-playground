package common

import (
	"net/http"

	sources "github.com/SabianF/htmx-playground/modules/common/data/sources"
	common_use_cases "github.com/SabianF/htmx-playground/modules/common/domain/use_cases"
)

func NewLogger(handler http.Handler) *sources.Logger {
	return &sources.Logger{Handler: handler}
}

func GetPage(w http.ResponseWriter, r *http.Request) {
	homePage := common_use_cases.GetHomePage()

	homePage.Render(r.Context(), w)
}
