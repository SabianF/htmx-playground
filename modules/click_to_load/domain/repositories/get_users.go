package click_to_load

import (
	"reflect"

	click_to_load_models "github.com/SabianF/htmx-playground/modules/click_to_load/data/models"
	click_to_load_repos "github.com/SabianF/htmx-playground/modules/click_to_load/data/repositories"
)

func GetUserProps() []string {

	var userProps []string
	userStruct := click_to_load_models.User{}
	intermediate := reflect.Indirect(reflect.ValueOf(userStruct))

	for i := 0; i < intermediate.NumField(); i++ {
		userProps = append(userProps, intermediate.Type().Field(i).Name)
	}

	return userProps
}

func GetUsers(page int, pagination_amount int) [][]string {

	var allUserDataStrings [][]string

	userDataModels := click_to_load_repos.GetUsers(page, pagination_amount)

	for _, user := range userDataModels {

		currentUserReflect := reflect.ValueOf(user)
		currentUserPropsStrings := make([]string, currentUserReflect.NumField())

		for i := 0; i < currentUserReflect.NumField(); i++ {

			currentUserPropString := currentUserReflect.Field(i).Interface().(string)
			currentUserPropsStrings[i] = currentUserPropString
		}

		allUserDataStrings = append(allUserDataStrings, currentUserPropsStrings)
	}

	return allUserDataStrings
}
