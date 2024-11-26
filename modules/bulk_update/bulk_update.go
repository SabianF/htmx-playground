
package bulk_update

import "github.com/a-h/templ"

func BulkUpdate() templ.Component {
	return bulkUpdate()
}

func BulkUpdateToast(updatedUsers map[string][]string) templ.Component {
	return bulkUpdateToast(updatedUsers)
}
