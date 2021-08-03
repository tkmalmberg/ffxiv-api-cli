package network

import (
	"ffxiv/resources"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	ItemID       int
	ItemResponse = resources.ItemSearchResponse{}
	// base_url     = "FFXIV_BASE_URL"
	// private_key  = "FFXIV_PRIVATE_KEY"
)

func ItemSearchByID(id int) interface{} {
	baseUrl := viper.GetString(base_url)
	pathUrl := "item/"
	requestUrl := fmt.Sprintf("%v%v", baseUrl, pathUrl)

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	req.Header.Add("private_key", viper.GetString(private_key))

	query := req.URL.Query()
	query.Add("ids", fmt.Sprint(id))

	req.URL.RawQuery = query.Encode()

	err = ApiRequestGet(*req, viper.GetString(private_key), &ItemResponse)
	if err != nil {
		log.Errorf("Error fetching API Request: %v", err)
	}

	return ItemResponse.Results[0]
}
