package utils

import (
	"encoding/json"
	"fmt"
	"github.com/c12s/cockpit/model"
	"io/ioutil"
)

const (
	// Path to files
	getStandaloneConfigFilePathJSON = "./standalone_config_files/standalone-config.json"
	getStandaloneConfigFilePathYAML = "./standalone_config_files/standalone-config.yaml"
)

func SaveStandaloneConfigResponseToFiles(response *model.SingleConfigGroupResponse, outputFormat string) error {
	if outputFormat == "json" {
		jsonData, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to convert to JSON: %v", err)
		}
		err = ioutil.WriteFile(getStandaloneConfigFilePathJSON, jsonData, 0644)
		if err != nil {
			return fmt.Errorf("failed to write JSON file: %v", err)
		}
		fmt.Printf("Standalone config saved to %s\n", getStandaloneConfigFilePathJSON)
	} else {
		yamlData, err := MarshalAppConfigResponseToYAML(response)
		if err != nil {
			return fmt.Errorf("failed to convert to YAML: %v", err)
		}
		err = ioutil.WriteFile(getStandaloneConfigFilePathYAML, yamlData, 0644)
		if err != nil {
			return fmt.Errorf("failed to write YAML file: %v", err)
		}
		fmt.Printf("Standalone config saved to %s\n", getStandaloneConfigFilePathYAML)
	}

	return nil
}
