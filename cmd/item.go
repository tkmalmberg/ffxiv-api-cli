package cmd

import (
	"ffxiv/network"
	"ffxiv/resources"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	ItemID       int
	ItemResponse = resources.ItemSearchResponse{}
	itemCmd      = &cobra.Command{
		Use:   "item",
		Short: "Fetch an item using its unique id.",
		Long:  "Get the details of an item by entering the unique iq",
		Run: func(cmd *cobra.Command, args []string) {
			i := network.ItemSearchByID(ItemID)
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
