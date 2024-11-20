package helloExample

import "github.com/a-h/templ"

func HelloDiv(name string) templ.Component {
	return hello(name)
}
