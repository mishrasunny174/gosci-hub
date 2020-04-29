package main

import (
	"log"
)

func main() {
	pdfBuff, err := libgoscihub.GetPdfFromArticleURL("https://link.springer.com/referenceworkentry/10.1007%2F978-1-4419-1428-6_634")
	if err != nil {
		log.Println(err)
	}
	log.Printf("%v\n", pdfBuff)
}
