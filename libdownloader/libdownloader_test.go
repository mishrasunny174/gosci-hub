package libdownloader

import (
	"testing"

	"github.com/dustin/go-humanize"
)

func TestDownloadFile(t *testing.T) {
	inputURL := "http://mirror.filearena.net/pub/speed/SpeedTest_16MB.dat?_ga=2.86128028.541775211.1588223194-1421992388.1588223194"
	inputFile := "16MB.dat"
	expectedFileSize := uint64(16 * 1024 * 1024)
	bytesDownloaded, err := DownloadFile(inputURL, inputFile)
	if err != nil {
		t.Fatal(err)
	}
	if bytesDownloaded != expectedFileSize {
		t.Fatalf("Download Test failed: Expected size %s and got %s\n", humanize.Bytes(expectedFileSize), humanize.Bytes(bytesDownloaded))
	}
}
