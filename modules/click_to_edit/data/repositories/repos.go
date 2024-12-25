package click_to_edit

import (
	"net/http"

	use_cases "github.com/SabianF/htmx-playground/modules/click_to_edit/domain/use_cases"
	components "github.com/SabianF/htmx-playground/modules/click_to_edit/presentation/components"
	pages "github.com/SabianF/htmx-playground/modules/click_to_edit/presentation/pages"
)

const ROUTE_PAGE string = "/click-to-edit"
const ROUTE_EDIT string = "/click-to-edit/edit"
const ROUTE_CANCEL string = "/click-to-edit/cancel"
const ROUTE_SAVE string = "/click-to-edit/save"

var clickToEditData = map[string]string{
	"firstName"	: "Joe",
	"lastName"	: "Shmoe",
	"email"			: "joe@shmoe.com",
}

func GetPage(w http.ResponseWriter, r *http.Request) {

	args := use_cases.LoadPageArgs{
		Route_edit: ROUTE_EDIT,
		FirstName:  clickToEditData["firstName"],
		LastName:   clickToEditData["lastName"],
		Email:      clickToEditData["email"],
	}

	use_cases.LoadPage(args).Render(r.Context(), w)
}

func GetEdit(w http.ResponseWriter, r *http.Request) {
	components.ClickToEditForm(
		ROUTE_SAVE,
		clickToEditData["firstName"],
		clickToEditData["lastName"],
		clickToEditData["email"],
	).Render(r.Context(), w)
}

func GetCancel(w http.ResponseWriter, r *http.Request) {
	pages.ClickToEditText(
		ROUTE_EDIT,
		clickToEditData["firstName"],
		clickToEditData["lastName"],
		clickToEditData["email"],
	).Render(r.Context(), w)
}

func GetSave(w http.ResponseWriter, r *http.Request) {
	firstName := r.PostFormValue("firstName")
	if firstName != "" && firstName != clickToEditData["firstName"] {
		clickToEditData["firstName"] = firstName
	}

	lastName := r.PostFormValue("lastName")
	if lastName != "" && lastName != clickToEditData["lastName"] {
		clickToEditData["lastName"] = lastName
	}

	email := r.PostFormValue("email")
	if email != "" && email != clickToEditData["email"] {
		clickToEditData["email"] = email
	}

	pages.ClickToEditText(
		ROUTE_EDIT,
		clickToEditData["firstName"],
		clickToEditData["lastName"],
		clickToEditData["email"],
	).Render(r.Context(), w)
}

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc(ROUTE_PAGE, GetPage)
	mux.HandleFunc(ROUTE_EDIT, GetEdit)
	mux.HandleFunc(ROUTE_SAVE, GetSave)
	mux.HandleFunc(ROUTE_CANCEL, GetCancel)
}
