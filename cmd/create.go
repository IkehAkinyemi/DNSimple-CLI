package main

import (
	"fmt"
	"log"

	"github.com/IkehAkinyemi/zone-records-cli/internal/zonerecord"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new zone record",
	Run: func(cmd *cobra.Command, args []string) {
		cfgFileDir, err := cmd.Flags().GetString("cfg-dir")
		if err != nil {
			log.Fatal(err)
		}

		ctrl, err := zonerecord.New(cfgFileDir)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Creating zone record...")

		err = ctrl.CreateZoneRecord(cmd)
		if err != nil {
			log.Fatalf(err.Error())
		}
	},
}

func init() {
	createCmd.Flags().String("name", "", "Name of the resource (required)")
	createCmd.Flags().String("type", "", "type of the record (A, AAAA, or SRV) (required)")
	createCmd.Flags().String("content", "", "content of the record (required)")
	createCmd.Flags().Int("priority", 0, "priority of the record (required for SRV records)")
	createCmd.Flags().Int("ttl", 14400, "Time to Live (TTL)")
}
