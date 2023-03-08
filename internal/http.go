package internal

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

var emptyResponse = gjson.Result{}

// GetProduct returns the product information from the endoflife.date API
func GetProduct(product string) (gjson.Result, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, "https://endoflife.date/api/"+product+".json", nil)
	if err != nil {
		return emptyResponse, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return emptyResponse, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return emptyResponse, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// convert response into string
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return emptyResponse, err
	}
	body := string(bodyBytes)

	return gjson.Parse(body), nil
}
