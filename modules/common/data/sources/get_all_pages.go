package common

import (
	bulk_update "github.com/SabianF/htmx-playground/modules/bulk_update/data/repositories"
	click_me "github.com/SabianF/htmx-playground/modules/click_me/data/repositories"
	click_to_edit "github.com/SabianF/htmx-playground/modules/click_to_edit/data/repositories"
	click_to_load "github.com/SabianF/htmx-playground/modules/click_to_load/data/repositories"
	entities "github.com/SabianF/htmx-playground/modules/common/domain/entities"
	delete_row "github.com/SabianF/htmx-playground/modules/delete_row/data/repositories"
	hello "github.com/SabianF/htmx-playground/modules/hello/data/repositories"
)

func GetAllPages() []entities.Page {
	return []entities.Page{
		{Route: hello.ROUTE_PAGE 					, Title: "Hello Example"},
		{Route: click_me.ROUTE_PAGE				, Title: "Click me"			},
		{Route: click_to_edit.ROUTE_PAGE	, Title: "Click to Edit"},
		{Route: bulk_update.ROUTE_PAGE		, Title: "Bulk Update"	},
		{Route: click_to_load.ROUTE_PAGE	, Title: "Click to Load"},
		{Route: delete_row.ROUTE_PAGE			, Title: "Delete Row"		},
	}
}
