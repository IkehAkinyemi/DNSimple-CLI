package main

import (
	"fmt"
	"log"

	"github.com/IkehAkinyemi/zone-records-cli/internal/zonerecord"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all zone records",
	Run: func(cmd *cobra.Command, args []string) {
		cfgFileDir, err := cmd.Flags().GetString("cfg-dir")
		if err != nil {
			log.Fatal(err)
		}

		ctrl, err := zonerecord.New(cfgFileDir)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Listing zone records...")

		err = ctrl.ListZoneRecords(cmd)
		if err != nil {
			log.Fatalf(err.Error())
		}
	},
}

func init() {
	listCmd.Flags().String("name", "", "filter records by name")
	listCmd.Flags().String("type", "", "filter records by type")
	listCmd.Flags().String("name_like", "", "filter records by name like a given string")
	createCmd.Flags().Int("ttl", 14400, "Time to Live (TTL)")
}
