package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dustin/go-humanize"
	"github.com/mishrasunny174/gosci-hub/libdownloader"
	"github.com/mishrasunny174/gosci-hub/libgoscihub"
)

func main() {
	url := flag.String("url", "", "Link of article or research paper you want to download")
	outFile := flag.String("outfile", "output.pdf", "Output file path")
	flag.Parse()
	if len(*url) == 0 {
		flag.Usage()
		return
	}
	pdfURL, err := libgoscihub.GetPDFURL(*url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found article\nStarting Download")
	bytesDownloaded, err := libdownloader.DownloadFile(pdfURL, *outFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Downloaded %s of size %d bytes\n", *outFile, humanize.Bytes(bytesDownloaded))
}
