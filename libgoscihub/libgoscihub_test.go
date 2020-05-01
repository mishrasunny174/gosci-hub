package libgoscihub

import (
	"regexp"
	"testing"
)

func TestGetPDFURL(t *testing.T) {
	inps := []string{
		"https://link.springer.com/referenceworkentry/10.1007%2F978-1-4419-1428-6_634",
		"https://www.researchgate.net/publication/282284055_When_Every_Byte_Counts_-_Writing_Minimal_Length_Shellcodes",
		"https://sci-hub.tw/https://www.researchgate.net/publication/261123624_JOP-alarm_Detecting_jump-oriented_programming-based_anomalies_in_applications",
		"https://www.researchgate.net/publication/324558339_Predicting_Cyber_Events_by_Leveraging_Hacker_Sentiment",
	}
	for _, inp := range inps {
		got, err := GetPDFURL(inp)
		if err != nil {
			t.Fatalf("Unknow error occured\nReturned value: %s\nERROR: %v\n", got, err)
		}
		matched, err := regexp.MatchString(`https://.*\.pdf`, got)
		if err != nil {
			t.Fatal(err)
		}
		if !matched {
			t.Fatalf("Failed: Unable to find the pdf for test case %s", inp)
		}
	}
}
