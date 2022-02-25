package powerstore

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) GetHost(clientInterface ClientInterface, hostURL string, hostName string) (*HostItem, error) {

	reqType := http.MethodGet

	// URL for volume resource type
	url := fmt.Sprintf("%s/api/rest/host?select=id,name?name=eq.%s", hostURL, hostName)

	body, err := clientInterface.doRequest(reqType, url, "")
	if err != nil {
		return nil, err
	}
	log.Println("body ===", string(body))
	var host []*HostItem
	err = json.Unmarshal(body, &host)
	if err != nil {
		return nil, err
	}
	return host[0], nil
}
