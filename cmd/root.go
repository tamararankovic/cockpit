package cmd

import (
	"fmt"
	auth "github.com/c12s/cockpit/cmd/auth"
	claim "github.com/c12s/cockpit/cmd/claim"
	create "github.com/c12s/cockpit/cmd/create"
	delete "github.com/c12s/cockpit/cmd/delete"
	get "github.com/c12s/cockpit/cmd/get"
	list "github.com/c12s/cockpit/cmd/list"
	put "github.com/c12s/cockpit/cmd/put"
	validate "github.com/c12s/cockpit/cmd/validate"
	"github.com/spf13/cobra"
	"os"
)

const (
	apiVersionFlag = "api-version"
)

func init() {
	RootCmd.AddCommand(auth.LoginCmd)
	RootCmd.AddCommand(auth.RegisterCmd)

	ListCmd.AddCommand(list.NodesCmd)
	list.NodesCmd.AddCommand(list.AllocatedNodesCmd)
	RootCmd.AddCommand(ListCmd)

	PutCmd.AddCommand(put.LabelsCmd)
	RootCmd.AddCommand(PutCmd)

	DeleteCmd.AddCommand(delete.DeleteNodeLabelsCmd)
	DeleteCmd.AddCommand(delete.DeleteSchemaCmd)
	RootCmd.AddCommand(DeleteCmd)

	ClaimCmd.AddCommand(claim.ClaimNodesCmd)
	RootCmd.AddCommand(ClaimCmd)

	GetCmd.AddCommand(get.GetSchemaCmd)
	GetCmd.AddCommand(ConfigCmd)
	ConfigCmd.AddCommand(get.GetConfigGroupCmd)
	get.GetSchemaCmd.AddCommand(get.GetSchemaVersionCmd)
	RootCmd.AddCommand(GetCmd)

	ValidateCmd.AddCommand(validate.ValidateSchemaVersionCmd)
	RootCmd.AddCommand(ValidateCmd)

	CreateCmd.AddCommand(create.CreateSchemaCmd)
	RootCmd.AddCommand(CreateCmd)

	RootCmd.PersistentFlags().String(apiVersionFlag, "1.0.0", "specify c12s API version")
}

var ClaimCmd = &cobra.Command{
	Use:   "claim",
	Short: "Claim resources",
}

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete resources",
}

var PutCmd = &cobra.Command{
	Use:   "put",
	Short: "Put resources",
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List resources",
}

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create resources",
}

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resources",
}

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manipulate with config",
}

var ValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Get resources",
}

var RootCmd = &cobra.Command{
	Use:   "cockpit",
	Short: "Cockpit is a CLI tool for interacting with the c12s system",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
