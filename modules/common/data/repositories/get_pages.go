package common

import (
	sources "github.com/SabianF/htmx-playground/modules/common/data/sources"
	entities "github.com/SabianF/htmx-playground/modules/common/domain/entities"
)

func GetAllPages() []entities.Page {
	return sources.GetAllPages()
}
