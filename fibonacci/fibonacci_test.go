package fibonacci

import (
	"log"
	"os"
	"testing"
)

var testCasesSlice = []struct {
	x, y          string
	expectedList  []string
	expectedError error
}{
	{
		x:             "0",
		y:             "10",
		expectedList:  []string{"0", "1", "1", "2", "3", "5", "8", "13", "21", "34", "55"},
		expectedError: nil,
	},
	{
		x:             "10",
		y:             "10",
		expectedList:  []string{"55"},
		expectedError: nil,
	},
}

var testCasesElement = []struct {
	i               int
	expectedElement string
}{
	{
		i:               10,
		expectedElement: "55",
	},
	{
		i:               0,
		expectedElement: "0",
	},
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestGetFibSlice(t *testing.T) {
	err := os.Setenv("redis_ip", "127.0.0.1")
	if err != nil {
		log.Println(err)
	}
	log.Println(os.Getenv("redis_ip"))
	err = os.Setenv("redis_port", "6379")
	if err != nil {
		log.Println(err)
	}
	for _, tc := range testCasesSlice {
		sl, _, err := GetFibSlice(tc.x, tc.y)
		if !equal(sl, tc.expectedList) {
			t.Errorf("GetFibSlice(%v, %v) got %v, expected %v", tc.x, tc.y, sl, tc.expectedList)
		} else if err != tc.expectedError {
			t.Errorf("GetFibSlice(%v, %v) got %v, expected %v", tc.x, tc.y, err, tc.expectedError)
		}
	}
}

func TestCalculateFibElement(t *testing.T) {
	for _, tc := range testCasesElement {
		el, _ := getFibElementFromCache(tc.i)
		if el != tc.expectedElement {
			t.Errorf("getFibElementFromCache(%v) got %v, expected %v", tc.i, el, tc.expectedElement)
		}
	}
}
