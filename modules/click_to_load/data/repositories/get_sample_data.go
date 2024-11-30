package click_to_edit

import (
	"encoding/json"
	"fmt"

	click_to_edit_models "github.com/SabianF/htmx-playground/modules/click_to_load/data/models"
	click_to_edit_sources "github.com/SabianF/htmx-playground/modules/click_to_load/data/sources"
)

func GetUsers(page int, pagination_amount int) []click_to_edit_models.User {

	user_data_raw := click_to_edit_sources.GetSampleUserData(page, pagination_amount)

	var user_data []click_to_edit_models.User

	err := json.Unmarshal(user_data_raw, &user_data)
	if err != nil {
		fmt.Printf("%v\n", err)
		panic("Failed to parse JSON user data")
	}

	return user_data
}
