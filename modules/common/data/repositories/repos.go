package common

import (
	"net/http"

	sources "github.com/SabianF/htmx-playground/modules/common/data/sources"
	common_use_cases "github.com/SabianF/htmx-playground/modules/common/domain/use_cases"
)

func AddMiddleware(handler http.Handler, middlewares ...sources.Adapter) http.Handler {

	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}

func Log(format string, v ...any) {
	sources.Printf(format, v...)
}

func GetPage(w http.ResponseWriter, r *http.Request) {
	homePage := common_use_cases.GetHomePage()

	homePage.Render(r.Context(), w)
}
