package network

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	base_url    = "FFXIV_BASE_URL"
	private_key = "FFXIV_PRIVATE_KEY"
)

func ApiRequestGet(req http.Request, token string, class interface{}) error {

	res, err := http.Get(req.URL.String())
	if err != nil {
		log.Error(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
	}

	err = json.Unmarshal(body, &class)
	if err != nil {
		log.Error(err)
	}

	return nil
}
