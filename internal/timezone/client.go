package timezone

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var endpoint = "http://apiip.net/api/check?ip=%v&accessKey=%v"

func GetTimezone(ip string) (response, error) {
	data := new(response)
	req, _ := http.NewRequest("GET", fmt.Sprintf(endpoint, ip, os.Getenv("TIMEZONE_API_KEY")), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return *data, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return *data, fmt.Errorf("status code: %v", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return *data, err
	}
	err = json.Unmarshal(body, data)
	if err != nil {
		return *data, err
	}
	return *data, nil
}
