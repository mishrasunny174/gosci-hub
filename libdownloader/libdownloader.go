package libdownloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/dustin/go-humanize"
)

// DownloadFailedError is the error struct to be thrown when
// there is an error while downloading the file
type DownloadFailedError struct{}

type downloadProgressWriter struct {
	totalSize uint64
}

func (wp *downloadProgressWriter) Write(data []byte) (int, error) {
	wp.totalSize += uint64(len(data))
	fmt.Printf("\r%s", strings.Repeat(" ", 40))
	fmt.Printf("\rDownloading... %s", humanize.Bytes(wp.totalSize))
	return len(data), nil
}

func (DownloadFailedError) Error() string {
	return "ERROR: Download Failed"
}

// DownloadFile function takes an url and output file name as arguments
// @param url: url of the file to download
// @param outputFile: path of the file to download to
// it return the size of file downloaded in bytes if there is no error otherwise an error is
// returned
func DownloadFile(url string, outputFile string) (uint64, error) {
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	outfile, err := os.Create(outputFile)
	if err != nil {
		return 0, err
	}
	defer outfile.Close()
	bytesDownloaded, err := io.Copy(outfile, io.TeeReader(response.Body, &downloadProgressWriter{}))
	if err != nil {
		return 0, DownloadFailedError{}
	}
	return uint64(bytesDownloaded), nil
}
