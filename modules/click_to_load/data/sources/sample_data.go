package click_to_load

import (
	"encoding/json"
	"strconv"
)

const (
	PAGINATION_AMOUNT          string = "pagination_amount"
	PAGINATION_VALID_RANGE_MIN int    = 5
	PAGINATION_VALID_RANGE_MAX int    = 100
)

func GetSampleUserData(page int, pagination_amount int) []byte {

	sample_users := []map[string]interface{}{}

	for i := page - 1; i < page*pagination_amount; i++ {

		new_user := map[string]interface{}{
			"name":  "Agent Smith",
			"email": strconv.Itoa(i) + "@matrix.com",
			"id":    "RandomId" + strconv.Itoa(i),
		}

		sample_users = append(sample_users, new_user)
	}

	sample_data_json, err := json.Marshal(sample_users)
	if err != nil {
		panic("Failed to get JSON user data")
	}

	return sample_data_json
}
