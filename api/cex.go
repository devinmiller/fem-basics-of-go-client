package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/devinmiller/fem-basics-of-go-client/models"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*models.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("three characters required; %d received", len(currency))
	}

	currency = strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, currency))

	if err != nil {
		return nil, err
	}

	var response RateResponse
	if res.StatusCode == http.StatusOK {
		data, err := io.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(data, &response); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("received status code %v", res.StatusCode)
	}

	lastPrice, err := strconv.ParseFloat(response.Last, 64)

	if err != nil {
		return nil, err
	}

	rate := models.Rate{Currency: currency, Price: lastPrice}
	return &rate, nil
}
