package loyalty

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LoyaltySystem interface {
	GetAccrual(orderNumber string) (float64, error)
}

type AccrualSystem struct {
	baseURL string
}

func NewAccrualSystem(url string) *AccrualSystem {
	return &AccrualSystem{
		baseURL: url,
	}
}

type accrualResponse struct {
	Accrual float64 `json:"accrual"`
	Status  string  `json:"status"`
	Order   string  `json:"order"`
}

func (a *AccrualSystem) GetAccrual(orderNumber string) (float64, error) {
	reqURL := fmt.Sprintf("%s/api/orders/%s", a.baseURL, orderNumber)
	fmt.Println("request URL", reqURL)
	resp, err := http.Get(reqURL)
	if err != nil {
		return 0, fmt.Errorf("error getting Accrual: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return 0, fmt.Errorf("error reading response body: %v", err)
		}

		var accrualResponse accrualResponse
		err = json.Unmarshal(body, &accrualResponse)
		if err != nil {
			return 0, fmt.Errorf("error unmarshaling response: %v", err)
		}

		log.Printf("result: %v", accrualResponse)
		return accrualResponse.Accrual, nil
	}

	log.Printf("response status code: %d", resp.StatusCode)
	return 0, fmt.Errorf("error getting Accrual: %v", err)
}
