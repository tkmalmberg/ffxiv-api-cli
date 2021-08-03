package cmd

import (
	"ffxiv/network"
	"ffxiv/resources"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	SearchResponse = resources.CharacterSearchResponse{}
	Name, Server   string
	characterCmd   = &cobra.Command{
		Use:              "character",
		TraverseChildren: true,
		Short:            "Get a Player Character",
		Long:             "Get a Player Character by entering the name of your character and server they exist on.",
		// Args:             cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			charID := network.CharacterSearch(Name, Server)
			if charID != 0 {
				network.GetCharacterData(charID)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(characterCmd)
	characterCmd.Flags().StringVarP(&Name, "name", "n", "", "Name of the player character (required)")
	characterCmd.Flags().StringVarP(&Server, "server", "s", "", "Server name the player character exists on")
	err := characterCmd.MarkFlagRequired("name")
	if err != nil {
		log.Error("Name is required for character search.")
	}
}
