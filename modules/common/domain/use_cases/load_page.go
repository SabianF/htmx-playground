package common

import (
	"github.com/a-h/templ"

	entities "github.com/SabianF/htmx-playground/modules/common/domain/entities"
	pages "github.com/SabianF/htmx-playground/modules/common/presentation/pages"
)

func GetHomePage(list_of_pages []entities.Page) templ.Component {
	return pages.HomePage(list_of_pages)
}
