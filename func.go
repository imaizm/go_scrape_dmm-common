package goScrapeDmmCommon

import (
	"regexp"
)

// GetItemCodeFromURL : get ItemCode from url
func GetItemCodeFromURL(url string) string {
	cidMatcher := regexp.MustCompile(`cid=([^/]+)`)
	itemCode := cidMatcher.FindString(url)
	itemCode = cidMatcher.ReplaceAllString(itemCode, "$1")
	return itemCode
}
