package test

import (
	"net/http"
	"testing"
)

func TestCurrencyGet(t *testing.T) {
	url := "http://golang.yigitaltunay.com/keys"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("network error")
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {

	} else {
		t.Error("Invalid response code")
	}
}
