package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "zone-records-cli",
	Short: "CLI for managing zone records via the DNSimple API",
}

func main() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(createCmd)

	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().String("cfg-dir", "", "Directory path to configuration file")
}
