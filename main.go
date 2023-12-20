package main

import (
	"fmt"
	"sync"

	"github.com/devinmiller/fem-basics-of-go-client/api"
)

func main() {
	currencies := []string{"btc", "eth", "bch"}
	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)
		go func(currency string) {
			getCurrencyData(currency)
			wg.Done()
		}(currency)
	}

	wg.Wait()

}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}
