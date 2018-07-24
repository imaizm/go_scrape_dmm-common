package goScrapeDmmCommon

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// GetItemCodeFromURL : get ItemCode from url
func GetItemCodeFromURL(url string) string {
	cidMatcher := regexp.MustCompile(`cid=([^/]+)`)
	itemCode := cidMatcher.FindString(url)
	itemCode = cidMatcher.ReplaceAllString(itemCode, "$1")
	return itemCode
}

// GetSampleImageList : get ItemCode from goquery.Document
func GetSampleImageList(doc *goquery.Document) []*SampleImage {
	var sampleImageList []*SampleImage

	sampleImageURLMatcher := regexp.MustCompile(`([^-]+)(-\d+\..+)`)

	doc.Find("#sample-image-block > a").Each(func(index int, selection *goquery.Selection) {
		sampleImage := SampleImage{}

		imgSrc, exists := selection.Find("img").First().Attr("src")
		if exists {
			sampleImage.ImageThumbURL = imgSrc

			imageURL :=
				sampleImageURLMatcher.ReplaceAllString(imgSrc, "$1") + "jp" +
					sampleImageURLMatcher.ReplaceAllString(imgSrc, "$2")

			sampleImage.ImageURL = imageURL
		}

		sampleImageList = append(sampleImageList, &sampleImage)
	})

	return sampleImageList
}
