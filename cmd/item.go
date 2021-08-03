package cmd

import (
	"ffxiv/network"
	"ffxiv/resources"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ItemID       int
	ItemResponse = resources.ItemSearchResponse{}
	itemCmd      = &cobra.Command{
		Use:   "item",
		Short: "Fetch an item using its unique id.",
		Long:  "Get the details of an item by entering the unique iq",
		Run: func(cmd *cobra.Command, args []string) {
			i := ItemSearchByID(ItemID)
			log.Infoln(i)
		},
	}
)

func init() {
	rootCmd.AddCommand(itemCmd)
	itemCmd.Flags().IntVarP(&ItemID, "itemid", "i", 0, "An item's unique ID (required)")
	err := itemCmd.MarkFlagRequired("itemid")
	if err != nil {
		log.Error("Item ID is required for item fetch.")
	}
}

func ItemSearchByID(id int) string {
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

	err = network.ApiRequestGet(*req, viper.GetString(private_key), &ItemResponse)
	if err != nil {
		log.Errorf("Error fetching API Request: %v", err)
	}

	return ItemResponse.Results[0].Name

}
