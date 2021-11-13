package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

const testUrl = "http://localhost:8080/"

func TestCurrencyGetUSD(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("%scurrency/USD", testUrl))
	if err != nil {
		t.Error("network error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode == http.StatusOK {
		log.Println(string(body))
	} else {
		t.Error("Invalid response code")
	}
}

func TestCurrencyGetEUR(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("%scurrency/EUR", testUrl))
	if err != nil {
		t.Error("network error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode == http.StatusOK {
		log.Println(string(body))
	} else {
		t.Error("Invalid response code")
	}
}

func TestCurrencyGetGBP(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("%scurrency/GBP", testUrl))
	if err != nil {
		t.Error("network error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode == http.StatusOK {
		log.Println(string(body))
	} else {
		t.Error("Invalid response code")
	}
}
