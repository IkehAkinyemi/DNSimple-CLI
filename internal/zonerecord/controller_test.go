package zonerecord

import (
	"flag"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

var configFileDir = flag.String("cfg-dir", "", "Directory path to configuration file")

func TestListZoneRecords(t *testing.T) {
	ctrl, err := New(*configFileDir)
	if err != nil {
		t.Fatalf("error reading config file: %+v", err)
	}

	// Set the output of the function to the buffer
	cmd := &cobra.Command{}

	// Set the value of the required flags
	cmd.Flags().String("cfg-dir", "", "")
	cmd.MarkFlagRequired("cfg-dir")
	cmd.Flags().Set("cfg-dir", *configFileDir)

	// Define flags
	cmd.Flags().String("name", "", "")
	cmd.Flags().String("type", "", "")
	cmd.Flags().String("name_like", "", "")

	// Set values
	cmd.Flags().Set("name", "test")
	cmd.Flags().Set("type", "A")
	cmd.Flags().Set("name_like", "test")

	// Call the listZoneRecords function
	ctrl.ListZoneRecords(cmd)

	// Check that the output of the function is as expected
	assert.Nil(t, err)
}

func TestCreateZoneRecord(t *testing.T) {
	ctrl, err := New(*configFileDir)
	if err != nil {
		t.Fatalf("error reading config file: %+v", err)
	}

	// Create a test CMD
	cmd := &cobra.Command{}

	// Set the value of the required flags
	cmd.Flags().String("cfg-dir", "", "")
	cmd.MarkFlagRequired("cfg-dir")
	cmd.Flags().Set("cfg-dir", *configFileDir)

	// Define flags
	cmd.Flags().String("name", "", "")
	cmd.Flags().String("type", "", "")
	cmd.Flags().String("content", "", "")
	cmd.Flags().Int("ttl", 0, "")
	cmd.Flags().Int("priority", 0, "")

	// Set values
	cmd.Flags().Set("name", "test")
	cmd.Flags().Set("type", "A")
	cmd.Flags().Set("content", "1.2.3.4")

	// Call the CreateZoneRecord function
	err = ctrl.CreateZoneRecord(cmd)

	// Check that the output of the function is as expected
	assert.Nil(t, err)

	// check for missing required fields
	cmd.Flags().Set("type", "SRV")
	cmd.Flags().Set("ttl", "0")

	err = ctrl.CreateZoneRecord(cmd)
	assert.NotNil(t, err)
}
