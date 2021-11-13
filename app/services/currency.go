package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"sync"

	"github.com/yigitaltunay/go-assignment/app/models"
)

var config models.Config

func SetConfig(configLoad models.Config) {
	config = configLoad
}

func SendRequestProvider(provider, currentType string, ch chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(provider)
	if err != nil {
		log.Println("network problem")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error calls io.ReadAll")
	}
	current := models.Currency{}
	json.Unmarshal(b, &current) // error did not hold, we are sure about the request
	price, _ := TypeCurrentWithName(currentType, current)
	ch <- price
}

func FindCurrencyCurrencies(currentType string) (models.Response, error) {
	ch := make(chan float64)
	var responses []float64
	var wg sync.WaitGroup

	for _, provider := range config.Providers {
		wg.Add(1)
		go SendRequestProvider(provider, currentType, ch, &wg)
	}

	// close the channel in the background
	go func() {
		wg.Wait()
		close(ch)
	}()
	// read from channel as they come in until its closed
	for res := range ch {
		//	log.Println(res)
		responses = append(responses, res)
	}

	if len(responses) == 0 {
		return models.Response{}, fmt.Errorf("no response")
	}
	sort.Float64s(responses)
	return models.Response{Data: responses[0], Currency: currentType}, nil
}

func TypeCurrentWithName(currentType string, current models.Currency) (float64, error) {
	switch currentType {
	case "USD":
		return current.Usd, nil
	case "EUR":
		return current.Eur, nil
	case "GBP":
		return current.Gbp, nil
	default:
		return 0, fmt.Errorf("current Type not supported")
	}
}
