package test

import (
	"testing"

	"github.com/yigitaltunay/go-assignment/app/models"
	"github.com/yigitaltunay/go-assignment/app/services"
	"github.com/yigitaltunay/go-assignment/config"
)

func TestTypeCurrentWithName(t *testing.T) {
	config.ConfigFilePath = "../config/config.json"
	config.InitServe()

	_, err := services.TypeCurrentWithName("USD", models.Currency{})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	_, err = services.TypeCurrentWithName("EUR", models.Currency{})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	_, err = services.TypeCurrentWithName("GBP", models.Currency{})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

}

func TestFindCurrencyCurrencies(t *testing.T) {

	_, err := services.FindCurrencyCurrencies("USD")

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	_, err = services.FindCurrencyCurrencies("EUR")

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	_, err = services.FindCurrencyCurrencies("GBP")

	if err != nil {
		t.Errorf("Error: %s", err)
	}

}
