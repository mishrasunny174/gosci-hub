package libgoscihub

import (
	"strings"
	"testing"
)

func TestGetPDFURL(t *testing.T) {
	inp := "https://link.springer.com/referenceworkentry/10.1007%2F978-1-4419-1428-6_634"
	got, err := getPDFURL(scihubURL, inp)
	expected := "https://twin.sci-hub.tw/5690/54fb8c9e958d81fe966762c612e4f836/10.1007@978-1-4419-1428-6634.pdf"
	if err != nil {
		t.Fatalf("Unknow error occured\nReturned value: %s\nERROR: %v\n", got, err)
	}
	if strings.Compare(got, expected) == 0 {
		t.Fatalf("Expected: %s\nGot: %s\n", expected, got)
	}
}
