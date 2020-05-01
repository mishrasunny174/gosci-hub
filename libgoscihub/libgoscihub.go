package libgoscihub

import (
	"net/http"
	"net/url"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// PDFNotFoundError is a error struct which will be used to handle the error
// when unable to find the pdf for the given article
type PDFNotFoundError struct{}

func (PDFNotFoundError) Error() string {
	return "Unable to find pdf for the given article"
}

var (
	scihubURLs = []string{
		"https://sci-hub.tw/",
		"https://sci-hub.im/",
		"https://scihub.wikicn.top/",
		"https://sci-hub.ren/",
		"https://sci-hub.se/",
		"https://sci-hub.shop/",
	}
)

// GetPDFURL function will take the scihub base url and article url and it will return
// the url of the pdf for that article
func GetPDFURL(articleURL string) (string, error) {
	data := url.Values{
		"request": {articleURL},
	}
	for _, scihubURL := range scihubURLs {
		resp, err := http.PostForm(scihubURL, data)
		if err != nil {
			continue
		}
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			continue
		}
		iframeSelection := doc.Find("iframe")
		pdfURL, ok := iframeSelection.Attr("src")
		if ok {
			if matched, err := regexp.MatchString(`https:.*`, pdfURL); err == nil && matched {
				return pdfURL, nil
			}
			pdfURL = "https:" + pdfURL
			return pdfURL, nil
		}
	}
	return "", PDFNotFoundError{}
}
