package click_to_load

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	sources "github.com/SabianF/htmx-playground/modules/click_to_load/data/sources"
	entities "github.com/SabianF/htmx-playground/modules/click_to_load/domain/entities"
	use_cases "github.com/SabianF/htmx-playground/modules/click_to_load/domain/use_cases"
)

const ROUTE_PAGE string = "/click-to-load"

const DEFAULT_PAGE int = 1
const DEFAULT_PAGINATION int = 10

func GetPage(w http.ResponseWriter, r *http.Request) {

	next_page := strconv.Itoa(DEFAULT_PAGE + 1)
	userProps := GetUserProps()
	users := fetchUserData(DEFAULT_PAGE, DEFAULT_PAGINATION)
	usersStringified := stringifyUsers(users)

	args := use_cases.LoadPageArgs{
		Route_get_users:  ROUTE_GET_USERS,
		Next_page:        next_page,
		UserProps:        userProps,
		UsersStringified: usersStringified,
	}

	use_cases.LoadPage(args).Render(r.Context(), w)
}

const ROUTE_GET_USERS string = "/click-to-load/users"

func GetUsers(w http.ResponseWriter, r *http.Request) {
	pageNumQuery := r.URL.Query().Get("page")

	pageNum, err := strconv.Atoi(pageNumQuery)
	if err != nil {
		panic("Error parsing query: page")
	}

	users := fetchUserData(pageNum, DEFAULT_PAGINATION)
	usersStringified := stringifyUsers(users)
	nextPageNumString := strconv.Itoa(pageNum + 1)

	args := use_cases.LoadMoreUsersArgs{
		Route_get_users: ROUTE_GET_USERS,
		Next_page: nextPageNumString,
		UserDataStrings: usersStringified,
	}

	use_cases.LoadMoreUsers(args).Render(r.Context(), w)
}

func GetUserProps() []string {

	var userProps []string
	userStruct := entities.User{}
	intermediate := reflect.Indirect(reflect.ValueOf(userStruct))

	for i := 0; i < intermediate.NumField(); i++ {
		userProps = append(userProps, intermediate.Type().Field(i).Name)
	}

	return userProps
}

func fetchUserData(page int, pagination_amount int) []entities.User {

	user_data_raw := sources.GetSampleUserData(page, pagination_amount)

	var user_data []entities.User

	err := json.Unmarshal(user_data_raw, &user_data)
	if err != nil {
		fmt.Printf("%v\n", err)
		panic("Failed to parse JSON user data")
	}

	return user_data
}

func stringifyUsers(users []entities.User) [][]string {

	var usersStringified [][]string

	for _, user := range users {

		currentUserReflect := reflect.ValueOf(user)
		currentUserPropsStrings := make([]string, currentUserReflect.NumField())

		for i := 0; i < currentUserReflect.NumField(); i++ {

			currentUserPropString := currentUserReflect.Field(i).Interface().(string)
			currentUserPropsStrings[i] = currentUserPropString
		}

		usersStringified = append(usersStringified, currentUserPropsStrings)
	}

	return usersStringified
}
