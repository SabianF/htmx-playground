package common

import (
	"github.com/a-h/templ"
	common_pages "github.com/SabianF/htmx-playground/modules/common/presentation/pages"
)

func GetHomePage() templ.Component {
	return common_pages.HomePage()
}
