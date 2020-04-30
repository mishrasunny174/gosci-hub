package libgoscihub

import (
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// PDFNotFoundError is a error struct which will be used to handle the error
// when unable to find the pdf for the given article
type PDFNotFoundError struct{}

func (PDFNotFoundError) Error() string {
	return "Unable to find pdf for the given article"
}

const (
	scihubURL = "https://sci-hub.tw/"
)

// GetPDFURL function will take the scihub base url and article url and it will return
// the url of the pdf for that article
func getPDFURL(scihubURL, articleURL string) (string, error) {
	var pdfURL string
	data := url.Values{
		"request": {articleURL},
	}
	resp, err := http.PostForm(scihubURL, data)
	if err != nil {
		return "", err
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	iframeSelection := doc.Find("iframe")
	href, ok := iframeSelection.Attr("src")
	if ok {
		pdfURL = href
	} else {
		return "", PDFNotFoundError{}
	}
	return pdfURL, nil
}
