package cmd

import (
	"fmt"
	"github.com/c12s/cockpit/clients"
	"github.com/c12s/cockpit/model"
	"github.com/c12s/cockpit/render"
	"github.com/c12s/cockpit/utils"
	"os"
	"time"

	"github.com/spf13/cobra"
)

const (
	claimNodesShortDescription = "Claim nodes for an organization based on specific criteria"
	claimNodesLongDescription  = "Claims nodes for an organization based on a defined query that specifies criteria like labels.\n\n" +
		"Example:\n" +
		"cockpit claim nodes --org 'org' --query 'labelKey >||=||!=||< value'\n" +
		"cockpit claim nodes --org 'org' --query 'memory-totalGB > 2'"

	// Flag Constants
	organizationFlag = "org"
	queryFlag        = "query"

	// Flag Shorthand Constants
	organizationFlagShortHand = "r"
	queryFlagShortHand        = "q"

	// Flag Descriptions
	organizationDesc = "Organization name (required)"
	queryDesc        = "Query label for finding specific nodes (required)"
)

var (
	org               string
	query             string
	claimNodeResponse model.ClaimNodesResponse
)

var ClaimNodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: claimNodesShortDescription,
	Long:  claimNodesLongDescription,
	Run:   executeClaimNodes,
}

func executeClaimNodes(cmd *cobra.Command, args []string) {
	requestBody, err := prepareClaimNodesRequest()
	if err != nil {
		fmt.Println("Error preparing request:", err)
		os.Exit(1)
	}

	if err := sendClaimNodeRequest(requestBody); err != nil {
		fmt.Printf("Error claiming nodes: %v\n", err)
		os.Exit(1)
	}

	render.RenderNodes(claimNodeResponse.Nodes)
}

func prepareClaimNodesRequest() (interface{}, error) {
	request := model.ClaimNodesRequest{
		Org: org,
	}

	nodeQueries, err := utils.CreateNodeQuery(query)
	if err != nil {
		return nil, err
	}
	request.Query = nodeQueries

	return request, nil
}

func sendClaimNodeRequest(requestBody interface{}) error {
	token, err := utils.ReadTokenFromFile()
	if err != nil {
		return fmt.Errorf("error reading token: %v", err)
	}

	url := clients.BuildURL("core", "v1", "ClaimOwnership")

	return utils.SendHTTPRequest(model.HTTPRequestConfig{
		URL:         url,
		Method:      "PATCH",
		Token:       token,
		RequestBody: requestBody,
		Response:    &claimNodeResponse,
		Timeout:     10 * time.Second,
	})
}

func init() {
	ClaimNodesCmd.Flags().StringVarP(&org, organizationFlag, organizationFlagShortHand, "", organizationDesc)
	ClaimNodesCmd.Flags().StringVarP(&query, queryFlag, queryFlagShortHand, "", queryDesc)

	ClaimNodesCmd.MarkFlagRequired(organizationFlag)
	ClaimNodesCmd.MarkFlagRequired(queryFlag)
}
