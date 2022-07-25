package pterodactyl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gamemann/Rust-Auto-Wipe/processor"
)

func SendAPIRequest(wipedata *processor.WipeData, request_type string, request_endpoint string, post_data map[string]string) (string, int, error) {
	// Initialize data and return code (status code).
	d := ""
	rc := -1

	// Compile our URL.
	urlstr := wipedata.APIURL + "/api/" + request_endpoint

	// Setup HTTP GET request.
	client := &http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(request_type, urlstr, nil)

	if err != nil {
		fmt.Println(err)

		return d, rc, err
	}

	// Set Application API token.
	req.Header.Set("Authorization", "Bearer "+wipedata.APIToken)

	// Accept only JSON.
	req.Header.Set("Accept", "application/json")

	// Check to see if we need to send post data.
	if request_type == "POST" {
		// Set POST data.
		for key, value := range post_data {
			req.PostForm.Add(key, value)
		}
	}

	// Perform HTTP request and check for errors.
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)

		return d, rc, err
	}

	// Set return code.
	rc = resp.StatusCode

	// Read body.
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)

		return d, rc, err
	}

	// Return data as a string.
	d = string(body)

	return d, rc, nil
}

func ProcessResponse(data []byte, structure *interface{}) error {
	err := json.Unmarshal(data, structure)

	return err
}
