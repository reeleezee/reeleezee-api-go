/*
Licensed under MIT license
(c) 2017 Reeleezee BV
*/
package misc

import "regexp"

// --------------------------------------------------------------------
// DynamicJsonValues, simple dynamic JSON representation
// --------------------------------------------------------------------
var re_nextlink = regexp.MustCompile("/api/[^/]+(.*)")

type DynamicJsonValues struct {
	NextLink string                   `json:"@odata.nextLink"`
	Value    []map[string]interface{} `json:"value"`
}

func (js *DynamicJsonValues) GetNextLink() string {
	if len(js.NextLink) > 0 {
		match := re_nextlink.FindStringSubmatch(js.NextLink)
		if match != nil && len(match) > 1 {
			return match[1]
		}
	}
	return ""
}
