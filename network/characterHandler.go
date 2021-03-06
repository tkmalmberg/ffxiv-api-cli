package network

import (
	"encoding/json"
	"ffxiv/resources"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	SearchResponse = resources.CharacterSearchResponse{}
	// base_url       = "FFXIV_BASE_URL"
	// private_key    = "FFXIV_PRIVATE_KEY"
)

func CharacterSearch(name string, server string) (charId int) {
	baseUrl := viper.GetString(base_url)
	pathUrl := "character/search"
	requestUrl := fmt.Sprintf("%v%v", baseUrl, pathUrl)

	log.Infof("Searching for %s on the server %s...", name, server)

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	req.Header.Add("private_key", viper.GetString(private_key))

	query := req.URL.Query()
	query.Add("name", name)
	query.Add("server", server)

	req.URL.RawQuery = query.Encode()

	log.Info("Sending request...")
	err = ApiRequestGet(*req, viper.GetString(private_key), &SearchResponse)
	if err != nil {
		log.Errorf("Error fetching API Request: %v", err)
	}
	log.Info("Request Fulfilled...")

	if len(SearchResponse.Results) > 1 {
		log.Info("Multiple Characters Found:")
		for _, item := range SearchResponse.Results {
			fmt.Println(item.Name)
		}
		charId = 0
		return charId
	}

	// TODO: Add cases for multiple results to be found and listed
	charId = SearchResponse.Results[0].ID

	return charId
}

func GetCharacterData(id int) {
	baseUrl := viper.GetString(base_url)
	pathUrl := fmt.Sprintf("character/%v", id)
	requestUrl := fmt.Sprintf("%v%v", baseUrl, pathUrl)
	Response := resources.CharResponse{}

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	req.Header.Add("private_key", viper.GetString(private_key))

	query := req.URL.Query()

	req.URL.RawQuery = query.Encode()

	log.Infof("Fetching character info for: %v", id)
	err = ApiRequestGet(*req, viper.GetString(private_key), &Response)
	if err != nil {
		log.Errorf("Error fetching API Request: %v", err)
	}

	fmtCharacter, err := json.MarshalIndent(Response.Character, "", "    ")
	if err != nil {
		log.Error(err)
	}

	log.Infof("%s", fmtCharacter)
}
