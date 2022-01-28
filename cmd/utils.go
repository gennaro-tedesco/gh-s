package cmd

import (
	"fmt"
	"net/url"
)

func checkNil(decoded interface{}, key string) string {
	val, ok := decoded.(map[string]interface{})[key]
	if ok && val != nil {
		return val.(string)
	}
	return ""
}

func parseInput(search string, lang string, desc string) url.Values {
	queryString := fmt.Sprintf("%s in:name", search)
	if lang != "" {
		queryString = queryString + fmt.Sprintf(" language:%s", lang)
	}
	if desc != "" {
		queryString = queryString + fmt.Sprintf(" %s in:description", desc)
	}
	query := url.Values{}
	query.Add("q", queryString)
	query.Add("sort", "stars")
	query.Add("per_page", "30")
	return query
}
