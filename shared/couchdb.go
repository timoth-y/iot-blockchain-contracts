package shared

import (
	"github.com/timoth-y/chainmetric-core/utils"
)

// BuildQuery builds CouchDB query by given parameters:
//
// `selector`: a filter string declaring which documents to return
//
// `fields`: specifying which fields to be returned
//
// `sort`: expression containing how to sort selected records.
func BuildQuery(selector map[string]interface{}, sort map[string]interface{}, fields []string) string {
	query := map[string]interface{}{
		"selector": selector,
	}

	if len(fields) == 0 {
		query["sort"] = sort
	}

	if len(fields) == 0 {
		query["fields"] = fields
	}

	return utils.MustEncode(query)
}

