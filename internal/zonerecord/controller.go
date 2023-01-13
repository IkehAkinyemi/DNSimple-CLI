package zonerecord

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/IkehAkinyemi/zone-records-cli/pkg/model"
	"github.com/IkehAkinyemi/zone-records-cli/pkg/utils"
	"github.com/spf13/cobra"
)

// Controller defines a zone record controller.
type Controller struct {
	cfg *model.Config
}

// New create a zone record controller.
func New(fileDir string) (*Controller, error) {
	cfg, err := utils.GetConfig(fileDir)
	if err != nil {
		return nil, err
	}

	return &Controller{cfg}, nil
}

// CreateZoneRecord create a new zone record either A, AAA or SRV 
// record type.
func (c *Controller) CreateZoneRecord(cmd *cobra.Command) error {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	recordType, err := cmd.Flags().GetString("type")
	if err != nil {
		return err
	}
	content, err := cmd.Flags().GetString("content")
	if err != nil {
		return err
	}
	ttl, err := cmd.Flags().GetInt("ttl")
	if err != nil {
		return err
	}
	priority, err := cmd.Flags().GetInt("priority")
	if err != nil {
		return err
	}

	// validate input
	if recordType != "A" && recordType != "AAAA" && recordType != "SRV" {
		return fmt.Errorf("invalid record type: %s", recordType)
	}
	if name == "" && content == "" {
		return fmt.Errorf("name and content are required")
	}
	if recordType == "SRV" && priority == 0 {
		return fmt.Errorf("priority are required for SRV records")
	}

	// build request body
	var requestBody string
	if recordType == "SRV" {
		requestBody = fmt.Sprintf(`{"name":"%s","type":"%s","content":"%s","ttl":%d,"priority":%d}`, name, recordType, content, ttl, priority)
	} else {
		requestBody = fmt.Sprintf(`{"name":"%s","type":"%s","content":"%s","ttl":%d}`, name, recordType, content, ttl)
	}

	// build URL for API request
	var domainURL string
	if c.cfg.Env != "dev" {
		domainURL = "api.dnsimple.com"
	} else {
		domainURL = "api.sandbox.dnsimple.com"
	}

	// create HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/v2/%d/zones/%s/records", domainURL, c.cfg.AccountID, c.cfg.ZoneName), strings.NewReader(requestBody))
	if err != nil {
		return err
	}

	// add access token and content type to request header
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.cfg.AccessToken))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	// make request to API
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// handle response
	if resp.StatusCode == 400 {
		io.Copy(os.Stdout, resp.Body)
		return fmt.Errorf("\nbad request: %s", resp.Status)
	} else if resp.StatusCode == 404 {
		io.Copy(os.Stdout, resp.Body)
		return fmt.Errorf("\nnot found: %s", resp.Status)
	} else if resp.StatusCode == 401 {
		io.Copy(os.Stdout, resp.Body)
		return fmt.Errorf("\nauthentication failed: %s", resp.Status)
	} else if resp.StatusCode != 201 {
		return fmt.Errorf("error making request: %s", resp.Status)
	}

	// holds the deserialize json data
	var parsedResp = model.Envelope{"data": model.ZoneRecord{}}

	// code to deserialize returned json data
	err = utils.ReadJSON(resp, &parsedResp)

	// Print the result to the terminal
	fmt.Fprintf(os.Stdout, "%+v", parsedResp)

	return err
}

// ListZoneRecords retrieves all zone records, with the option
// to filter by name and type and name like a given string.
func (c *Controller) ListZoneRecords(cmd *cobra.Command) error {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	recordType, err := cmd.Flags().GetString("type")
	if err != nil {
		return err
	}
	nameLike, err := cmd.Flags().GetString("name_like")
	if err != nil {
		return err
	}

	// build query parameters for filtering
	queryParams := url.Values{}
	if name != "" {
		queryParams.Add("name", name)
	}
	if recordType != "" {
		queryParams.Add("type", recordType)
	}
	if nameLike != "" {
		queryParams.Add("name_like", nameLike)
	}

	// build URL for API request
	var apiURL string
	if c.cfg.Env != "dev" {
		apiURL = fmt.Sprintf("https://api.dnsimple.com/v2/%d/zones/%s/records?%s", c.cfg.AccountID, c.cfg.ZoneName, queryParams.Encode())
	} else {
		apiURL = fmt.Sprintf("https://api.sandbox.dnsimple.com/v2/%d/zones/%s/records?%s", c.cfg.AccountID, c.cfg.ZoneName, queryParams.Encode())
	}

	// create HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return err
	}

	// add access token to request header
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.cfg.AccessToken))
	req.Header.Add("Accept", "application/json")

	// make request to API
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// handle response
	if resp.StatusCode == 404 {
		io.Copy(os.Stdout, resp.Body)
		return fmt.Errorf("\nnot found: %s", resp.Status)
	} else if resp.StatusCode == 401 {
		io.Copy(os.Stdout, resp.Body)
		return fmt.Errorf("\nauthentication failed: %s", resp.Status)
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("error making request: %s", resp.Status)
	}

	// holds the list of zone records
	var zoneRecords model.ZoneRecords

	// code to deserialize returned json data
	err = utils.ReadJSON(resp, &zoneRecords)

	// Print the result to the terminal
	fmt.Fprintf(os.Stdout, "%+v", zoneRecords)

	return err
}
