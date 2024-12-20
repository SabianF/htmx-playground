package helloExample

import (
	pages "github.com/SabianF/htmx-playground/modules/hello/presentation/pages"
	"github.com/a-h/templ"
)

func LoadPage(name string) templ.Component {
	return pages.Hello(name)
}
