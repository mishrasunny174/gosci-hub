package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mishrasunny174/gosci-hub/libgoscihub"
)

func writeFile(filePath string, data []byte) (int64, error) {
	outFile, err := os.Create(filePath)
	if err != nil {
		return -1, err
	}
	defer outFile.Close()
	bytesWritten, err := io.Copy(outFile, bytes.NewReader(data))
	if err != nil {
		return -1, err
	}
	return bytesWritten, nil
}

func main() {
	url := flag.String("url", "", "Link of article or research paper you want to download")
	outFile := flag.String("outfile", "output.pdf", "Output file path")
	flag.Parse()
	if len(*url) == 0 {
		flag.Usage()
		return
	}
	data, err := libgoscihub.GetPDFFromArticleURL(*url)
	if err != nil {
		log.Fatal(err)
	}
	bytesWritten, err := writeFile(*outFile, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Article downloaded and saved as %s of size %d bytes\n", *outFile, bytesWritten)
}
